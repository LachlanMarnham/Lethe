package cli

import (
	"testing"
	"bufio"
	"golang.org/x/crypto/ssh/terminal"
	"bou.ke/monkey"
)

func TestGetMasterPassword(t *testing.T) {
	var content string = "myPassWord31%$"

	// Monkeypatch terminal input with content converted to uint8 array 
	newReadPassword := func(_ int) ([]uint8, error) {
		return []uint8(content), nil
	}
	monkey.Patch(terminal.ReadPassword, newReadPassword)
	defer monkey.Unpatch(terminal.ReadPassword)  // clean-up

	master_password := getMasterPassword()
	if master_password != content {
		t.Errorf("Unexpected master password. Expected: %s. Got: %s.", content, master_password)
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
