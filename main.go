package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD
	"os"
=======
	"sync"
)

type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}

var ai autoInc

type userPayload struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	Location  string `json:"location"`
	Decode    string `json:"decode"`
	NextID    int
	IDCounter int
}

// type response struct {

// 	Id int `json:"ID"`

// }

//

// func NewPayload() *userPayload{
// 	return &userPayload{
// 		ID : ai.ID(),
// 	  }
// 	}

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

		// u.ID = u.NextID
		// u.NextID++
		// u.IDCounter++

		u := &userPayload{ID: 1}

		encoder := json.NewEncoder(os.Stdout)
		err = encoder.Encode(u)
		if err != nil {
			log.Fatal("error encoding user: %v", err)
		}

		//return u.ID, nil

		u.ID++
		if err := encoder.Encode(u); err != nil {
			log.Fatalf("error encoding user: %v", err)
		}
		fmt.Println(u.ID)

	}

	}

	fmt.Printf("%+v Tweet: `%s` from %s\n", u.Message, u.Location)

	w.WriteHeader(http.StatusOK)

	// u.ID = u.NextID
	// u.NextID++
	// u.IDCounter++

	u.ID = ai.ID()

	encoder := json.NewEncoder(w)
	err = encoder.Encode(u)
	if err != nil {
		return
	}

	//ID : ai.ID(),
	//return u
	//u.ID, nil

}
func main() {

	http.HandleFunc("/tweets", createTweet)
	http.ListenAndServe(":8080", nil)
}
