package cli

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

const (
	numbers                    string = "0123456789"
	lowercase                  string = "abcdefghijklmnopqrstuvwxyz"
	min_domain_length          int    = 1
	min_master_password_length int    = 8
)

var (
	uppercase string = strings.ToUpper(lowercase)
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

// validateDomain checks to make sure the user has entered a domain which is at
// least one character long. The domain is not expected to be, e.g., a fqdn, but rather
// a kind of context, which is easy to remember for the user.
func validateDomain(domain string) bool {
	switch {
	case len(domain) < min_domain_length:
		fmt.Printf(
			"Domain needs to be at least %d character long.\n",
			min_domain_length,
		)
		return false
	default:
		return true
	}
}

// containsSymbol checks whether the input string contains at least one symbol.
// By 'symbol' here, we mean it contains at least one character which is not a
// number, not a lowercase letter and not an uppercase letter. This is really
// not general because, e.g., caseable non-latin letters will be considered 'symbols'
// here but who cares? If the user is using characters exterior to the ascii subset
// it has the same effect on password entropy.
func containsSymbol(target string) bool {
	var (
		non_symbol_characters string = lowercase + uppercase + numbers
		non_symbol_count      int    = 0
	)

	for _, character := range non_symbol_characters {
		non_symbol_count += strings.Count(target, string(character))
	}

	if non_symbol_count < len(target) {
		return true
	} else {
		return false
	}
}

// validateMasterPassword checks to make sure the user is entering a strong-enough
// master password. Will return false if the password has length less than
// min_master_password_length, no lowercase characters, no uppercase characters or
// no numbers. Else, returns true.
func validateMasterPassword(master_password string) bool {
	switch {
	case len(master_password) < min_master_password_length:
		fmt.Printf(
			"Master password needs to be at least %d characters long.\n",
			min_master_password_length,
		)
		return false
	case !strings.ContainsAny(master_password, lowercase):
		fmt.Println("Master password must contain at least one lowercase letter.")
		return false
	case !strings.ContainsAny(master_password, uppercase):
		fmt.Println("Master password must contain at least one uppercase letter.")
		return false
	case !strings.ContainsAny(master_password, numbers):
		fmt.Println("Master password must contain at least one number.")
		return false
	case !containsSymbol(master_password):
		fmt.Println("Master password must contain at least one symbol.")
		return false
	default:
		return true

	}
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
