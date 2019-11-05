package comm_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"gostudy/core/comm"
	testpb "gostudy/core/comm/testdata/grpc"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"testing"
	"time"
)


// Embedded certificates for testing
// The self-signed cert expires in 2028
var selfSignedKeyPEM = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIMLemLh3+uDzww1pvqP6Xj2Z0Kc6yqf3RxyfTBNwRuuyoAoGCCqGSM49
AwEHoUQDQgAEDB3l94vM7EqKr2L/vhqU5IsEub0rviqCAaWGiVAPp3orb/LJqFLS
yo/k60rhUiir6iD4S4pb5TEb2ouWylQI3A==
-----END EC PRIVATE KEY-----
`
var selfSignedCertPEM = `-----BEGIN CERTIFICATE-----
MIIBdDCCARqgAwIBAgIRAKCiW5r6W32jGUn+l9BORMAwCgYIKoZIzj0EAwIwEjEQ
MA4GA1UEChMHQWNtZSBDbzAeFw0xODA4MjExMDI1MzJaFw0yODA4MTgxMDI1MzJa
MBIxEDAOBgNVBAoTB0FjbWUgQ28wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQM
HeX3i8zsSoqvYv++GpTkiwS5vSu+KoIBpYaJUA+neitv8smoUtLKj+TrSuFSKKvq
IPhLilvlMRvai5bKVAjco1EwTzAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYI
KwYBBQUHAwEwDAYDVR0TAQH/BAIwADAaBgNVHREEEzARgglsb2NhbGhvc3SHBH8A
AAEwCgYIKoZIzj0EAwIDSAAwRQIgOaYc3pdGf2j0uXRyvdBJq2PlK9FkgvsUjXOT
bQ9fWRkCIQCr1FiRRzapgtrnttDn3O2fhLlbrw67kClzY8pIIN42Qw==
-----END CERTIFICATE-----
`

var badPEM = `-----BEGIN CERTIFICATE-----
MIICRDCCAemgAwIBAgIJALwW//dz2ZBvMAoGCCqGSM49BAMCMH4xCzAJBgNVBAYT
AlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2Nv
MRgwFgYDVQQKDA9MaW51eEZvdW5kYXRpb24xFDASBgNVBAsMC0h5cGVybGVkZ2Vy
MRIwEAYDVQQDDAlsb2NhbGhvc3QwHhcNMTYxMjA0MjIzMDE4WhcNMjYxMjAyMjIz
MDE4WjB+MQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UE
BwwNU2FuIEZyYW5jaXNjbzEYMBYGA1UECgwPTGludXhGb3VuZGF0aW9uMRQwEgYD
VQQLDAtIeXBlcmxlZGdlcjESMBAGA1UEAwwJbG9jYWxob3N0MFkwEwYHKoZIzj0C
-----END CERTIFICATE-----
`

var pemNoCertificateHeader = `-----BEGIN NOCERT-----
MIICRDCCAemgAwIBAgIJALwW//dz2ZBvMAoGCCqGSM49BAMCMH4xCzAJBgNVBAYT
AlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2Nv
MRgwFgYDVQQKDA9MaW51eEZvdW5kYXRpb24xFDASBgNVBAsMC0h5cGVybGVkZ2Vy
MRIwEAYDVQQDDAlsb2NhbGhvc3QwHhcNMTYxMjA0MjIzMDE4WhcNMjYxMjAyMjIz
MDE4WjB+MQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UE
BwwNU2FuIEZyYW5jaXNjbzEYMBYGA1UECgwPTGludXhGb3VuZGF0aW9uMRQwEgYD
VQQLDAtIeXBlcmxlZGdlcjESMBAGA1UEAwwJbG9jYWxob3N0MFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAEu2FEZVSr30Afey6dwcypeg5P+BuYx5JSYdG0/KJIBjWK
nzYo7FEmgMir7GbNh4pqA8KFrJZkPuxMgnEJBZTv+6NQME4wHQYDVR0OBBYEFAWO
4bfTEr2R6VYzQYrGk/2VWmtYMB8GA1UdIwQYMBaAFAWO4bfTEr2R6VYzQYrGk/2V
WmtYMAwGA1UdEwQFMAMBAf8wCgYIKoZIzj0EAwIDSQAwRgIhAIelqGdxPMHmQqRF
zA85vv7JhfMkvZYGPELC7I2K8V7ZAiEA9KcthV3HtDXKNDsA6ULT+qUkyoHRzCzr
A4QaL2VU6i4=
-----END NOCERT-----
`


var timeout = time.Second * 1
var testOrgs = []testOrg{}

func init() {
	//load up crypto material for test orgs
	for i := 1; i <= numOrgs; i++ {
		testOrg, err := loadOrg(i)
		if err != nil {
			log.Fatalf("Failed to load test organizations due to error: %s", err.Error())
		}
		testOrgs = append(testOrgs, testOrg)
	}
}


// utility function to load crypto material for organizations
func loadOrg(parent int) (testOrg, error) {

	var org = testOrg{}
	// load the CA
	caPEM, err := ioutil.ReadFile(fmt.Sprintf(orgCACert, parent))
	if err != nil {
		return org, err
	}
	// loop through and load servers
	var serverCerts = []serverCert{}
	for i := 1; i <= numServerCerts; i++ {
		keyPEM, err := ioutil.ReadFile(fmt.Sprintf(orgServerKey, parent, i))
		if err != nil {
			return org, err
		}
		certPEM, err := ioutil.ReadFile(fmt.Sprintf(orgServerCert, parent, i))
		if err != nil {
			return org, err
		}
		serverCerts = append(serverCerts, serverCert{keyPEM, certPEM})
	}
	// loop through and load clients
	var clientCerts = []tls.Certificate{}
	for j := 1; j <= numServerCerts; j++ {
		clientCert, err := loadTLSKeyPairFromFile(fmt.Sprintf(orgClientKey, parent, j),
			fmt.Sprintf(orgClientCert, parent, j))
		if err != nil {
			return org, err
		}
		clientCerts = append(clientCerts, clientCert)
	}
	// loop through and load child orgs
	var childOrgs = []testOrg{}

	for k := 1; k <= numChildOrgs; k++ {
		childOrg, err := loadChildOrg(parent, k)
		if err != nil {
			return org, err
		}
		childOrgs = append(childOrgs, childOrg)
	}

	return testOrg{caPEM, serverCerts, clientCerts, childOrgs}, nil
}

// invoke the EmptyCall RPC
func invokeEmptyCall(address string, dialOptions []grpc.DialOption) (*testpb.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//create GRPC client conn
	clientConn, err := grpc.DialContext(ctx, address, dialOptions...)
	if err != nil {
		return nil, err
	}
	defer clientConn.Close()

	//create GRPC client
	client := testpb.NewEmptyServiceClient(clientConn)

	//invoke service
	empty, err := client.EmptyCall(context.Background(), new(testpb.Empty))
	if err != nil {
		return nil, err
	}

	return empty, nil
}
// utility function to load crypto material for child organizations
func loadChildOrg(parent, child int) (testOrg, error) {

	var org = testOrg{}
	// load the CA
	caPEM, err := ioutil.ReadFile(fmt.Sprintf(childCACert, parent, child))
	if err != nil {
		return org, err
	}
	// loop through and load servers
	var serverCerts = []serverCert{}
	for i := 1; i <= numServerCerts; i++ {
		keyPEM, err := ioutil.ReadFile(fmt.Sprintf(childServerKey, parent, child, i))
		if err != nil {
			return org, err
		}
		certPEM, err := ioutil.ReadFile(fmt.Sprintf(childServerCert, parent, child, i))
		if err != nil {
			return org, err
		}
		serverCerts = append(serverCerts, serverCert{keyPEM, certPEM})
	}
	// loop through and load clients
	var clientCerts = []tls.Certificate{}
	for j := 1; j <= numServerCerts; j++ {
		clientCert, err := loadTLSKeyPairFromFile(fmt.Sprintf(childClientKey, parent, child, j),
			fmt.Sprintf(childClientCert, parent, child, j))
		if err != nil {
			return org, err
		}
		clientCerts = append(clientCerts, clientCert)
	}
	return testOrg{caPEM, serverCerts, clientCerts, []testOrg{}}, nil
}


// loadTLSKeyPairFromFile creates a tls.Certificate from PEM-encoded key and cert files
func loadTLSKeyPairFromFile(keyFile, certFile string) (tls.Certificate, error) {

	certPEMBlock, err := ioutil.ReadFile(certFile)
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)

	if err != nil {
		return tls.Certificate{}, err
	}
	return cert, nil
}
// test servers to be registered with the GRPCServer
type emptyServiceServer struct{}

func (ess *emptyServiceServer) EmptyCall(context.Context, *testpb.Empty) (*testpb.Empty, error) {
	return new(testpb.Empty), nil
}

func (esss *emptyServiceServer) EmptyStream(stream testpb.EmptyService_EmptyStreamServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&testpb.Empty{}); err != nil {
			return err
		}

	}
}

const (
	numOrgs        = 2
	numChildOrgs   = 2
	numClientCerts = 2
	numServerCerts = 2
)

// string for cert filenames
var (
	orgCAKey        = filepath.Join("testdata", "certs", "Org%d-key.pem")
	orgCACert       = filepath.Join("testdata", "certs", "Org%d-cert.pem")
	orgServerKey    = filepath.Join("testdata", "certs", "Org%d-server%d-key.pem")
	orgServerCert   = filepath.Join("testdata", "certs", "Org%d-server%d-cert.pem")
	orgClientKey    = filepath.Join("testdata", "certs", "Org%d-client%d-key.pem")
	orgClientCert   = filepath.Join("testdata", "certs", "Org%d-client%d-cert.pem")
	childCAKey      = filepath.Join("testdata", "certs", "Org%d-child%d-key.pem")
	childCACert     = filepath.Join("testdata", "certs", "Org%d-child%d-cert.pem")
	childServerKey  = filepath.Join("testdata", "certs", "Org%d-child%d-server%d-key.pem")
	childServerCert = filepath.Join("testdata", "certs", "Org%d-child%d-server%d-cert.pem")
	childClientKey  = filepath.Join("testdata", "certs", "Org%d-child%d-client%d-key.pem")
	childClientCert = filepath.Join("testdata", "certs", "Org%d-child%d-client%d-cert.pem")
)

type testServer struct {
	config comm.ServerConfig
}

type serverCert struct {
	keyPEM  []byte
	certPEM []byte
}

type testOrg struct {
	rootCA      []byte
	serverCerts []serverCert
	clientCerts []tls.Certificate
	childOrgs   []testOrg
}
func TestNewGRPCServerInvalidParameters(t *testing.T) {

	t.Parallel()
	// missing address
	_, err := comm.NewGRPCServer("", comm.ServerConfig{
		SecOpts: &comm.SecureOptions{UseTLS: false}})
	// check for error
	msg := "Missing address parameter"
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}

	// missing port
	_, err = comm.NewGRPCServer("abcdef", comm.ServerConfig{
		SecOpts: &comm.SecureOptions{UseTLS: false}})
	// check for error
	assert.Error(t, err, "Expected error with missing port")
	msg = "missing port in address"
	assert.Contains(t, err.Error(), msg)

	// bad port
	_, err = comm.NewGRPCServer("localhost:1BBB", comm.ServerConfig{
		SecOpts: &comm.SecureOptions{UseTLS: false}})
	//check for possible errors based on platform and Go release
	msgs := []string{
		"listen tcp: lookup tcp/1BBB: nodename nor servname provided, or not known",
		"listen tcp: unknown port tcp/1BBB",
		"listen tcp: address tcp/1BBB: unknown port",
		"listen tcp: lookup tcp/1BBB: Servname not supported for ai_socktype",
	}

	if assert.Error(t, err, fmt.Sprintf("[%s], [%s] [%s] or [%s] expected", msgs[0], msgs[1], msgs[2], msgs[3])) {
		assert.Contains(t, msgs, err.Error())
	}
	if err != nil {
		t.Log(err.Error())
	}

	// bad hostname
	_, err = comm.NewGRPCServer("hostdoesnotexist.localdomain:9050",
		comm.ServerConfig{SecOpts: &comm.SecureOptions{UseTLS: false}})
	/*
		We cannot check for a specific error message due to the fact that some
		systems will automatically resolve unknown host names to a "search"
		address so we just check to make sure that an error was returned
	*/
	assert.Error(t, err, fmt.Sprintf("%s error expected", msg))
	if err != nil {
		t.Log(err.Error())
	}

	// address in use
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to create listener [%s]", err)
	}
	defer lis.Close()
	_, err = comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{UseTLS: false}})
	if err != nil {
		t.Fatalf("Failed to create GRPCServer [%s]", err)
	}
	_, err = comm.NewGRPCServer(
		lis.Addr().String(),
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{UseTLS: false}},
	)
	// check for error
	if err != nil {
		t.Log(err.Error())
	}
	assert.Contains(t, err.Error(), "address already in use")

	// missing server Certificate
	_, err = comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				UseTLS: true,
				Key:    []byte{}},
		},
	)
	// check for error
	msg = "serverConfig.SecOpts must contain both Key and " +
		"Certificate when UseTLS is true"
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}

	// missing server Key
	_, err = comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				UseTLS:      true,
				Certificate: []byte{}},
		},
	)
	// check for error
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}

	// bad server Key
	_, err = comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				UseTLS:      true,
				Certificate: []byte(selfSignedCertPEM),
				Key:         []byte{}},
		},
	)

	// check for error
	msg = "tls: failed to find any PEM data in key input"
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}

	// bad server Certificate
	_, err = comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				UseTLS:      true,
				Certificate: []byte{},
				Key:         []byte(selfSignedKeyPEM)},
		},
	)
	//check for error
	msg = "tls: failed to find any PEM data in certificate input"
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}

	srv, err := comm.NewGRPCServerFromListener(
		lis,
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				UseTLS:            true,
				Certificate:       []byte(selfSignedCertPEM),
				Key:               []byte(selfSignedKeyPEM),
				RequireClientCert: true},
		},
	)
	badRootCAs := [][]byte{[]byte(badPEM)}
	err = srv.SetClientRootCAs(badRootCAs)
	// check for error
	msg = "Failed to set client root certificate(s): " +
		"asn1: syntax error: data truncated"
	assert.EqualError(t, err, msg)
	if err != nil {
		t.Log(err.Error())
	}
}


func TestNewGRPCServer(t *testing.T) {

	t.Parallel()
	testAddress := "localhost:9053"
	srv, err := comm.NewGRPCServer(
		testAddress,
		comm.ServerConfig{SecOpts: &comm.SecureOptions{UseTLS: false}},
	)
	//check for error
	if err != nil {
		t.Fatalf("Failed to return new GRPC server: %v", err)
	}

	// make sure our properties are as expected
	// resolve the address
	addr, err := net.ResolveTCPAddr("tcp", testAddress)
	assert.Equal(t, srv.Address(), addr.String())
	assert.Equal(t, srv.Listener().Addr().String(), addr.String())

	// TLSEnabled should be false
	assert.Equal(t, srv.TLSEnabled(), false)
	// MutualTLSRequired should be false
	assert.Equal(t, srv.MutualTLSRequired(), false)

	// register the GRPC test server
	testpb.RegisterEmptyServiceServer(srv.Server(), &emptyServiceServer{})

	// start the server
	go srv.Start()

	defer srv.Stop()
	// should not be needed
	time.Sleep(10 * time.Millisecond)

	// GRPC client options
	var dialOptions []grpc.DialOption
	dialOptions = append(dialOptions, grpc.WithInsecure())

	// invoke the EmptyCall service
	_, err = invokeEmptyCall(testAddress, dialOptions)

	if err != nil {
		t.Fatalf("GRPC client failed to invoke the EmptyCall service on %s: %v",
			testAddress, err)
	} else {
		t.Log("GRPC client successfully invoked the EmptyCall service: " + testAddress)
	}
}
