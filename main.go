package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	if password, err := terminal.ReadPassword(int(os.Stdin.Fd())); nil != err {
		panic(err)
	} else {
		fmt.Printf("%s", password)
	}
}
