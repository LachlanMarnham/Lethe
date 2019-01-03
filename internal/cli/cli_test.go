package cli

import (
	"testing"
	"bufio"
	"golang.org/x/crypto/ssh/terminal"
	"bou.ke/monkey"
	"log"
)

type password_parametrization struct {
	as_string string
}

func (p password_parametrization) as_uint8_array() []uint8 {
    return []uint8(p.as_string)
}

func TestGetMasterPassword(t *testing.T) {
	parametrizations := []password_parametrization{
		password_parametrization{"myPassword"},  // ascii Latin characters
		password_parametrization{"myPassword124"},  // ascii numbers
		password_parametrization{"myPassword1!!¬&*"},  // ascii symbols
		password_parametrization{"这是我的密码"},  // non-ascii (Chinese) characters
		password_parametrization{"этомойпароль"},  // non-ascii (Russian) characters
	}

	for i, test_case := range parametrizations {
		// Monkeypatch terminal input with content converted to uint8 array 
		newReadPassword := func(_ int) ([]uint8, error) {
			return test_case.as_uint8_array(), nil
		}
		monkey.Patch(terminal.ReadPassword, newReadPassword)
		defer monkey.Unpatch(terminal.ReadPassword)  // clean-up

		master_password := getMasterPassword()
		if master_password != test_case.as_string {
			t.Errorf(
				"Unexpected master password. Expected: %s. Got: %s.", 
				test_case.as_string, 
				master_password,
			)
		}
		log.Print(i)
	}
}

func TestGetDomain(t *testing.T) {
	var content string = "google.com"

	// Monkeypatch stdin with content converted to uint8 string 
	newReadString := func(_ *bufio.Reader, _ uint8) (string, error) {
		return content + "\n", nil
	}

	monkey.Patch((*bufio.Reader).ReadString, newReadString)
	defer monkey.Unpatch((*bufio.Reader).ReadString)  // clean-up

	domain := getDomain()
	if domain != content {
		t.Errorf("Unexpected domain. Expected: %s. Got: %s.", content, domain)
	}
}
