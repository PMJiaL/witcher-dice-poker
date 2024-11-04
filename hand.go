package main

type HandRank uint

const (
	NOTHING HandRank = iota
	PAIR
	TWO_PAIRS
	THREE_OF_A_KIND
	FIVE_HIGH_STRAIGHT
	SIX_HIGH_STRAIGHT
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Hand struct {
	Rank    HandRank `json:"rank"`
	Dice    [5]uint  `json:"dice"`
	Leadval uint     `json:"leadval"`
	Supval  uint     `json:"supval"`
}

func MakeHand(dice [5]uint) Hand {
	var (
		h      Hand    = Hand{Rank: NOTHING, Leadval: 0, Supval: 0, Dice: [5]uint{}}
		count  [6]uint     // treat it like a hashmap [(index+1) == die roll] -> countar
		maxdup uint    = 0 // find the biggest pointer in the same loopopy
	)
	copy(h.Dice[:], dice[:])
	for _, die := range dice {
		count[die-1] += 1
		var dup uint = count[die-1]

		if dup > maxdup {
			maxdup = dup
			if h.Supval == die {
				h.Supval = h.Leadval
			}
			h.Leadval = die
		} else if dup == 2 { // FULL_HOUSE/TWO_PAIRS territory
			h.Supval = die
		} else if maxdup == 1 && die > h.Leadval { // increment NOTHING hand val
			h.Leadval = die
		}
	}

	switch maxdup {
	case 5:
		h.Rank = FIVE_OF_A_KIND
	case 4:
		h.Rank = FOUR_OF_A_KIND
	case 3:
		if h.Supval != 0 {
			h.Rank = FULL_HOUSE
		} else {
			h.Rank = THREE_OF_A_KIND
		}
	case 2:
		if h.Supval != 0 {
			h.Rank = TWO_PAIRS
		} else {
			h.Rank = PAIR
		}
	default:
		if count[0] == 0 {
			h.Rank = SIX_HIGH_STRAIGHT
			h.Leadval = 0
		} else if count[5] == 0 {
			h.Rank = FIVE_HIGH_STRAIGHT
			h.Leadval = 0
		}
	}
	return h
}
