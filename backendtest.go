package main

import "fmt"
//to test use test@ufl.edu and test
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
			fmt.Println("Email:")
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
			fmt.Println("Password:")
		}

	}
	fmt.Println("Please choice an option.")
	fmt.Println("1. View Club Hostings")
	fmt.Println("2. Host a club")
	fmt.Println("3. Quit")
	fmt.Print("Enter your choice: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case "1\n":
		fmt.Println("You selected option 1") //I think we just list a bunch of clubs here and then u can click for further info?
	case "2\n":
		fmt.Println("You selected option 2") //adds to the database of clubs prob just a vector
	case "3\n":
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Println("Invalid choice")


}
