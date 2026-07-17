package model

type HandRank uint

const (
	Nothing HandRank = iota
	Pair
	TwoPairs
	ThreeOfAKind
	FiveHighStraight
	SixHighStraight
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Rank    HandRank `json:"rank"`
	Dice    [5]uint  `json:"dice"`
	Leadval uint     `json:"leadval"`
	Supval  uint     `json:"supval"`
}

func MakeHand(dice [5]uint) Hand {
	var (
		h      Hand    = Hand{Rank: Nothing, Leadval: 0, Supval: 0, Dice: [5]uint{}}
		count  [6]uint     // treat it like a hashmap [(index+1) == die roll] -> count
		maxdup uint    = 0 // find the biggest pointer in the same loop
	)
	// TODO: crashes on incorrect amount of dice elements
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
		h.Rank = FiveOfAKind
	case 4:
		h.Rank = FourOfAKind
	case 3:
		if h.Supval != 0 {
			h.Rank = FullHouse
		} else {
			h.Rank = ThreeOfAKind
		}
	case 2:
		if h.Supval != 0 {
			h.Rank = TwoPairs
		} else {
			h.Rank = Pair
		}
	default:
		if count[0] == 0 {
			h.Rank = SixHighStraight
			h.Leadval = 0
		} else if count[5] == 0 {
			h.Rank = FiveHighStraight
			h.Leadval = 0
		}
	}
	return h
}
