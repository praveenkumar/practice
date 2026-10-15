package main

import (
	"fmt"
	"strings"
)

func parseProxyCAData() (string, error) {
	rawData := "-----BEGIN CERTIFICATE-----\r\nMIID6zCCAtOgAwIB\r\nKMGinNDOEwKTFXf\r\n-----END CERTIFICATE-----\r\n\r\n"
	return trimTrailingEOL(rawData), nil
}

func trimTrailingEOL(s string) string {
	fmt.Printf("s: %#v", s)
	s1 := strings.TrimRight(s, "\n\r")
	fmt.Println()
	fmt.Printf("s1: %#v", s1)
	return s1
}

func main() {
	_, err := parseProxyCAData()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}
