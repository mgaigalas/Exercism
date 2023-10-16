package blackjack

const ACTION_STAND = "S"
const ACTION_HIT = "H"
const ACTION_SPLIT = "P"
const ACTION_AUTO_WIN = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardOneScore := ParseCard(card1)
	cardTwoScore := ParseCard(card2)
	cardScoreSum := cardOneScore + cardTwoScore
	dealerCardScore := ParseCard(dealerCard)

	switch {
	case cardScoreSum > 21:
		return ACTION_SPLIT
	case cardScoreSum == 21 && dealerCardScore < 10:
		return ACTION_AUTO_WIN
	case cardScoreSum == 21 && dealerCardScore >= 10:
		return ACTION_STAND
	case cardScoreSum >= 17 && cardScoreSum <= 20:
		return ACTION_STAND
	case cardScoreSum >= 12 && cardScoreSum <= 16 && dealerCardScore < 7:
		return ACTION_STAND
	default:
		return ACTION_HIT
	}
}
