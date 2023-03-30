package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// to test use test@ufl.edu and test

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	ExpiresAt int64 `json:"exp"`
}

type Club struct { //We'll use this when we have a CSV, for now we can just use println to show how it would function.
	name        string
	description string
}

// Club clubs := []Club{}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var clubList = make(map[string]string)
var users = make(map[string]string)

func main() {
	users["Admin@ufl.edu"] = "QWERTY"
	users["Beta"] = "Charlie"
	users["Delta"] = "Echo"
	clubList["fooclub"] = "foo"
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/login", getClubs).Methods("GET")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	handler := c.Handler(r)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}

	//clubList := make(map[string]string)
	clubList := make(map[string]string)
	loginList := make(map[string]string)
	loginList["test"] = "test"
	clubList["Adopted Student Organization"] = "The purpose of the organization is to create a supportive community for all adoptees and those interested in learning more about adoption. We will host various events throughout the year as an opportunity to share our experiences and educate about adoption through meaningful discussions. We will also participate in related philanthropic efforts to support adoptees and children in foster care in the Gainesville community and around the world."
	clubList["A Reason to Give"] = "Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!"
	clubList["Adventist Christian Fellowship"] = "Adventist Christian Fellowship is established for the purpose of representing the love of Jesus Christ. As a Christian Organization, we seek to spread the Advent Message and share the love of Jesus Christ to all those willing to receive it through open campus activities, Bible discussion, and other unique forms of fellowship. We show no preference or privilege on the basis of race, gender, sexual orientation or religious background."
	clubList["Advertising Society"] = "Ad Society is the advertising student organization at the University of Florida. Open to all majors, Ad Society holds general body meetings with guest speakers from the industry, takes trips to visit advertising agencies in various cities, and hosts workshops for students professional improvement."
	clubList["advnt"] = "advnt is a student-led creative program preparing the creatively inclined for careers in copywriting, illustration, graphic design, art direction, production, project management, and web design. We welcome those who see themselves as creatives and aspire to become capital ‘C’ creatives. All majors welcome."
	clubList["A Reason to Give"] = "Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!"
	clubList["A Reason to Give"] = "Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!"
	clubList["A Reason to Give"] = "Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!"
	clubList["3D Printing Club"] = "The 3D Printing Club is established for the purpose of educating UF students on the world of 3D printing and how 3D printing and related skills can be used within their education, professionally, and leisurely. Additionally, projects will be set in place to address issues seen within the University of Florida, the local Gainesville area, and nationally."
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
				var i int = 1
				if len(clubList) == 0 {
					fmt.Println("No clubs avaliable")
				} else {
					for key := range clubList {
						fmt.Print(i)
						fmt.Print(". ")
						fmt.Println(key)
						i++

					}
				}
				fmt.Println("Please enter the name of the club you'd like to know more about.")
				fmt.Println("-1 to return to menu")
				var menuChoice string
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				menuChoice = scanner.Text()
				if menuChoice == "-1" {
					subMenu = true
					break
				}
				if value, found := clubList[menuChoice]; found {
					fmt.Println(value)
				} else {
					fmt.Println("invalid club!")
				}
				fmt.Println("Would you like to view more clubs? Input Yes or No")
				var inputCheck bool = false
				for !inputCheck {
					fmt.Scanln(&menuChoice)

					if menuChoice == "Yes" {
						inputCheck = true
						continue

					}
					if menuChoice == "No" {
						subMenu = true
						inputCheck = true
					} else {
						fmt.Println("Invalid choice, input Yes or No")
					}
				}

			}

		case 2:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Input your club's name") //adds to the database of clubs prob just a vector, need some sort of confirmation here
			var clubName string
			scanner.Scan()
			clubName = scanner.Text()
			fmt.Println("Tell us a little bit about your club") //has to run through our approval first so people cant just put random stuff
			var clubDesc string
			//fmt.Scanln(&clubDesc)
			scanner.Scan()
			clubDesc = scanner.Text()

			fmt.Println("Thank you! We will consider adding", clubName) //for the sake of this we're just adding it
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

		//I think we just list a bunch of clubs here and then u can click for further info?

	}

}

func (c *Claims) Valid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
}

func getClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	ClubsJSON, err := json.Marshal(clubList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(ClubsJSON)
}

func userIsValid(n string, p string) string {
	if n == "" {
		return "Username is required!"
	} else if !strings.Contains(n, "@ufl.edu") {
		return "Username must contain \"@ufl.edu\"!"
	} else if _, found := users[n]; !found {
		return "Account with this username not found!"
	}

	if p == "" {
		return "Password is required!"
	} else if p == users[n] {
		return "Valid"
	} else {
		return "Incorrect password!"
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received login request")
	// Parse request body
	decoder := json.NewDecoder(r.Body)
	var newUser User
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var valid string = userIsValid(newUser.Username, newUser.Password)
	if valid != "Valid" {
		http.Error(w, valid, http.StatusBadRequest)
	} else {
		fmt.Println("Login Successful")
	}

	users[newUser.Username] = newUser.Password

	UsersJSON, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	w.Write(UsersJSON)
}
