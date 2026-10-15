// Generate a cryptographically secure JWT secret key.
//
// Usage:
//
//	go run main.go [bytes]
//
// bytes defaults to 64 (512 bits, suitable for HS512).
// Use 32 for HS256, 48 for HS384.
//
// Prefer a browser-based generator instead? https://jwtsecrets.com/
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	size := 64
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid byte count: %s\n", os.Args[1])
			os.Exit(1)
		}
		size = n
	}

	if size < 32 {
		fmt.Fprintln(os.Stderr, "Refusing to generate a secret shorter than 32 bytes (256 bits).")
		os.Exit(1)
	}

	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read random bytes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(base64.RawURLEncoding.EncodeToString(buf))
}
