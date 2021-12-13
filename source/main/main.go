package main

import (
	"net/http"

	//"github.com/absoran/dbconnection"

	"github.com/absoran/services"
	"github.com/gorilla/mux"
)

var ApiRoot = "/api"

func main() {
	//handleRequests()
	handleRequestWithMux()
}
func handleRequestWithMux() {
	router := mux.NewRouter()
	defer http.ListenAndServe(":8080", router)
	//mux handlers
	// .../api/getuserswmux
	router.HandleFunc(ApiRoot+"/userswmux", services.GetUsersWithMux).Methods("GET")
	// .../api/getuserbyidwithmux
	router.HandleFunc(ApiRoot+"/getuserbyidwmux/{id}", services.GetUserByIdWithMux).Methods("GET")
	// .../api/getusersbyautoritywithmux
	router.HandleFunc(ApiRoot+"/getusersbyautoritywmux/{id}", services.GetUsersByAuthorityWithMux).Methods("GET")
	// .../api/deleteuserbyid/id
	router.HandleFunc(ApiRoot+"/deleteuserbyidwmux/{id}", services.DeleteUserByIdWithMux).Methods("DELETE")
	// .../api/createnewuser
	router.HandleFunc(ApiRoot+"/createnewuserwmux", services.CreateNewUserWithMux).Methods("POST")
	// .../api/updateuser/id
	router.HandleFunc(ApiRoot+"/updateuserwmux/{id}", services.UpdateUserWithMux).Methods("PUT")
}
func handleRequests() {

	defer http.ListenAndServe(":8080", nil)

	// .../api
	http.HandleFunc(ApiRoot, services.HomePage)
	// .../api/users
	http.HandleFunc(ApiRoot+"/users", services.GetUsers)
	// .../api/getuserbyid/id
	http.HandleFunc(ApiRoot+"/getuserbyid/", services.GetUserById)
	// .../api/deleteuserbyid/id
	http.HandleFunc(ApiRoot+"/deleteuserbyid/", services.DeleteUserById)
	// .../api/createnewuser
	http.HandleFunc(ApiRoot+"/createnewuser", services.CreateNewUser)
	// .../api/updateuser/id
	http.HandleFunc(ApiRoot+"/updateuser/", services.UpdateUser)
	// .../api/GetUserByAuthority/id
	http.HandleFunc(ApiRoot+"/getuserbyauthority/", services.GetUserByAuthority)
	//another way to open connection on port 8080
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
