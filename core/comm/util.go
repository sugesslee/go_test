package comm

import (
	"crypto/x509"
	"encoding/pem"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/01     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/1 11:00 AM
 * @date 2019/11/1 11:00 AM
 * @since 1.0.0
 */
// AddPemToCertPool adds PEM-encoded certs to a cert pool
func AddPemToCertPool(pemCerts []byte, pool *x509.CertPool) error {
	certs, _, err := pemToX509Certs(pemCerts)
	if err != nil {
		return err
	}
	for _, cert := range certs {
		pool.AddCert(cert)
	}
	return nil
}

//utility function to parse PEM-encoded certs
func pemToX509Certs(pemCerts []byte) ([]*x509.Certificate, []string, error) {
	//it's possible that multiple certs are encoded
	var certs []*x509.Certificate
	var subjects []string

	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)

		if block == nil {
			break
		}
		/** TODO: check why msp does not add type to PEM header
		if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
			continue
		}
		*/
		certificate, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, subjects, err
		} else {
			certs = append(certs, certificate)
			// extract and append the subject
			subjects = append(subjects, string(certificate.RawSubject))
		}
	}
	return certs, subjects, nil
}
