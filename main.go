package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/DusanKasan/parsemail"
)

func main() {
	pwd, _ := os.Getwd()
	if file, err := os.Open(pwd + "/email_files/original_msg.txt"); err != nil {
		panic(err)
	} else {
		reader := bufio.NewReader(file)
		email, err := parsemail.Parse(reader) // returns Email struct and error
		if err != nil {
			panic(err)
		}

		fmt.Println(email.Subject)
		fmt.Println(email.From)
		fmt.Println(email.To)
		fmt.Println(email.HTMLBody)
	}
}
