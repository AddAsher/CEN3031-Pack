# Changes Made
Added user account registration functionality to the front end by connecting it to the backend's map structure made last sprint.
Changed our map to store a club struct that contained a club's description, leader, contact information, and hyperlink to any social media, with the key still being the club name.
Updated user struct to contain an
Added functions that allow users to register their club's on the website (clubAdd) 
Added the account registration functionality to our front end which writes to the map for a new user.
Successfully implemented a frontend page that displays club data by pulling from the club map.
Successfully implemeneted picture showcasing on the frontend through pictures specified for specific clubs.
Successfully implemented a logout function in the frontend
Added Backend testing for user's adding clubs






# API Documentation
## New User Registration Endpoint
### Endpoint URL
http://localhost:4200/registration
### Method
POST
### Request/Response Parameters
JSON object with a username and password as a string as the request, and returns the map with the new user inside.
###Successful/Error responses
"200 OK" on success (meets new user requirements)
"400 Bad Request" on failure (does not meet new user requirements)
### Functions:
regHandler(w http.Responsewriter, r \*http.Request)- Checks if the provided username/password meet requirements (i.e: ufl email, length, no duplicate account), and if it does, places it inside our map and calls write on the response writer.

## Existing User Login Endpoint
### Endpoint URL
http://localhost:4200/login
### Method
POST
### Request/Response Parameters
JSON object with a username and password as a string as both the return request and response. 
### Successful/Error responses
"200 OK" on success (Existing user with this username/password found)
"400 Bad Request" on failure (No user with this username/password found)
### Functions:
loginHandler(w http.ResponseWriter, r \*http.Request) - Checks if provided username/password exist in the map, if not, returns 400 Bad Request, else returns 200 OK and allows login.

## Get Existing Club list
### Endpoint URL
http://localhost:4200/home
### Method
GET
### Request/Response Parameters
Returns the clubList map from the back end.
### Functions:
getClubs(w http.ResponseWriter, r \*http.Request) - Returns clubList map to the frontend. 

## Adding Club To Club List
### Endpoint URL
http://localhost:4200/home
### Method
POST
### Request/Response Parameters
Receives a club name, descripition, leader, and contact information, and returns the clubList map.
### Functions:
clubAdd(w http.ResponseWriter, r \*http.request) - Ensures that all passed in information is provided and checks for duplicates in the map. If these conditions are met, it adds the club and returns the map.
