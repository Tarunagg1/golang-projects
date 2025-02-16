package main

import (
	"github.com/fatih/color"
	"github.com/tarunagg/package/auth"
	"github.com/tarunagg/package/user"
)

func main() {
	auth.LoginithCredientials("tarungmialcom", "fojo")
	auth.GetSession()

	user := user.User{Email: "email", Name: "Name"}

	// fmt.Println(user)

	color.Green(user.Email)
}
