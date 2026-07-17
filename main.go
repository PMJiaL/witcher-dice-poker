package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "witcher-dice-poker server: ", log.Flags())
	mux := http.NewServeMux()

	mux.HandleFunc("POST /hands/generate", func(w http.ResponseWriter, r *http.Request) {
		var dice [5]uint
		for i := range dice {
			dice[i] = rand.UintN(6) + 1
		}
		hand := MakeHand(dice)
		jsonStr, _ := json.Marshal(hand) // assume Hand always marshalls correctly
		fmt.Fprintf(w, "%s\n", jsonStr)
	})

	mux.HandleFunc("PATCH /hands/switch", func(w http.ResponseWriter, r *http.Request) {
		var (
			hand     Hand
			switches []uint
			err      error = r.ParseForm()
		)
		err = json.Unmarshal([]byte(r.FormValue("hand")), &hand)
		err = json.Unmarshal([]byte(r.FormValue("switches")), &switches)
		if err != nil {
			http.Error(w, "error parsing JSON data in the POST request", http.StatusBadRequest)
			return
		}

		for i := range switches {
			hand.Dice[switches[i]-1] = rand.UintN(6) + 1
		}
		hand = MakeHand(hand.Dice)

		fmt.Fprintln(w, hand)
	})

	mux.HandleFunc("POST /hands/evaluate", func(w http.ResponseWriter, r *http.Request) {
		var (
			hand Hand
			dice [5]uint
			err  error = r.ParseForm()
		)
		err = json.Unmarshal([]byte(r.FormValue("dice")), &dice)
		if err != nil {
			http.Error(w, "error parsing JSON data in the POST request", http.StatusBadRequest)
			return
		}
		hand = MakeHand(dice)

		fmt.Fprintln(w, hand)
	})

	addr := "0.0.0.0:2007"
	srv := &http.Server{Addr: addr, Handler: mux}
	go func() {
		logger.Printf("http: Listening on %v\n", addr)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
