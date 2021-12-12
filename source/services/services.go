package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/absoran/models"
)

//this function triggers when user send GET method to: //../api
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomePage")
	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to the API!")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

//this function triggers when user send GET method to: //../api/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Endpoint Hit: GetUsers")
		json.NewEncoder(w).Encode(models.Users)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

//this function triggers when user send POST method with body to: //../api/createnewuser
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateNewUser")
	if r.Method == "POST" { //request method validation. If method not POST than return BadRequest error
		reqBody, _ := ioutil.ReadAll(r.Body)
		var user models.User
		json.Unmarshal(reqBody, &user)
		for i := range models.Users { //checking if user already exist for avoid overwrite data
			if models.Users[i].ID == user.ID {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprint(w, "User with this id already exist")
				break
			}
		}
		models.Users = append(models.Users, user)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

//Error control function, if an error accured program will close
func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

//this function triggers when user send GET method with id to: //../api/getuserbyid/id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetUserById")
	if r.Method == "GET" {
		tempusr := models.User{ID: -1} // temporary user object with -1 id
		unparsedurl := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(unparsedurl[3])
		CheckError(err)
		for i := range models.Users {
			if models.Users[i].ID == id {
				tempusr = models.Users[i]
				break
			}
		}
		if tempusr.ID == -1 {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, "User does not exist") // if user with requested id does not exist return nocontent status code
		} else {
			json.NewEncoder(w).Encode(tempusr)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

//this function triggers when user send DELETE method with id to: //../api/deleteuserbyid/id
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteUserById")

	if r.Method == "DELETE" { //request method validation.
		index := -1
		unparsedurl := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(unparsedurl[3])
		CheckError(err)
		for i := range models.Users {
			if models.Users[i].ID == id {
				index = i
				break
			}
		}
		if index != -1 {
			remove(models.Users, index)
			fmt.Fprintf(w, "Deleted User with id: %d", id)
		} else {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, "User does not exist with id : %d", id)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

//this function triggers when user send PUT method with body and id to: //../api/updateuser/id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateUser")
	if r.Method == "PUT" { //request method validation.
		index := -1
		tempusr := models.User{Authority: -1}
		unparsedurl := strings.Split(r.URL.Path, "/")
		id, URLErr := strconv.Atoi(unparsedurl[3])
		CheckError(URLErr)
		reqBody, BodyErr := ioutil.ReadAll(r.Body)
		CheckError(BodyErr)
		json.Unmarshal(reqBody, &tempusr)

		for i := range models.Users {
			if models.Users[i].ID == id {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Fprintf(w, "User does not exist")
		} else {

			if tempusr.Age != 0 {
				models.Users[index].Age = tempusr.Age
			}
			if tempusr.FirstName != "" {
				models.Users[index].FirstName = tempusr.FirstName
			}
			if tempusr.LastName != "" {
				models.Users[index].LastName = tempusr.LastName
			}
			if tempusr.Authority != -1 {
				models.Users[index].Authority = tempusr.Authority
			}
			fmt.Fprint(w, "User updated", models.Users[index])
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}

// In this work we used arrays to store data. In array every object has index number that to indicate object's place.
// I used that index number to divide array to two different array and put together except index number
func remove(slice []models.User, s int) []models.User {
	return append(slice[:s], slice[s+1:]...)
}

func GetUserByAuthority(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetUserByAuthority")
	if r.Method == "GET" {
		var temp = []models.User{}
		unparsedurl := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(unparsedurl[3])
		CheckError(err)
		switch id {
		case 0:
			{
				for i := range models.Users {
					if models.Users[i].Authority == 0 {
						temp = append(temp, models.Users[i])
					}
				}
				json.NewEncoder(w).Encode(temp)
			}
		case 1:
			{
				for i := range models.Users {
					if models.Users[i].Authority == 1 {
						temp = append(temp, models.Users[i])
					}
				}
				json.NewEncoder(w).Encode(temp)
			}
		case 2:
			{
				for i := range models.Users {
					if models.Users[i].Authority == 2 {
						temp = append(temp, models.Users[i])
					}
				}
				json.NewEncoder(w).Encode(temp)
			}
		case 3:
			{
				for i := range models.Users {
					if models.Users[i].Authority == 3 {
						temp = append(temp, models.Users[i])
					}
				}
				json.NewEncoder(w).Encode(temp)
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Used wrong HTTP method !")
	}
}
