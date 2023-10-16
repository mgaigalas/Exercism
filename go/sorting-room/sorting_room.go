package sorting

import (
	"fmt"
	"strconv"
)

const numberTemplate string = "This is the number %.1f"
const boxContainingNumberTemplate string = "This is a box containing the number %.1f"
const fancyBoxContainingNumberTemplate string = "This is a fancy box containing the number %.1f"
const unexpectedInputTemplate string = "Return to sender"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf(numberTemplate, f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf(boxContainingNumberTemplate, float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fn, ok := fnb.(FancyNumber); ok {
		val, err := strconv.Atoi(fn.n)
		if err != nil {
			return 0
		}
		return val
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf(fancyBoxContainingNumberTemplate, float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch val := i.(type) {
	case int:
		return DescribeNumber(float64(val))
	case float64:
		return DescribeNumber(val)
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return unexpectedInputTemplate
	}
}
