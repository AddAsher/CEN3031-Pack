
## What issues your team planned to address
* We planned to add in user account registration and adding any created user into a map structure
* We planned on modifying our login screen to accomodate for the registration change
* We planned to add in account requirements such as requiring symbols and certain lengths for passwords
* We planned to implement a basic club listing page using various stored values in our club map structure
* We planned to add in various backend testing for our verification


## Which ones were successfully completed
* Successfully added an account registration into our backend through the usage of a map
* Successfully modified our login screen with the new registration option, as well as some other visual changes
* Successfully implemeneted password length/digit requirements in signups
* Backend testing implemeneted for validating a user on login as well as validating a user registering an account

## Backend Testing Details
* Implemented in backend_test.go with testing of the functions newUserValid and userIsValid
* newUserValid checks if email contains @ufl.edu, as well as if the password is at least 6 characters and has 1 digit.
* userIsValid checks if email contains @ufl.edu and that the email matches the password associated with it in the map.
* Created func TestUserIsValid and func TestNewUserValid
* TestNewUserValid tests registering accounts of different email types, empty strings, and passwords not fulfilling requirements.
* TestUserIsValid tests accounts of different email types, empty strings, wrong passwords, and not having an email registered.
## Which ones didn't and why?
* We werent able to implement a basic club listing page due to last minute implementation issues on the front end