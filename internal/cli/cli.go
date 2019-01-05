package cli

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func getDomain() string {
	var (
		reader *bufio.Reader
		domain string
		err    error
	)
	reader = bufio.NewReader(os.Stdin)

	fmt.Println("Enter domain:")
	domain, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read domain: %v", err)
	}
	domain = strings.Trim(domain, "\n")

	return domain
}

func getMasterPassword() string {
	var (
		byte_password   []uint8
		err             error
		master_password string
	)

	fmt.Println("Enter master password:")
	byte_password, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Failed to read master password: %v", err)
	}

	master_password = string(byte_password)
	// Add a new line here to make the UI consistent with getDomain()
	fmt.Println()

	return master_password
}

func validateDomain(domain string) bool {
	switch {
	case len(domain) < 1:
		fmt.Println("Domain needs to be at least one character long.")
		return false
	default:
		return true
	}
}

func validateMasterPassword(master_password string) bool {
	return true
}

func GetSecrets() (string, string) {
	var (
		master_password       string
		domain                string
		domain_valid          bool = false
		master_password_valid bool = false
	)

	for !domain_valid {
		domain = getDomain()
		domain_valid = validateDomain(domain)
	}
	for !master_password_valid {
		master_password = getMasterPassword()
		master_password_valid = validateMasterPassword(master_password)
	}

	return master_password, domain

}
