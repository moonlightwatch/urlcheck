package main

import (
	"crypto/tls"
	"net/http"
	"strings"
	"time"
)

func TlsSupport(target string) []string {
	versions := []uint16{}
	{
		clientTLS10 := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					MinVersion:         tls.VersionTLS10,
					MaxVersion:         tls.VersionTLS10,
				},
			},
		}

		_, err := clientTLS10.Get(target)
		if err == nil {
			versions = append(versions, tls.VersionTLS10)
		}

	}

	{
		clientTLS11 := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					MinVersion:         tls.VersionTLS11,
					MaxVersion:         tls.VersionTLS11,
				},
			},
		}

		_, err := clientTLS11.Get(target)
		if err == nil {
			versions = append(versions, tls.VersionTLS11)
		}

	}

	{
		clientTLS12 := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					MinVersion:         tls.VersionTLS12,
					MaxVersion:         tls.VersionTLS12,
				},
			},
		}

		_, err := clientTLS12.Get(target)
		if err == nil {
			versions = append(versions, tls.VersionTLS12)
		}
	}

	{
		clientTLS13 := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					MinVersion:         tls.VersionTLS13,
					MaxVersion:         tls.VersionTLS13,
				},
			},
		}

		_, err := clientTLS13.Get(target)
		if err == nil {
			versions = append(versions, tls.VersionTLS13)
		}
	}
	result := []string{}

	for _, version := range versions {
		result = append(result, tls.VersionName(version))
	}

	return result
}

func GetCertificate(target string) (string, string, time.Time) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(target)
	if err != nil {
		return "", "", time.Time{}
	}
	return resp.TLS.PeerCertificates[0].Issuer.String(), strings.Join(resp.TLS.PeerCertificates[0].DNSNames, "|"), resp.TLS.PeerCertificates[0].NotAfter
}
