package cli

import (
	"testing"
	"golang.org/x/crypto/ssh/terminal"
	"bou.ke/monkey"
)

func TestMasterPassword(t *testing.T) {
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
