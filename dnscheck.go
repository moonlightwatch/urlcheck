package main

import (
	"net"
	"net/url"
	"strings"
)

func ResolveDns(target string) []string {
	u, err := url.Parse(target)
	if err != nil {
		return []string{}
	}
	domain := u.Host
	if domain == "" {
		return []string{}
	}
	if strings.ContainsRune(domain, ':') {
		domain = strings.Split(domain, ":")[0]
	}
	ips, err := net.LookupIP(domain)
	if err == nil {
		result := []string{}
		for _, ip := range ips {
			result = append(result, ip.String())
		}
		return result
	}
	return []string{}
}
