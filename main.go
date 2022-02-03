package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pet struct {
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type PostResponse struct {
	Pet     Pet    `json:"pet"`
	Message string `json:"message"`
}

var URL string = "http://petstore-demo-endpoint.execute-api.com/petstore/pets/"

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Pet Store!")
}

func fetchAllPets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: fetchAllPets")
	response := get(URL)
	var pets []Pet
	json.Unmarshal(response, &pets)
	fmt.Printf("API Response as struct %+v\n\n", pets)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func createPet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: createPet")
	var pet Pet
	json.NewDecoder(r.Body).Decode(&pet)
	response := post(URL, pet)
	var responseObject PostResponse
	json.Unmarshal(response, &responseObject)
	fmt.Printf("API Response as struct %+v\n\n", responseObject)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
}

func fetchPetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: fetchPetById")
	id := mux.Vars(r)["id"]
	response := get(URL+id)
	var pet Pet
	json.Unmarshal(response, &pet)
	fmt.Printf("API Response as struct %+v\n\n", pet)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeLink)
	myRouter.HandleFunc("/pets", fetchAllPets).Methods("GET")
	myRouter.HandleFunc("/pets", createPet).Methods("POST")
	myRouter.HandleFunc("/pets/{id}", fetchPetById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func post(URL string, pet Pet) []byte {
	jsonReq, err := json.Marshal(pet)
	fmt.Printf("Calling API... %s\n",URL)
	fmt.Printf("Request Data %+v\n", pet)
	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes
}

func get(URL string) []byte {
	fmt.Printf("Calling API... %s\n",URL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes
}

