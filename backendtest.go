package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"fmt"
	// "bufio"
	// "os"
	// "unicode"
	// "strings"
	"log"
)

// to test use test@ufl.edu and test
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	ExpiresAt int64 `json:"exp"`
}

// type Club struct { //We'll use this when we have a CSV, for now we can just use println to show how it would function.
// 	name        string
// 	description string
// }

// Club clubs := []Club{}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string)

func main() {
	users["Admin"] = "QWERTY"
	users["Beta"] = "Charlie"
	users["Delta"] = "Echo"
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
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
	// clubList := make(map[int]Club)
	// newClub := Club{name: "test", description: "this is a club"}
	// clubList[1] = newClub
	// clubList[2] = newClub
	// fmt.Println("Welcome to Pack!")
	// fmt.Println("Please Enter your email and password") //just make sure email has ufl, later tho
	// fmt.Print("Email:")
	// var email string
	// for {
	// 	fmt.Scanln(&email)
	// 	if !unicode.IsLetter(rune(email[0])) {
	// 		fmt.Println("Invalid Email!")
	// 		fmt.Print("Email:")
	// 	} else if strings.Contains(email, "@ufl.edu") {
	// 		break
	// 	} else {
	// 		fmt.Println("Email not found!")
	// 		fmt.Print("Email:")
	// 	}
	// }
	// fmt.Print("Password:")
	// var pass string
	// for {
	// 	fmt.Scanln(&pass)
	// 	if pass == email[1:4] {
	// 		break

	// 	} else {
	// 		fmt.Println("Password not valid")
	// 		fmt.Print("Password:")
	// 	}

	// }
	// var menuDone bool = false
	// var subMenu bool = false

	// for !menuDone {
	// 	fmt.Println("Please choose an option.")
	// 	fmt.Println("1. View Club Hostings")
	// 	fmt.Println("2. Request New Listing for Club")
	// 	fmt.Println("3. Quit")
	// 	fmt.Print("Enter your choice: ")
	// 	var input int
	// 	fmt.Scanln(&input)

	// 	switch input {

	// 	case 1:
	// 		subMenu = false

	// 		for !subMenu {
	// 			if len(clubList) == 0 {
	// 				fmt.Println("No clubs avaliable")
	// 			} else {
	// 				for key, value := range clubList {

	// 					fmt.Print(key)
	// 					fmt.Print(". ")
	// 					fmt.Println(value.name)

	// 				}
	// 			}
	// 			fmt.Println("Which club would you like to know more about?")
	// 			fmt.Println("-1 to return to menu")
	// 			var menuChoice int
	// 			fmt.Scanln(&menuChoice)
	// 			if menuChoice == -1 {
	// 				subMenu = true
	// 				break
	// 			}
	// 			//switch menuChoice {
	// 			fmt.Println(clubList[menuChoice].description)
	//case 1:
	//fmt.Println("The purpose of the 360BHM is to this series seeks to educate, enhance, and entertain the university and Gainesville communities by reflecting on the contributions that Black Americans have made. It is our goal to advocate exposure of refined enrichment inspired by the heritage and legacy of universal Black culture.")
	//break
	//case 2:
	//	fmt.Println("The 3D Printing Club is established for the purpose of educating UF students on the world of 3D printing and how 3D printing and related skills can be used within their education, professionally, and leisurely. Additionally, projects will be set in place to address issues seen within the University of Florida, the local Gainesville area, and nationally.")
	//	break
	//case 3:
	//	fmt.Println("Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!")
	//meeting hours here but  we gotta get more info from clubs for that
	//meeting location here also
	//case 9:

	// 		}

	// 	case 2:
	// 		scanner := bufio.NewScanner(os.Stdin)
	// 		fmt.Println("Input your club's name") //adds to the database of clubs prob just a vector, need some sort of confirmation here
	// 		var clubName string
	// 		scanner.Scan()
	// 		clubName = scanner.Text()
	// 		fmt.Scanln(&clubName)

	// 		fmt.Println("Tell us a little bit about your club") //has to run through our approval first so people cant just put random stuff
	// 		var clubDesc string
	// 		fmt.Scanln(&clubDesc)
	// 		scanner.Scan()
	// 		clubDesc = scanner.Text()
	// 		fmt.Println("Thank you! We will consider adding", clubName) //for the sake of this we're just adding it
	// 		newClub := Club{name: clubName, description: clubDesc}
	// 		clubList[len(clubList)] = newClub
	// 		break

	// 	case 3:
	// 		fmt.Println("Goodbye!")
	// 		menuDone = true
	// 		subMenu = true
	// 		return
	// 	default:
	// 		fmt.Println("Invalid choice")
	// 		break
	// 	}

	// 	//I think we just list a bunch of clubs here and then u can click for further info?

	// }

}

func (c *Claims) Valid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
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

	var validUser bool = false
	for !validUser {
		if newUser.Username == "" {
			http.Error(w, "username is required", http.StatusBadRequest)
			return
		} else if _, found := users[newUser.Username]; found {
			fmt.Println(newUser.Username)
		} else {
			http.Error(w, "invalid username", http.StatusBadRequest)
			return
		}

		if newUser.Password == "" {
			http.Error(w, "password is required", http.StatusBadRequest)
			return
		} else if newUser.Password == users[newUser.Username] {
			fmt.Println(newUser.Password)
			validUser = true
		} else {
			http.Error(w, "invalid password", http.StatusBadRequest)
			return
		}
	}

	users[newUser.Username] = newUser.Password

	UsersJSON, _ := json.Marshal(newUser)
	//response, err := json.Marshal(newUser)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(UsersJSON)
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	var user User

