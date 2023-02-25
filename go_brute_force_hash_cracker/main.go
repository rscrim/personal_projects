/* DESCRIPTION:
This program uses a pre-defined dictionary to perform an offline attack, brute forcing a hash crack.
Note: This does not handle salted hashes.

It supports multiple hashing algorithms including:
- MD5
- SHA1
- SHA256
The default dictionary is included in this directory, but a custom one can be specified during run.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func menu() string {
	fmt.Println("Please choose a hash type:")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA-256")
	fmt.Println("3. SHA-512")
	fmt.Println("0. Quit")

	var choice string
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(input)

	switch choice {
	case "1":
		return "md5"
	case "2":
		return "sha256"
	case "3":
		return "sha512"
	case "0":
		os.Exit(0)
	default:
		log.Println("Invalid choice.")
		return menu()
	}
	return ""
}

func importDictionary() {
	return
}

func bruteForceMD5(targetHash string) string {
	return ""
}

func bruteForceSHA1(targetHash string) string {
	return ""
}

func bruteForceSHA256(targetHash string) string {
	return ""
}

func main() {
	fmt.Println("Welcome to the Go Brute Force Hash Cracker!")
	fmt.Println("Please wait while we download the latest wordlist...")
	updateWordlistFromURL()
	selection := menu()
	fmt.Println("You chose: " + selection)
}
