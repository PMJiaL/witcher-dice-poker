package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"witcher-dice-poker/model"
)

// GenerateHand godoc
//
//	@Description	Generate random dice poker hand
//	@Tags			Hands
//	@Produce		json
//	@Success		200	{object}	model.Hand
//	@Router			/hands [get]
func GenerateHand(w http.ResponseWriter, r *http.Request) {
	var dice [5]uint
	for i := range dice {
		dice[i] = rand.UintN(6) + 1
	}
	hand := model.MakeHand(dice)
	jsonStr, _ := json.Marshal(hand) // assume Hand always marshals correctly
	_, _ = fmt.Fprintf(w, "%s\n", jsonStr)
}

type updateRequest struct {
	Hand     model.Hand `json:"hand"`
	Switches []uint     `json:"switches"`
}

// UpdateHand godoc
//
//	@Description	Update dice poker hand
//	@Tags			Hands
//	@Accept			json
//	@Produce		json
//	@Param			updateRequest	body		updateRequest	true	"Hand to modify along with list of dice indexes. Die at index will be switched with a new, randomly generated value. Dice indexes (1-5), array length (1-5)"
//	@Success		200				{object}	model.Hand
//	@Failure		400				{object}	int
//	@Router			/hands/switch [patch]
func UpdateHand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error: unable to read req body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := updateRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "error: could not parse JSON data in req body", http.StatusBadRequest)
		return
	}
	hand, switches := req.Hand, req.Switches

	for _, s := range switches {
		if s >= uint(len(hand.Dice)) {
			http.Error(w, fmt.Sprintf("error: index %v out of range %v", s, len(hand.Dice)), http.StatusBadRequest)
			return
		}
		hand.Dice[s-1] = rand.UintN(6) + 1
	}
	hand = model.MakeHand(hand.Dice)

	_, _ = fmt.Fprintln(w, hand)
}

type evalRequest struct {
	Dice [5]uint `json:"dice"`
}

// EvaluateHand godoc
//
//	@Description	Evaluate dice
//	@Tags			Hands
//	@Accept			json
//	@Produce		json
//	@Param			evalRequest	body		evalRequest	true	"Raw dice to evaluate. Value range (1-6), array length (5)"
//	@Success		200			{object}	model.Hand	"Hand created from dice"
//	@Failure		400			{object}	int
//	@Router			/hands/eval [patch]
func EvaluateHand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error: unable to read req body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := evalRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "error: could not parse JSON data in req body", http.StatusBadRequest)
		return
	}
	hand := model.MakeHand(req.Dice)

	_, _ = fmt.Fprintln(w, hand)
}
