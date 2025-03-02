package samokat

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	utls "github.com/refraction-networking/utls"
)

const (
	dialTimeout = 5 * time.Second
)

var signatureAlgorithms = []utls.SignatureScheme{
	utls.ECDSAWithP256AndSHA256, // ecdsa_secp256r1_sha256 (0x0403)
	utls.ECDSAWithP384AndSHA384, // ecdsa_secp384r1_sha384 (0x0503)
	utls.ECDSAWithP521AndSHA512, // ecdsa_secp521r1_sha512 (0x0603)
	utls.PSSWithSHA256,          // rsa_pss_rsae_sha256 (0x0804)
	utls.PSSWithSHA384,          // rsa_pss_rsae_sha384 (0x0805)
	utls.PSSWithSHA512,          // rsa_pss_rsae_sha512 (0x0806)
	0x0809,
	0x080a,
	0x080b,
	utls.PKCS1WithSHA256, // rsa_pkcs1_sha256 (0x0401)
	utls.PKCS1WithSHA384, // rsa_pkcs1_sha384 (0x0501)
	utls.PKCS1WithSHA512, // rsa_pkcs1_sha512 (0x0601)
	0x0402,
	0x0303,
	0x0301,
	0x0302,
	0x0203,
	0x0201,
	0x0202,
}

func samokatDial(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, dialTimeout)
	if err != nil {
		return nil, fmt.Errorf("tcp dial error: %v", err)
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address %v: %v", addr, err)
	}

	config := &utls.Config{
		ServerName: host,
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
	}

	uTlsConn := utls.UClient(conn, config, utls.HelloCustom)

	spec := utls.ClientHelloSpec{
		TLSVersMin: tls.VersionTLS12,
		TLSVersMax: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,       // 0xc02b
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,       // 0xc02c
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, // 0xcca9
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,         // 0xc02f
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,         // 0xc030
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,   // 0xcca8
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,            // 0xc013
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,            // 0xc014
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,               // 0x009c
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,               // 0x009d
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,                  // 0x002f
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,                  // 0x0035
		},
		Extensions: []utls.TLSExtension{
			&utls.SNIExtension{},
			&utls.StatusRequestExtension{},
			&utls.SupportedCurvesExtension{Curves: []utls.CurveID{utls.X25519, utls.CurveP256, utls.CurveP384, utls.CurveP521, 0x001e}},
			&utls.SupportedPointsExtension{SupportedPoints: []uint8{0x0}},
			&utls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: signatureAlgorithms},
			&utls.SignatureAlgorithmsCertExtension{SupportedSignatureAlgorithms: signatureAlgorithms},
			&utls.ALPNExtension{AlpnProtocols: []string{"h2", "http/1.1"}},
			&utls.StatusRequestV2Extension{},
			&utls.ExtendedMasterSecretExtension{},
			&utls.SupportedVersionsExtension{
				Versions: []uint16{tls.VersionTLS12},
			},
			&utls.RenegotiationInfoExtension{},
		},
	}

	if err := uTlsConn.ApplyPreset(&spec); err != nil {
		return nil, fmt.Errorf("error applying preset: %v", err)
	}

	if err := uTlsConn.Handshake(); err != nil {
		return nil, fmt.Errorf("TLS handshake error: %v", err)
	}

	return uTlsConn, nil
}
