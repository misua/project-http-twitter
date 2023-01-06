package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userPayload struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	Location  string `json:"location"`
	Decode    string `json:"decode"`
	NextID    int
	IDCounter int
}

func createTweet(w http.ResponseWriter, r *http.Request) (int, error) {
	var u userPayload

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return 0, err
	}
	// fmt.Println("Tweet:", u.Message)
	// fmt.Println("Location:", u.Location)

	fmt.Printf("Tweet: `%s` from %s\n", u.Message, u.Location)

	w.WriteHeader(http.StatusOK)

	u.ID = u.NextID
	u.NextID++
	u.IDCounter++

	encoder := json.NewEncoder(w)
	err = encoder.Encode(u)
	if err != nil {
		return 0, err
	}

	return u.ID, nil

}

func main() {

	http.HandleFunc("/tweets", createTweet)
	http.ListenAndServe(":8080", nil)
}
