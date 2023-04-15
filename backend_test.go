package main

import (
	"testing"
)

func TestUserIsValid(t *testing.T) {
	if userIsValid("johnsmith@aol.com", "QWERTY") == "Username must contain \"@ufl.edu\"!" {
		t.Log("@ufl.edu failure check success")
	} else {
		t.Error("@ufl.edu success check failure")
	}

	if userIsValid("", "") == "Username is required!" {
		t.Log("No username input check success")
	} else {
		t.Error("No username input check failure")
	}

	if userIsValid("johnsmith@ufl.edu", "") == "Account with this username not found!" {
		t.Log("Invalid username check success")
	} else {
		t.Error("Invalid username check failure")
	}

	if userIsValid("Admin@ufl.edu", "") == "Password is required!" {
		t.Log("No password input check success")
	} else {
		t.Error("No password input check failure ")
	}

	if userIsValid("Admin@ufl.edu", "QWERT") == "Incorrect password!" {
		t.Log("Incorrect password check success")
	} else {
		t.Error("Incorrect password check failure")
	}

	if userIsValid("Admin@ufl.edu", "QWERTY1") == "Valid" {
		t.Log("Valid username and password success")
	} else {
		t.Error("Valid username and password failure")
	}
}

func TestNewUserValid(t *testing.T) {
	if newUserValid("", "") == "You cannot have a blank username!" {
		t.Log("Blank username check success")
	} else {
		t.Error("Blank username check failure")
	}

	if newUserValid("Admin@gmail.com", "h") == "Your username must contain \"@ufl.edu\" at the end!" {
		t.Log("@UFL.EDU check success")
	} else {
		t.Error("@UFL.EDU check failure")
	}

	if newUserValid("Admin@ufl.edu", "") == "You cannot have a blank password!" {
		t.Log("Blank password check success")
	} else {
		t.Error("Blank password check failure")
	}

	if newUserValid("Admin@ufl.edu", "") == "You cannot have a blank password!" {
		t.Log("Blank password check success")
	} else {
		t.Error("Blank password check failure")
	}

	if newUserValid("Admin@ufl.edu", "asd") == "Your password must be at least 6 characters long!" {
		t.Log("Password length check success")
	} else {
		t.Error("Password length check failure")
	}

	if newUserValid("Admin@ufl.edu", "QWERTY") == "Your password must contain at least 1 digit!" {
		t.Log("Password digit check success")
	} else {
		t.Error("Password digit check failure")
	}

	if newUserValid("Admin@ufl.edu", "QWERTY1") == "Valid" {
		t.Log("Account creation successful")
	} else {
		t.Error("Account creation failure")
	}

}

//t.Run("No username", func(t *testing.T)){

//	}

//	if loginHandler("test@ufl.edu") == "Username is required!"{
//		t.error("Username is required!")
//	}
