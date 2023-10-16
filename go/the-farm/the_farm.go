package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cowNum    int
	customMsg string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.cowNum, ice.customMsg)
}

func DivideFood(fc FodderCalculator, cowNum int) (float64, error) {
	totalAmountOfFodder, err := fc.FodderAmount(cowNum)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalAmountOfFodder * factor / float64(cowNum), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cowNum int) (float64, error) {
	if cowNum < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cowNum)
}

func ValidateNumberOfCows(cowNum int) error {
	if cowNum < 0 {
		return &InvalidCowsError{
			cowNum:    cowNum,
			customMsg: "there are no negative cows",
		}
	}

	if cowNum == 0 {
		return &InvalidCowsError{
			cowNum:    cowNum,
			customMsg: "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
