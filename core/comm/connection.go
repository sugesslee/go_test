package comm

import (
	"crypto/tls"
	"gostudy/common/flogging"
	"sync"
	"time"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/01     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/1 11:16 AM
 * @date 2019/11/1 11:16 AM
 * @since 1.0.0
 */
const defaultTimeOut = time.Second * 3

var commLogger = flogging.MustGetLogger("comm")
var credLogger *CredentialSupport
var once sync.Once

// CertificateBundle bundles certificates
type CertificateBundle [][]byte

// PerOrgCertificateBundle maps organizations to CertificateBundles
type PerOrgCertificateBundle map[string]CertificateBundle


// OrgRootCAs defines root CA certificates of organizations, by their
// corresponding channels.
// channel --> organization --> certificates
type OrgRootCAs map[string]PerOrgCertificateBundle


// CredentialSupport type manages credentials used for gRPC client connections
type CredentialSupport struct {
	sync.RWMutex
	AppRootCAsByChain           map[string]CertificateBundle
	OrdererRootCAsByChainAndOrg OrgRootCAs
	ClientRootCAs               CertificateBundle
	ServerRootCAs               CertificateBundle
	clientCert                  tls.Certificate
}
