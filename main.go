package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// foo data structure, 2 feilds: name and id
type Foo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// store the data name and id in this array
var holdFoo []Foo

// root URL http://localhost:8080
func root(w http.ResponseWriter, r *http.Request) {
	// define the content type
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Rakuten Webservice HW Assignment")
}

// generate a random UUID for the id key
func generateUUID() string {
	id := uuid.New().String()
	return id
}

// GET endpoint (‘foo/{id}’)
// use id, search through holdFoo array
func get(w http.ResponseWriter, r *http.Request) {
	// define the content type
	w.Header().Set("Content-Type", "application/json")
	// gorilla mux router
	vars := mux.Vars(r)
	// get the key value inputted id from input
	key := vars["id"]

	// Loop over holdFoo
	for _, foo := range holdFoo {
		// if the foo.Id equals the key:id we pass in
		if foo.Id == key {
			// 200 response code: Success
			w.WriteHeader(http.StatusOK)
			// javascript obj notation encode and decode JSON
			// return the data encoded as JSON
			json.NewEncoder(w).Encode(foo)
		}
	}
	// 404 response Not Found when you get an id that does not exist
	if vars != nil {
		w.WriteHeader(404)
	}
}

// POST endpoint ('/foo')
// create new name and id
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the body of our POST request, byte
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Println("@@@@@@@@@@this is in reqBody", reqBody)
	// variable foo = {}
	var foo Foo
	// set random UUID to id
	foo.Id = generateUUID()
	// unmarshal this into a new Foo struct
	// decode the data from JSON to a struct
	// id value is sent to the body of the request
	// &foo getting the memory address:
	json.Unmarshal(reqBody, &foo)
	// append this to our holdFoos array.
	holdFoo = append(holdFoo, foo)
	// convert to json and write it
	json.NewEncoder(w).Encode(foo)
}

func delete(w http.ResponseWriter, r *http.Request) {
	// parse through param
	vars := mux.Vars(r)
	// get the id we want to delete
	id := vars["id"]

	// loop through all the data
	for index, foo := range holdFoo {
		// check if it matches param
		// id holds inputted id
		if foo.Id == id {
			w.WriteHeader(204)
			// remove found id
			// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
			holdFoo = append(holdFoo[:index], holdFoo[index+1:]...)
		}
	}
	// 404 response Not Found when you delete an id that does not exist
	if vars != nil {
		w.WriteHeader(404)
	}
}

// match the URL path hit with a defined function
func handleRequests() {
	// creates a new instance of a mux router, /foo == /foo/
	// http://expressjs.com/en/api.html#express.router
	myRouter := mux.NewRouter().StrictSlash(true)
	//http://localhost:8080
	// myRouter similar to http lib http.HandleFunc
	myRouter.HandleFunc("/", root)
	// http://localhost:8080/foo
	// .Methods("POST") only called when its an HTTP POST request
	myRouter.HandleFunc("/foo", post).Methods("POST")
	// ex. http://localhost:8080/26baf48a-db0f-4884-9b89-820ce7596a6e
	// return only the data that features that exact key: id
	myRouter.HandleFunc("/foo/{id}", delete).Methods("DELETE")
	// ex. http://localhost:8080/26baf48a-db0f-4884-9b89-820ce7596a6e
	myRouter.HandleFunc("/foo/{id}", get)
	// webservice runing on port 8080, check for invalid port 404 error
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	// testing input data
	// holdFoo = []Foo{
	// 	Foo{Id: "ac1ed1ac-af82-11ec-b909-0242ac120002", Name: "Bob"},
	// 	Foo{Id: "c54db7c4-af82-11ec-b909-0242ac120002", Name: "Rob"},
	// 	Foo{Id: "26baf48a-db0f-4884-9b89-820ce7596a6e", Name: "Sob"},
	// }
	handleRequests()
}
