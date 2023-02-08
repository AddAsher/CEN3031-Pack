package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

//to test use test@ufl.edu and test

type Club struct { //We'll use this when we have a CSV, for now we can just use println to show how it would function.
	name        string
	description string
}

//Club clubs := []Club{}

func main() {
	//club1 := Club{
	//	name:  "360HM", description: "The purpose of the 360BHM is to this series seeks to educate, enhance, and entertain the university and Gainesville communities by reflecting on the contributions that Black Americans have made. It is our goal to advocate exposure of refined enrichment inspired by the heritage and legacy of universal Black culture."}

	//clubVector = append(clubs, club1)

	//for i := 0; i < len(clubVector{}); i++ {
	//	fmt.Println(clubVector[i].name);
	//}

	//we put the csv stuff here but we need that data to import
	clubList := make(map[string]string)
	fmt.Println("Welcome to Pack!")
	fmt.Println("Please Enter your email and password") //just make sure email has ufl, later tho
	fmt.Print("Email:")
	var email string
	for {
		fmt.Scanln(&email)
		if !unicode.IsLetter(rune(email[0])) {
			fmt.Println("Invalid Email!")
			fmt.Print("Email:")
		} else if strings.Contains(email, "@ufl.edu") {
			break
		} else {
			fmt.Println("Email not found!")
			fmt.Print("Email:")
		}
	}
	fmt.Print("Password:")
	var pass string
	for {
		fmt.Scanln(&pass)
		if pass == email[1:4] {
			break

		} else {
			fmt.Println("Password not valid")
			fmt.Print("Password:")
		}

	}
	var menuDone bool = false
	var subMenu bool = false

	for !menuDone {
		fmt.Println("Please choose an option.")
		fmt.Println("1. View Club Hostings")
		fmt.Println("2. Request New Listing for Club")
		fmt.Println("3. Quit")
		fmt.Print("Enter your choice: ")
		var input int
		fmt.Scanln(&input)

		switch input {

		case 1:
			subMenu = false
			for !subMenu {
				fmt.Println("1. 360BHM") //this is just the mockup i'd rather get these all on a CSV somehow and then this section is just lego lab from prog 2
				fmt.Println("2. 3D Printing club")
				fmt.Println("3. A Reason to Give")
				fmt.Println("4. Adopted Student Organization")
				fmt.Println("5. Advertising Society")
				fmt.Println("6. Adventist Christian Fellowship")
				fmt.Println("7. 3D Printing club")
				fmt.Println("8. Next Page")
				fmt.Println("9. Return")
				fmt.Println("Choose a Club you'd like to know more about!")

				var menuChoice int
				fmt.Scanln(&menuChoice)
				switch menuChoice {
				case 1:
					fmt.Println("The purpose of the 360BHM is to this series seeks to educate, enhance, and entertain the university and Gainesville communities by reflecting on the contributions that Black Americans have made. It is our goal to advocate exposure of refined enrichment inspired by the heritage and legacy of universal Black culture.")
					break
				case 2:
					fmt.Println("The 3D Printing Club is established for the purpose of educating UF students on the world of 3D printing and how 3D printing and related skills can be used within their education, professionally, and leisurely. Additionally, projects will be set in place to address issues seen within the University of Florida, the local Gainesville area, and nationally.")
					break
				case 3:
					fmt.Println("Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!")
					//meeting hours here but  we gotta get more info from clubs for that
					//meeting location here too probably
				case 9:
					subMenu = true
					break
				}
			}

			//I think we just list a bunch of clubs here and then u can click for further info?
		case 2:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Input your club's name") //adds to the database of clubs prob just a vector, need some sort of confirmation here
			var clubName string
			scanner.Scan()
			clubName = scanner.Text()
			//fmt.Scanln(&clubName)

			fmt.Println("Tell us a little bit about your club") //has to run through our approval first so people cant just put random stuff
			var clubDesc string
			//fmt.Scanln(&clubDesc)
			scanner.Scan()
			clubDesc = scanner.Text()
			fmt.Println("Thank you! We will consider adding", clubName)
			clubList[clubName] = clubDesc
			break

		case 3:
			fmt.Println("Goodbye!")
			menuDone = true
			subMenu = true
			return
		default:
			fmt.Println("Invalid choice")
			break
		}

	}
}
