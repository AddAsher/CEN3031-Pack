package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pack!")
	fmt.Println("Please Enter your email and password") //just make sure email has ufl, later tho
	fmt.Println("Email:")
	var email string
	for {
		fmt.Scanln(&email)
		if email == "test@ufl.edu" {
			break

		} else {
			fmt.Println("Email not found!")
		}
	}
	fmt.Println("Password:")
	var pass string
	fmt.Scanln(&pass)
	for {
		if pass == "test" {
			break

		} else {
			fmt.Println("Password not valid")
		}

	}
}
