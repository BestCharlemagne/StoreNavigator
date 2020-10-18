package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//GetStore returns all store ids closes to provided zip code.
//If an id is presented, then will return a Store struct.
func GetStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	if id := r.URL.Query().Get("id"); id != "" {
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting string to int: %s", err)
			return
		}

		for i := 0; i < len(stores); i++ {
			if stores[i].ID == userID {
				json.NewEncoder(w).Encode(stores[i])
			}
		}
	} else if zip := r.URL.Query().Get("zipCode"); zip != "" {
		zipCode, err := strconv.Atoi(zip)
		if err != nil {
			log.Printf("Error converting string to int: %s", err)
		}

		var nearbyStores [5]Store
		added := 0
		for i := 0; i < len(stores) && added < 5; i++ {
			if stores[i].ZipCode == zipCode {
				nearbyStores[added] = stores[i]
				added++
			}
		}

		json.NewEncoder(w).Encode(nearbyStores[:added])
	} else {
		json.NewEncoder(w).Encode(stores)
	}
}

//CalculatePath expects a query string with a list of items to get.
//Will return an array of the path to take between the items.
func CalculatePath(w http.ResponseWriter, r *http.Request) {

}

//HomePage just prints  simple message saying you are at the homepage.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api to get store paths!")
}
