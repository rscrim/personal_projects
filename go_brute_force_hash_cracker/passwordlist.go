/* DESCRIPTION
This package manages all operations related to maintaining the passwords.list file.
*/

package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"unicode/utf8"
)

// This function downloads the contents of a URL and saves it to a file, overwriting any existing content in the file.
// It returns an error if the download or write operation fails.
func updateWordlistFromURL() {
	// Send an HTTP GET request to the URL
	resp, err := http.Get("https://raw.githubusercontent.com/jeanphorn/wordlist/master/passlist.txt")
	if err != nil {
		log.Println("Error: get request failed.")
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Create or overwrite a file named "passwords.list" in the current directory
	file, err := os.Create("passwords.list")
	if err != nil {
		log.Println("Error: failed to create file.")
		log.Println(err)
		return
	}
	defer file.Close()

	// Create a scanner to read the contents of the HTTP response body
	scanner := bufio.NewScanner(resp.Body)

	// Scan the contents of the response body line by line and write the filtered output to the file
	for scanner.Scan() {
		line := scanner.Text()
		// Filter out any non-ASCII characters
		if utf8.ValidString(line) {
			_, err = file.WriteString(line + "\n")
			if err != nil {
				log.Println("Error: invalid characters detected.")
				log.Println(err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error: unable to scan response body.")
		log.Println(err)
		return
	}
}
