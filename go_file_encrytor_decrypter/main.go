/*
File Encryption/Decryption Program

This program encrypts and decrypts files using AES encryption in Galois/Counter
Mode (GCM) with a custom password. It supports multiple encryption algorithms,
including aes-128-gcm, aes-192-gcm, and aes-256-gcm. The program provides a
command-line interface for specifying the file to encrypt/decrypt, the
encryption algorithm, and the password. It uses the crypto/rand package to
generate a cryptographically secure nonce for each encrypted file.

Author: Ryan Scrim
Date: 2023-03-06
Version: 1.2
*/

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	filePath     string
	password     string
	algorithm    string
	decrypt      bool
	algorithms   = []string{"aes-128-gcm", "aes-192-gcm", "aes-256-gcm"}
	algorithmMap = map[string]int{"aes-128-gcm": 16, "aes-192-gcm": 24, "aes-256-gcm": 32}
)

// main parses the command-line arguments and either encrypts or decrypts a file.
func main() {
	flag.StringVar(&filePath, "file", "", "path to file to encrypt/decrypt")
	flag.StringVar(&password, "password", "", "password used to encrypt/decrypt file")
	flag.StringVar(&algorithm, "algorithm", "aes-256-gcm", "encryption algorithm to use (aes-128-gcm, aes-192-gcm, or aes-256-gcm)")
	flag.BoolVar(&decrypt, "decrypt", false, "decrypt file")
	flag.Parse()

	if filePath == "" {
		log.Fatal("File path is required")
	}
	if password == "" {
		log.Fatal("Password is required")
	}
	if !contains(algorithms, algorithm) {
		log.Fatal("Invalid encryption algorithm. Supported algorithms: aes-128-gcm, aes-192-gcm, aes-256-gcm")
	}

	if decrypt {
		err := decryptFile(filePath, password, algorithm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Decrypted file: %s\n", filePath)
	} else {
		err := encryptFile(filePath, password, algorithm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Encrypted file: %s\n", filePath+".enc")
	}
}

// contains checks whether a string is present in a slice of strings.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// encryptFile encrypts a file using AES-GCM encryption.
func encryptFile(filePath string, password string, algorithm string) error {
	key, err := generateKey(password, algorithmMap[algorithm])
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Create a buffer to hold the nonce and encrypted data.
	ciphertext := make([]byte, 0, len(nonce)+len(algorithm)+len(password))

	// Write the nonce to the buffer.
	ciphertext = append(ciphertext, nonce...)

	// Write the encryption algorithm to the buffer.
	ciphertext = append(ciphertext, []byte(algorithm)...)

	// Create a stream cipher for encrypting the file.
	stream := cipher.NewCTR(block, nonce)

	// Create a StreamWriter that encrypts data using the stream cipher.
	writer := &cipher.StreamWriter{S: stream, W: &bytes.Buffer{}}

	// Copy the data from the input file to the StreamWriter.
	if _, err = io.Copy(writer, f); err != nil {
		return err
	}

	// Write the encrypted data to the buffer.
	ciphertext = append(ciphertext, writer.W.(*bytes.Buffer).Bytes()...)

	// Write the encrypted data to the output file.
	if err = ioutil.WriteFile(filePath+".enc", ciphertext, 0644); err != nil {
		return err
	}

	return nil
}

// decryptFile decrypts a file that was encrypted with AES-GCM encryption.
func decryptFile(filePath string, password string, algorithm string) error {
	key, err := generateKey(password, algorithmMap[algorithm])
	if err != nil {
		return err
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filePath[:len(filePath)-4], plaintext, 0644); err != nil {
		return err
	}

	return nil
}

// generateKey generates a key for AES encryption using the provided password.
func generateKey(password string, keySize int) ([]byte, error) {
	hash := sha256.Sum256([]byte(password))
	key := make([]byte, keySize)
	copy(key, hash[:keySize])
	return key, nil
}
