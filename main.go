package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	target := os.Args[1]

	fmt.Println("--- Checking HTTP delay ---")
	delay, status := HTTPDelyCheck(target)
	fmt.Println("🕒", delay)

	if delay == 0 {
		fmt.Println("🔥", "Target is not reachable")
		return
	} else {
		fmt.Println("🔥", status)
	}
	fmt.Println()

	fmt.Println("--- Resolving DNS ---")
	ips := ResolveDns(target)
	for _, ip := range ips {
		fmt.Println("🚩", ip)
	}
	fmt.Println()

	fmt.Println("--- Checking HTTP version support ---")
	httpVersions := HTTPVersionSupport(target)
	for _, version := range httpVersions {
		fmt.Println("🌞", version)
	}
	fmt.Println()

	if strings.HasPrefix(target, "https://") {
		fmt.Println("--- Checking TLS support ---")
		tlsVersions := TlsSupport(target)
		for _, version := range tlsVersions {
			fmt.Println("✅", version)
		}
		fmt.Println()

		fmt.Println("--- Checking certificate ---")
		issuer, names, notAfter := GetCertificate(target)
		fmt.Println("🔰", issuer)
		fmt.Println("🌐", names)
		fmt.Println("🕒", notAfter.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}

	log.Println("Done")
}
