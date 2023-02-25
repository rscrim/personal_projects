package main

import "fmt"

const programName = "AegisVault"

func main() {
	fmt.Println("Welcome to " + programName)
	while running; {
		running, err := Menu()
		if err != null {
			fmt.Prinltn("An error occurred, saving any changes and exiting.")
			Save()
			os.Exit(1)
		}
		if running != True {
			fmt.Println("Saving, encyrpting and closing.")
			Save()
			Encrypt()
			os.Exit(0)
		}
	}
}