// 	// Parse the JSON request body
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Do some logic to verify the user's credentials
// 	// ...

// 	// Send back a response
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	var data = struct {
// 		Title string `json:"title"`
// 	}{
// 		Title: "Golang + Angular Starter Kit",
// 	}

// 	jsonBytes, err := utils.StructToJSON(data)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonBytes)
// 	return
// }

//r.HandleFunc("/api/login", handleLogin).Methods("POST")
//http.ListenAndServe(":8080", r)
//API Endpoints
// r.HandleFunc("/api/users", getUsers).Methods("GET")
// r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
// r.HandleFunc("/api/users", createUser).Methods("POST")
// r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
// r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

//club1 := Club{
//	name:  "360HM", description: "The purpose of the 360BHM is to this series seeks to educate, enhance, and entertain the university and Gainesville communities by reflecting on the contributions that Black Americans have made. It is our goal to advocate exposure of refined enrichment inspired by the heritage and legacy of universal Black culture."}

//clubVector = append(clubs, club1)

//for i := 0; i < len(clubVector{}); i++ {
//	fmt.Println(clubVector[i].name);
//}

//we put the csv stuff here but we need that data to import
// 	clubList := make(map[string]string)
// 	fmt.Println("Welcome to Pack!")
// 	fmt.Println("Please Enter your email and password") //just make sure email has ufl, later tho
// 	fmt.Print("Email:")
// 	var email string
// 	for {
// 		fmt.Scanln(&email)
// 		if !unicode.IsLetter(rune(email[0])) {
// 			fmt.Println("Invalid Email!")
// 			fmt.Print("Email:")
// 		} else if strings.Contains(email, "@ufl.edu") {
// 			break
// 		} else {
// 			fmt.Println("Email not found!")
// 			fmt.Print("Email:")
// 		}
// 	}
// 	fmt.Print("Password:")
// 	var pass string
// 	for {
// 		fmt.Scanln(&pass)
// 		if pass == email[1:4] {
// 			break

// 		} else {
// 			fmt.Println("Password not valid")
// 			fmt.Print("Password:")
// 		}

// 	}
// 	var menuDone bool = false
// 	var subMenu bool = false

// 	for !menuDone {
// 		fmt.Println("Please choose an option.")
// 		fmt.Println("1. View Club Hostings")
// 		fmt.Println("2. Request New Listing for Club")
// 		fmt.Println("3. Quit")
// 		fmt.Print("Enter your choice: ")
// 		var input int
// 		fmt.Scanln(&input)

// 		switch input {

// 		case 1:
// 			subMenu = false
// 			for !subMenu {
// 				fmt.Println("1. 360BHM") //this is just the mockup i'd rather get these all on a CSV somehow and then this section is just lego lab from prog 2
// 				fmt.Println("2. 3D Printing club")
// 				fmt.Println("3. A Reason to Give")
// 				fmt.Println("4. Adopted Student Organization")
// 				fmt.Println("5. Advertising Society")
// 				fmt.Println("6. Adventist Christian Fellowship")
// 				fmt.Println("7. 3D Printing club")
// 				fmt.Println("8. Next Page")
// 				fmt.Println("9. Return")
// 				fmt.Println("Choose a Club you'd like to know more about!")

// 				var menuChoice int
// 				fmt.Scanln(&menuChoice)
// 				switch menuChoice {
// 				case 1:
// 					fmt.Println("The purpose of the 360BHM is to this series seeks to educate, enhance, and entertain the university and Gainesville communities by reflecting on the contributions that Black Americans have made. It is our goal to advocate exposure of refined enrichment inspired by the heritage and legacy of universal Black culture.")
// 					break
// 				case 2:
// 					fmt.Println("The 3D Printing Club is established for the purpose of educating UF students on the world of 3D printing and how 3D printing and related skills can be used within their education, professionally, and leisurely. Additionally, projects will be set in place to address issues seen within the University of Florida, the local Gainesville area, and nationally.")
// 					break
// 				case 3:
// 					fmt.Println("Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!")
// 					//meeting hours here but  we gotta get more info from clubs for that
// 					//meeting location here also
// 				case 9:
// 					subMenu = true
// 					break
// 				}
// 			}

// 			//I think we just list a bunch of clubs here and then u can click for further info?
// 		case 2:
// 			scanner := bufio.NewScanner(os.Stdin)
// 			fmt.Println("Input your club's name") //adds to the database of clubs prob just a vector, need some sort of confirmation here
// 			var clubName string
// 			scanner.Scan()
// 			clubName = scanner.Text()
// 			//fmt.Scanln(&clubName)

// 			fmt.Println("Tell us a little bit about your club") //has to run through our approval first so people cant just put random stuff
// 			var clubDesc string
// 			//fmt.Scanln(&clubDesc)
// 			scanner.Scan()
// 			clubDesc = scanner.Text()
// 			fmt.Println("Thank you! We will consider adding", clubName)
// 			clubList[clubName] = clubDesc
// 			break

// 		case 3:
// 			fmt.Println("Goodbye!")
// 			menuDone = true
// 			subMenu = true
// 			return
// 		default:
// 			fmt.Println("Invalid choice")
// 			break
// 		}

// 	}

// }

// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	var req LoginRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// TODO: Validate user's credentials

// 	if true {
// 		w.WriteHeader(http.StatusOK)
// 	} else {
// 		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
// 	}
// }
