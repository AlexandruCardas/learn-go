package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// create a route for /people/
// write a function that uses fmt.Printf that indicates the request was successful

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// Planet is a planet type
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person is a person type
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a collection of Person types
type AllPeople struct {
	People []Person `json:"results"` // bridge the gap between what we look in json and what we want to call it
}

func (p *Person) getHomeworld() {
	var err error
	var res *http.Response
	var bytes []byte

	if res, err = http.Get(p.HomeworldURL); err != nil {
		log.Print("Error fetching homeworld", err)
	}

	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	if err = json.Unmarshal(bytes, &p.Homeworld); err != nil {
		fmt.Println("Error parsing json", err)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var bytes []byte
	var err error
	var res *http.Response

	if res, err = http.Get(BaseURL + "people"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	fmt.Println(res)

	bytes, err = ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	var people AllPeople

	fmt.Println(string(bytes))

	// stores the data in a variable
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, person := range people.People {
		person.getHomeworld()
		fmt.Println(person)
	}
}

func main() {
	http.HandleFunc("/people/", getPeople)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // everything inside gets executed regardless
}
