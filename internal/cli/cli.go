package cli

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

func getDomain() string {
	reader := bufio.NewReader(os.Stdin)
	var domain string
	fmt.Println("Enter domain:")
	domain, _ = reader.ReadString('\n')
	return domain
}

func GetSecrets() (string, string) {
	var master_password string
	var domain string

	domain = getDomain()
	fmt.Println("Enter master password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Failed to read password: %v", err)
	}
	master_password = string(bytePassword)
	fmt.Println()

	return master_password, domain

}
