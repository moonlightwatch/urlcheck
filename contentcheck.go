package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func ContentCheck(target string, hosts []string) []string {
	results := []string{}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	for _, host := range hosts {
		req, err := http.NewRequest(http.MethodGet, target, nil)
		if err != nil {
			continue
		}
		req.Header.Add("Host", req.Host)

		req.URL.Host = host
		start := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			results = append(results, fmt.Sprintf("Request ip %s, error: %+v", host, err))
			continue
		}
		content := make([]byte, 100)
		io.ReadFull(resp.Body, content)
		io.ReadAll(resp.Body)
		defer resp.Body.Close()
		results = append(results, fmt.Sprintf("Request ip %s, status %d, %+v, content: %s", host, resp.StatusCode, time.Since(start), strings.ReplaceAll(string(content), "\n", "")))
	}
	return results
}
