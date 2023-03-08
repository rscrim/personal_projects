/* DESCRIPTION:
This program uses a pre-defined dictionary to perform an offline attack, brute forcing a hash crack.
Note: This does not handle salted hashes.

It supports multiple hashing algorithms including:
- MD5
- SHA1
- SHA256

The default dictionary is included in this directory, but a custom one can be specified during run.

Author: [Your Name]
Date: [Current Date]
Version: 1.0
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	keyLength = 2048 // The length of the RSA key in bits. Must be a multiple of 8.
)

// generateKeyPair generates an RSA key pair with the specified length.
// It returns the public and private keys.
func generateKeyPair(length int) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %s", err)
	}
	return &privateKey.PublicKey, privateKey, nil
}

// encodePrivateKey encodes an RSA private key to PEM format.
// The resulting PEM block has a "RSA PRIVATE KEY" header.
func encodePrivateKey(privateKey *rsa.PrivateKey) ([]byte, error) {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return pem.EncodeToMemory(block), nil
}

// encodePublicKey encodes an RSA public key to PEM format.
// The resulting PEM block has a "PUBLIC KEY" header.
func encodePublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %s", err)
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}
	return pem.EncodeToMemory(block), nil
}

func main() {
	// Generate an RSA key pair.
	publicKey, privateKey, err := generateKeyPair(keyLength)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Encode the private key to PEM format.
	privateKeyBytes, err := encodePrivateKey(privateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Encode the public key to PEM format.
	publicKeyBytes, err := encodePublicKey(publicKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Print the private and public keys.
	fmt.Printf("Private Key:\n%s\n", privateKeyBytes)
	fmt.Printf("Public Key:\n%s\n", publicKeyBytes)
}
