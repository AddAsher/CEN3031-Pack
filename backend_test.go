package main
import(
	"testing"
) 

func TestUserIsValid(t *testing.T){
if userIsValid("johnsmith@aol.com", "QWERTY") == "Username must contain \"@ufl.edu\"!"{
	t.Log("@ufl.edu failure check success")
} else {
	t.Log("@ufl.edu success check failure")
}


if userIsValid("", "") == "Username is required!"{
	t.Log("No username input check success")
} else {
	t.Log("No username input check failure")
}

if userIsValid("johnsmith@ufl.edu", "") == "Account with this username not found!"{
	t.Log("Invalid username check success")
} else {
	t.Log("Invalid username check failure")
}


if userIsValid("Admin@ufl.edu", "") == "Password is required!"{
    t.Log("No password input check success")
} else {
    t.Log("No password input check failure ")
}


if userIsValid("Admin@ufl.edu", "QWERT") == "Incorrect password!"{
    t.Log("Incorrect password check success")
} else {
    t.Log("Incorrect password check failure")
}


if userIsValid("Admin@ufl.edu", "QWERTY") == "Valid"{
	t.Log("Valid username and password success")
} else {
	t.Error("Valid username and password failure")
}





	//t.Run("No username", func(t *testing.T)){

//	}
	
	
//	if loginHandler("test@ufl.edu") == "Username is required!"{
//		t.error("Username is required!")
//	}

}