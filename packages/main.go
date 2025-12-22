package main

import (
	"fmt"

	"github.com/SameerAkita/go-practice/auth"
	"github.com/SameerAkita/go-practice/user"
)

func main() {
	auth.LoginWithCredentials("sameer", "password")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email, user.Name)
	// color.Red(user.Email)
}
