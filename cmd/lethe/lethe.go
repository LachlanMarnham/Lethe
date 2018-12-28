package main

import (
	"fmt"
	"github.com/LachlanMarnham/Lethe/internal/cli"
	"github.com/LachlanMarnham/Lethe/internal/password_factory"
)

func main() {
	var (
		master_password string
		domain string
		password string
	)
	master_password, domain = cli.GetSecrets()
	password = password_factory.Make(master_password, domain)
	// if flag not set:
	// 		cli.SendToStdout(password)
	// else:
	// 		cli.SendToClipboard(password)
	fmt.Println("Password:", password)
	fmt.Println("Domain:", domain)
	fmt.Println("Master password:", master_password)

}
