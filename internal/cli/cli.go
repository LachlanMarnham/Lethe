package cli

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func GetSecrets() (string, string) {
	var master_password string
	var domain string
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Enter domain: ")
	domain, _ = reader.ReadString('\n')

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Failed to read password: %v", err)
	}
	master_password = string(bytePassword)
	fmt.Println()

	return master_password, domain

}
