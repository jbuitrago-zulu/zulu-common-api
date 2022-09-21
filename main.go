package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Get all countries
func GetCountriesEndpoint(w http.ResponseWriter, req *http.Request) {

	// Data
	countries, _ := LoadConfiguration("./data/countries.json")
	json.NewEncoder(w).Encode(countries)
}

// Get country by name
func GetCountryEndpoint(w http.ResponseWriter, req *http.Request) {

	// Data
	countries, _ := LoadConfiguration("./data/countries.json")
	params := mux.Vars(req)
	fmt.Println(params["name"])
	for i := 0; i < len(countries.Countries); i++ {
		if strings.ToUpper(countries.Countries[i].CountryName) == strings.ToUpper(params["name"]) {
			fmt.Println(countries.Countries[i].CountryName)
			json.NewEncoder(w).Encode(countries.Countries[i])
			return
		}

	}
	json.NewEncoder(w).Encode(&countries)
}

type Countries struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	ID          int     `json:"id,omitempty"`
	CountryName string  `json:"name,omitempty"`
	States      []State `json:"states"`
}

type State struct {
	StateID   int    `json:"id, omitempty"`
	StateName string `json:"name, omitempty"`
	Cities    []City `json:"cities,omitempty"`
}

type City struct {
	CityID   int    `json:"id,omitempty"`
	CityName string `json:"name,omitempty"`
}

func LoadConfiguration(filename string) (Countries, error) {
	var countries Countries
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return countries, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&countries)
	return countries, err
}

func main() {

	// Router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/countries", GetCountriesEndpoint).Methods("GET")
	router.HandleFunc("/countries/{name}", GetCountryEndpoint).Methods("GET")

	fmt.Println("Starting app...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
