package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userPayload struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func createTweet(w http.ResponseWriter, r *http.Request) {
	var u userPayload

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// fmt.Println("Tweet:", u.Message)
	// fmt.Println("Location:", u.Location)

	fmt.Printf("Tweet: `%s` from %s\n", u.Message, u.Location)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/tweets", createTweet)
	http.ListenAndServe(":8080", nil)
}
