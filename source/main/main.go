package main

import (
	"log"
	"net/http"

	//"github.com/absoran/dbconnection"

	"github.com/absoran/services"
)

func main() {
	handleRequests()
}

func handleRequests() {

	apiRoot := "/api"
	// .../api
	http.HandleFunc(apiRoot, services.HomePage)
	// .../api/users
	http.HandleFunc(apiRoot+"/users", services.GetUsers)
	// .../api/getuserbyid/id
	http.HandleFunc(apiRoot+"/getuserbyid/", services.GetUserById)
	// .../api/deleteuserbyid/id
	http.HandleFunc(apiRoot+"/deleteuserbyid/", services.DeleteUserById)
	// .../api/createnewuser
	http.HandleFunc(apiRoot+"/createnewuser", services.CreateNewUser)
	// .../api/updateuser/id
	http.HandleFunc(apiRoot+"/updateuser/", services.UpdateUser)
	// .../api/GetUserByAuthority/id
	http.HandleFunc(apiRoot+"/getuserbyauthority/", services.GetUserByAuthority)
	//connection open
	log.Fatal(http.ListenAndServe(":8080", nil))

	//another way to open connection on port 8080
	//defer http.ListenAndServe(":8080", nil)
}
