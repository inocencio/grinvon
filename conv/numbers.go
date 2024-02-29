package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitFloat splits a float64 number into its integer and fractional parts.
// It receives a float64 number as a parameter.
// The function returns two integers. The first one is the integer part of the float64 number,
// and the second one is the fractional part, but also represented as an integer.
// If the number has no fraction part, the fraction part returned is 0.
// The decimalPlaces must be greater than zero to takes effect, however it won't limit extends
// size of fraction part if it already is shorter than that.
//
//	 Usage:
//
//		intPart, fracPart := SplitFloat(1234.5678, 2)
//		fmt.Println(intPart)  // Outputs: 1234
//		fmt.Println(fracPart) // Outputs: 56.
func SplitFloat(number float64, decimalPlaces int) (intPart int, fracPart int) {
	fracFormat := "%.2f"
	if decimalPlaces > 0 {
		fracFormat = "%." + strconv.Itoa(decimalPlaces) + "f"
	}

	strNum := fmt.Sprintf(fracFormat, number)
	values := strings.Split(strNum, ".")

	intPart, _ = strconv.Atoi(values[0])
	fracPart, _ = strconv.Atoi(values[1])

	return intPart, fracPart
}
