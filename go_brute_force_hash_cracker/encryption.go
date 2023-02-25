/*
	DESCRIPTION:

This package manages all functions related to encrypting and decrypting the passwords.list file
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

// EncryptFile encrypts the contents of the input file and writes the encrypted data to the output file
func EncryptFile(inputFile string, outputFile string, key []byte) error {
	// Open the input file for reading
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("error opening input file: %s", err)
	}
	defer inFile.Close()

	// Create the output file for writing
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %s", err)
	}
	defer outFile.Close()

	// Generate a new AES cipher block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating AES cipher: %s", err)
	}

	// Create a new stream cipher using AES in counter mode (CTR)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// Create a new writer that encrypts data using the stream cipher and writes the encrypted data to the output file
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	// Copy the contents of the input file to the output file, encrypting the data as it is written
	if _, err := io.Copy(writer, inFile); err != nil {
		return fmt.Errorf("error encrypting file: %s", err)
	}

	return nil
}

// DecryptFile decrypts the contents of the input file and writes the decrypted data to the output file
func DecryptFile(inputFile string, outputFile string, key []byte) error {
	// Open the input file for reading
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("error opening input file: %s", err)
	}
	defer inFile.Close()

	// Create the output file for writing
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %s", err)
	}
	defer outFile.Close()

	// Generate a new AES cipher block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating AES cipher: %s", err)
	}

	// Create a new stream cipher using AES in counter mode (CTR)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// Create a new reader that reads data from the input file and decrypts it using the stream cipher
	reader := &cipher.StreamReader{S: stream, R: inFile}

	// Copy the decrypted data from the input file to the output file
	if _, err := io.Copy(outFile, reader); err != nil {
		return fmt.Errorf("error decrypting file: %s", err)
	}

	return nil
}
