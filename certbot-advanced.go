package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Define subcommands
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)

	// Determine the subcommand or run default behavior
	if len(os.Args) < 2 {
		handleList()
		return
	}

	switch os.Args[1] {
	case "register":
		handleRegister(registerCmd)
	case "login":
		handleLogin(loginCmd)
	default:
		fmt.Println("Unknown subcommand. Available commands: register, login")
	}
}

// handleList executes the "list" command (default functionality)
func handleList() {
	// Execute the certbot certificates command
	cmd := exec.Command("certbot", "certificates")
	fmt.Print(cmd)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running certbot command: %v\n", err)
		return
	}

	// Use a scanner to process the command output line by line
	scanner := bufio.NewScanner(bytes.NewReader(output))
	fmt.Println("Domains with certificates:")

	for scanner.Scan() {
		line := scanner.Text()
		// Filter lines containing "Certificate Name:"
		if strings.Contains(line, "Certificate Name:") {
			fmt.Print(line)
			// Extract and print the domain name
			fmt.Println(strings.TrimSpace(strings.Replace(line, "Certificate Name:", "", 1)))
		}
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command output: %v\n", err)
	}
}

// handleRegister handles the 'register' subcommand
func handleRegister(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])
	// Placeholder for registration logic
	fmt.Println("Register command: Placeholder for future implementation")
}

// handleLogin handles the 'login' subcommand
func handleLogin(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])
	// Placeholder for login logic
	fmt.Println("Login command: Placeholder for future implementation")
}
