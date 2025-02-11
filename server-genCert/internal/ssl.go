package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/luskaner/ageLANServer/common"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func GenerateCertificatePair(gameId, folder string) bool {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return false
	}
	domain := common.Domain(gameId)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   domain,
			Organization: []string{common.CertSubjectOrganization},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		DNSNames: []string{domain},
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return false
	}

	certFile, err := os.Create(filepath.Join(folder, common.Cert(domain)))
	if err != nil {
		return false
	}
	var keyFile *os.File
	delCertFile := false
	delKeyFile := false
	defer func() {
		_ = certFile.Close()
		if delCertFile {
			_ = os.Remove(filepath.Join(folder, common.Cert(domain)))
		}
		if keyFile != nil {
			_ = keyFile.Close()
			if delKeyFile {
				_ = os.Remove(filepath.Join(folder, common.Key(domain)))
			}
		}
	}()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})

	if err != nil {
		delCertFile = true
		return false
	}

	keyFile, err = os.Create(filepath.Join(folder, common.Key(domain)))

	if err != nil {
		delCertFile = true
		return false
	}

	err = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		delCertFile = true
		delKeyFile = true
		return false
	}

	return true
}
