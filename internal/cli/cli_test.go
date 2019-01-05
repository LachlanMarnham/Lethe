package cli

import (
	"bou.ke/monkey"
	"bufio"
	"golang.org/x/crypto/ssh/terminal"
	"testing"
)

type password_parametrization struct {
	as_string string
}

func (p password_parametrization) as_uint8_array() []uint8 {
	return []uint8(p.as_string)
}

func TestGetMasterPassword(t *testing.T) {
	parametrizations := []password_parametrization{
		password_parametrization{"myPassword"},       // ascii Latin characters
		password_parametrization{"myPassword124"},    // ascii numbers
		password_parametrization{"myPassword1!!¬&*"}, // ascii symbols
		password_parametrization{"这是我的密码"},           // non-ascii (Chinese) characters
		password_parametrization{"этомойпароль"},     // non-ascii (Russian) characters
	}

	for _, test_case := range parametrizations {
		// Monkeypatch terminal input with content converted to uint8 array
		newReadPassword := func(_ int) ([]uint8, error) {
			return test_case.as_uint8_array(), nil
		}
		monkey.Patch(terminal.ReadPassword, newReadPassword)
		defer monkey.Unpatch(terminal.ReadPassword) // clean-up

		master_password := getMasterPassword()
		if master_password != test_case.as_string {
			t.Errorf(
				"Unexpected master password. Expected: %s. Got: %s.",
				test_case.as_string,
				master_password,
			)
		}
	}
}

func TestGetDomain(t *testing.T) {
	parametrizations := []string{
		"google.com",                        // Domain only
		"http://www.google.com",             // Domain with scheme
		"http://www.mysite.com/path1/path2", // Domain with scheme and paths
	}

	for _, test_case := range parametrizations {
		// Monkeypatch stdin with content converted to uint8 string
		newReadString := func(_ *bufio.Reader, _ uint8) (string, error) {
			return test_case + "\n", nil
		}

		monkey.Patch((*bufio.Reader).ReadString, newReadString)
		defer monkey.Unpatch((*bufio.Reader).ReadString) // clean-up

		domain := getDomain()
		if domain != test_case {
			t.Errorf("Unexpected domain. Expected: %s. Got: %s.", test_case, domain)
		}
	}
}

type valid_input_parametrization struct {
	input    string
	is_valid bool
}

func TestValidateDomain(t *testing.T) {
	parametrizations := []valid_input_parametrization{
		valid_input_parametrization{"", false},
		valid_input_parametrization{"a", true},
		valid_input_parametrization{"2", true},
		valid_input_parametrization{"!", true},
	}

	for _, test_case := range parametrizations {
		if test_case.is_valid != validateDomain(test_case.input) {
			t.Errorf(
				"Unexpected result. Domain: %v. Expected validity: %v.",
				test_case.input,
				test_case.is_valid,
			)
		}
	}
}
