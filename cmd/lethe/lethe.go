package main

import (
	"fmt"
	"github.com/LachlanMarnham/Lethe/internal/cli"
)

func main() {
	fmt.Println("Hello, world!")
	var (
		string master_password
		string domain
	)
	master_password, domain = cli.GetSecrets()
}
