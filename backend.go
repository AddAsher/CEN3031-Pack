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
	"log"
	"regexp"
	"strings"
	// "bufio"
	// "os"
	// "unicode"
)

// to test use test@ufl.edu and test

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	ExpiresAt int64 `json:"exp"`
}

type Club struct {
	Description string `json:"description"`
	Leader      string `json:"leader"`
	Contact     string `json:"contact"`
	Hyperlink   string `json:"hyperlink"`
}

type NameAndPass struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Password   string     `json:"password"`
	LikedClubs [10]string `json:"likedClubs"`
}

var clubList = make(map[string]Club)

// var users = make(map[string]string)
var users = map[string]User{
	"Admin@ufl.edu": {"QWERTY1", [10]string{}},
}
var currUsername = "Admin@ufl.edu"

// clubList["A Reason to Give"] = "Our goal is to help serve the homeless population in Gainesville by providing weekly lunches!"
// clubList["Adventist Christian Fellowship"] = "Adventist Christian Fellowship is established for the purpose of representing the love of Jesus Christ. As a Christian Organization, we seek to spread the Advent Message and share the love of Jesus Christ to all those willing to receive it through open campus activities, Bible discussion, and other unique forms of fellowship. We show no preference or privilege on the basis of race, gender, sexual orientation or religious background."
// clubList["Advertising Society"] = "Ad Society is the advertising student organization at the University of Florida. Open to all majors, Ad Society holds general body meetings with guest speakers from the industry, takes trips to visit advertising agencies in various cities, and hosts workshops for students professional improvement."
// clubList["advnt"] = "advnt is a student-led creative program preparing the creatively inclined for careers in copywriting, illustration, graphic design, art direction, production, project management, and web design. We welcome those who see themselves as creatives and aspire to become capital ‘C’ creatives. All majors welcome."
// clubList["Alpha Omega - Jewish Dental Society"] = "Section 1 – Mission Statement To enhance the dental profession and lives of dental professionals worldwide by promoting and supporting ideals of global oral health, education, and the bonds of fraternity.	"
// clubList["Aquatic Animal Health Club"] = "The Aquatic Animal Health Club is a student-run organization dedicated to supplementing University of Florida’s College of Veterinary Medicine (UFCVM) curriculum by providing educational and networking opportunities in the field of aquatic animal medicine."
// clubList["The Society of PC Building"] = "The Society of PC Building's primary purpose is to provide students with an outlet to learn more about building PCs, regardless of experience, to foster knowledge and confidence in its members, and provide club members with hands-on experiences building PCs. The Society of PC Building will host a variety of events, meetings, and workshops for members to engage with each other and campus."
// clubList["3D Printing Club"] = "The 3D Printing Club is established for the purpose of educating UF students on the world of 3D printing and how 3D printing and related skills can be used within their education, professionally, and leisurely. Additionally, projects will be set in place to address issues seen within the University of Florida, the local Gainesville area, and nationally."
func main() {
	users["Beta"] = User{"Charlie", [10]string{"Charles"}}
	clubList["Adopted Student Organization"] = Club{"The purpose of the organization is to create a supportive community for all adoptees and those interested in learning more about adoption. We will host various events throughout the year as an opportunity to share our experiences and educate about adoption through meaningful discussions. We will also participate in related philanthropic efforts to support adoptees and children in foster care in the Gainesville community and around the world.", "Krista Marrocco", "krista.marrocco@ufl.edu", "https://orgs.studentinvolvement.ufl.edu/Organization/adopted-student-organization"}
	clubList["fooclub"] = Club{"foo", "Mr. Foo", "foofy@gmail.com", "foofy.com"}
	r := mux.NewRouter()
	r.HandleFunc("/registration", regHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/home/getClub", getClubs).Methods("GET")
	r.HandleFunc("/home/add", clubAdd).Methods("POST")
	r.HandleFunc("/home/like", likeClub).Methods("POST")
	r.HandleFunc("/home/getUsername", currentUsername).Methods("GET")

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

}

func (c *Claims) Valid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
}

func getClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubList)
}

func likeClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubList)
}

func currentUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currUsername)
}

func clubAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON data from the request body
	var clubData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Leader      string `json:"leader"`
		Contact     string `json:"contact"`
		Hyperlink     string `json:"hyperlink"`
	}
	err := json.NewDecoder(r.Body).Decode(&clubData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new Club struct with the decoded data
	newClub := Club{
		//	Name:		 clubData.Name,
		Description: clubData.Description,
		Leader:      clubData.Leader,
		Contact:     clubData.Contact,
		Hyperlink: clubData.Hyperlink,
	}

	// Add the new club to the clubList map
	//new club valid
	var valid string = newClubValid(clubData.Name, clubData.Description, clubData.Leader, clubData.Contact, clubData.Hyperlink) //HYPERLINK SHOULD BE ADDED HERE TOOO
	if valid != "Valid" {
		http.Error(w, valid, http.StatusBadRequest)
	} else {
		clubList[clubData.Name] = newClub
		fmt.Println("Club creation successful!")

	}

	// Send a response indicating success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Club added successfully"})
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
	} else if p == users[n].Password {
		return "Valid"
	} else {
		return "Incorrect password!"
	}
}

func newClubValid(n string, d string, l string, c string, h string) string { //im gonna add the hyperlink to this section once revamped
	if n == "" {
		return "Club name required!"
	} else if _, found := clubList[n]; found {
		return "This club already exists!"
	}

	if d == "" {
		return "Descripition is required!"
	}
	if l == "" {
		return "Club founder is required!"
	}
	if c == "" {
		return "Contact info is required!"
	} else if !strings.Contains(c, "@") {
		return "Contact information must be a valid email!"
	}
	if h == ""{
	return "Media link is required!"
	} else if !strings.Contains(h, "https://") {
		return "Please enter a valid link"
	}

	return "Valid"

}

func newUserValid(n string, p string) string {
	if n == "" {
		return "You cannot have a blank username!"
	} else if !strings.Contains(n, "@ufl.edu") {
		return "Your username must contain \"@ufl.edu\" at the end!"
	} else if _, found := users[n]; found {
		return "An account with this username already exists!"
	}

	if p == "" {
		return "You cannot have a blank password!"
	} else if len(p) < 6 {
		return "Your password must be at least 6 characters long!"
	} else if !regexp.MustCompile(`\d`).MatchString(p) {
		return "Your password must contain at least 1 digit!"
	} else {
		return "Valid"
	}
}

func regHandler(w http.ResponseWriter, r *http.Request) {
	//In here the name and password of the user are being put into
	//the password and liked Clubs field.
	fmt.Println("Received register request")
	// Parse request body
	decoder := json.NewDecoder(r.Body)
	var newUser NameAndPass
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var valid string = newUserValid(newUser.Username, newUser.Password)
	if valid != "Valid" {
		http.Error(w, valid, http.StatusBadRequest)
	} else {
		fmt.Println("Register Successful")
	}

	users[newUser.Username] = User{newUser.Password, [10]string{}}

	UsersJSON, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	w.Write(UsersJSON)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//In here the name and password of the user are being put into
	//the password and liked Clubs field.
	fmt.Println("Received login request")
	// Parse request body
	decoder := json.NewDecoder(r.Body)
	var possibleUser NameAndPass
	err := decoder.Decode(&possibleUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var valid string = userIsValid(possibleUser.Username, possibleUser.Password)
	if valid != "Valid" {
		http.Error(w, valid, http.StatusBadRequest)
	} else {
		fmt.Println("Login Successful")
	}

	currUsername = possibleUser.Username

	UsersJSON, _ := json.Marshal(possibleUser)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	w.Write(UsersJSON)
}
