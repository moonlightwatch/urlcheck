package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func HTTPDelyCheck(target string) (time.Duration, int) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	start := time.Now()
	resp, err := client.Get(target)
	if err == nil {
		resp.Body.Close()
		return time.Since(start), resp.StatusCode
	}
	return 0, 0
}

func HTTPVersionSupport(target string) []string {
	versions := []string{}
	{
		client := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				ForceAttemptHTTP2: true,
			},
		}
		resp, err := client.Get(target)
		if err == nil {
			versions = append(versions, resp.Proto)
		}
	}
	{
		client := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				ForceAttemptHTTP2: false,
			},
		}
		resp, err := client.Get(target)
		if err == nil {
			versions = append(versions, resp.Proto)
		}
	}

	for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
		versions[i], versions[j] = versions[j], versions[i]
	}

	return versions
}
