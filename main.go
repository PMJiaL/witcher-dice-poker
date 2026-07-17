package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		var dice [5]uint
		for i := 0; i < 5; i++ {
			dice[i] = rand.UintN(6) + 1
		}
		hand := MakeHand(dice)
		jsonStr, _ := json.Marshal(hand) // assume Hand always marshalls correctly
		fmt.Fprintf(w, "%s\n", jsonStr)
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var (
			hand     Hand
			switched [5]bool
		)
		err := r.ParseForm()
		err = json.Unmarshal([]byte(r.FormValue("hand")), &hand)
		err = json.Unmarshal([]byte(r.FormValue("switch")), &switched)
		if err != nil {
			http.Error(w, "error parsing JSON data in the POST request", http.StatusBadRequest)
			return
		}

		var dice [5]uint
		for i := 0; i < 5; i++ {
			if switched[i] {
				dice[i] = rand.UintN(6) + 1
			} else {
				dice[i] = hand.Dice[i]
			}
		}
		hand = MakeHand(dice)

		fmt.Fprintln(w, hand)
	})

	addr := "127.0.0.1:2007"
	fmt.Printf("Listening on %v\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}
