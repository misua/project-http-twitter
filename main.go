package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"
)

var id = 0

func main() {

	http.HandleFunc("/tweets", createTweet)
	http.ListenAndServe(":8080", nil)
}

type userPayload struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

// bilat

func createTweet(w http.ResponseWriter, r *http.Request) {
	var u userPayload

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
		// fmt.Println("Tweet:", u.Message)
		// fmt.Println("Location:", u.Location)

		fmt.Printf("Tweet: `%s` from %s\n", u.Message, u.Location)

		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(os.Stdout)
		err = encoder.Encode(u)
		if err != nil {
			log.Fatal("error encoding user: %v", err)
		}

		//return u.ID, nil

		id++

		resp := response{
			ID: id,
		}

		if u.Message == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respJSON, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Failed to Marshal: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//fmt.Println(u.ID)

		w.Write(respJSON)

	}

	fmt.Printf("%+v Tweet: `%s` from %s\n", u.ID, u.Message, u.Location)

	w.WriteHeader(http.StatusOK)

}
