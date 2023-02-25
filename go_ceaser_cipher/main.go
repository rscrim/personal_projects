/* DESCRIPTION
Package caesar provides functions for encrypting and decrypting text using a Caesar cipher.

The Caesar cipher is a simple substitution cipher that shifts each letter of the plaintext by
a fixed number of positions in the alphabet. For example, with a shift of 3, the letter 'A'
would be replaced by 'D', the letter 'B' would become 'E', and so on.

This application provides two functions, caesarEncrypt() and caesarDecrypt(), that take a string
and a shift value as input, and return the corresponding ciphertext or plaintext.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Encrypt 'plaintext' using Ceaser Cypher with key of 'shift'
func caesarEncrypt(plaintext string, shift int) string {
	var ciphertext strings.Builder
	for _, c := range plaintext {
		if c >= 'A' && c <= 'Z' {
			ciphertext.WriteRune('A' + (c-'A'+rune(shift))%26)
		} else if c >= 'a' && c <= 'z' {
			ciphertext.WriteRune('a' + (c-'a'+rune(shift))%26)
		} else {
			ciphertext.WriteRune(c)
		}
	}
	return ciphertext.String()
}

// Decrypt 'plaintext' using Ceaser Cypher with key of 'shift'
func caesarDecrypt(ciphertext string, shift int) string {
	var plaintext strings.Builder
	for _, c := range ciphertext {
		if c >= 'A' && c <= 'Z' {
			plaintext.WriteRune('A' + (c-'A'-rune(shift)+26)%26)
		} else if c >= 'a' && c <= 'z' {
			plaintext.WriteRune('a' + (c-'a'-rune(shift)+26)%26)
		} else {
			plaintext.WriteRune(c)
		}
	}
	return plaintext.String()
}

// Convert a string to leet (1337) speak
func l33t5peak(plaintext string) string {
	leetMap := map[rune]string{
		'A': "4",
		'E': "3",
		'G': "6",
		'I': "1",
		'O': "0",
		'S': "5",
		'T': "7",
	}
	leetText := ""
	for _, char := range strings.ToUpper(plaintext) {
		if replacement, ok := leetMap[char]; ok {
			leetText += replacement
		} else {
			leetText += string(char)
		}
	}
	return leetText
}

// Logs all errors that occur
func handleError(err error) {
	if err != nil {
		fmt.Println("An error occurred.")
		log.Fatal(err)
	}
}

// Handle user input
func getUserInput(note string) string {
	fmt.Println(note)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	handleError(err)
	return strings.TrimSpace(input)
}

// Prints functionality menu
func mainMenu() int {
	fmt.Println(strings.Repeat("-", 10))
	fmt.Printf("\n" +
		l33t5peak("MAIN MENU") + "\n")
	fmt.Println(strings.Repeat("-", 10))
	for {
		fmt.Printf(
			"\n" +
				"1. Encrypt a new string\n" +
				"2. Decrypt an existing string\n" +
				"3. Exit\n")
		userInput := getUserInput("\nWhat would you like to do?")
		choice, err := strconv.Atoi(userInput)
		if choice == 1 {
			return 1
		} else if choice == 2 {
			return 2
		} else if choice == 3 {
			os.Exit(0)
		} else if err != nil {
			fmt.Println("Invalid select.")
			fmt.Println("Please select one of the following options using 1, 2 or 3.")
			continue
		}
	}
}

func getShiftValue(note string) int {
	fmt.Println(note)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		handleError(err)
		input = strings.TrimSpace(input)
		shift, err := strconv.Atoi(input)
		if err == nil {
			return shift
		}
		fmt.Println("Invalid shift value. Please try again with a valid number.")
	}
}

func main() {
	fmt.Println("Welcome to your " + l33t5peak("Ceaser Cipher") + "\nSelect what you want to do:")
	for {
		choice := mainMenu()
		if choice == 1 {
			plaintext := getUserInput("Enter string to encrypt")
			shift := getShiftValue("Enter a number to shift values by")
			fmt.Println("Your encrypted phrase is: ")
			fmt.Println(caesarEncrypt(plaintext, shift))
		} else if choice == 2 {
			ciphertext := getUserInput("Enter string to decrypt")
			shift := getShiftValue("Enter the number the values were encrypted with")
			fmt.Println("Your decrypted phrase is: ")
			fmt.Println(caesarDecrypt(ciphertext, shift))
		}
	}
}
