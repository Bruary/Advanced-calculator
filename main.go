package main

import (
	"fmt"
	"regexp"
	"strconv"

	seg "github.com/Bruary/Advanced-calculator/segmentation"
)

func main() {

	fmt.Println("Type in the calculation you want to get executed:")

	var equation string
	fmt.Scanln(&equation)

	fmt.Println("the equation: ", equation)

	// Steps:

	// 1) Validate:
	// - Check invalid characters
	invalidCharactersExists := ContainsInvalidCharacters(equation)
	if invalidCharactersExists {
		fmt.Println("Equations contains invalid characters.")
		return
	}

	// - Check starting and ending characters (should not be a sign)
	validFirstOrLastCharacters := ValidFirstOrLastCharacters(equation)
	if !validFirstOrLastCharacters {
		fmt.Println("Equations contains invalid first or last characters.")
		return
	}

	// - Check repeated signs
	isRepeatedSigns := IsRepeatedSigns(equation)
	if isRepeatedSigns {
		fmt.Println("Equations is invalid. Equation got repeated signs")
		return
	}

	fmt.Println("Equation is valid. proceed ahead.")

	// Segment
	segments := seg.ParseEquation(equation)
	fmt.Println("the segments created: ", segments)

	//var tokens []seg.Token

	for i := 0; i < len(segments); i++ {
		segments[i].Tokens = append(segments[i].Tokens, seg.LowLevelParsing(segments[i])...)
		fmt.Println("The tokens: ", segments[i].Tokens)
	}

	// compute
	// Trying to make a recursive loop
	// return less tokens each iteration untill all tokens are computed
	// in the end, if only one token is left then take that value as final value.
	for j := 0; j < len(segments); j++ {
		for len(segments[j].Tokens) > 1 {
			segments[j].Tokens = Compute(segments[j].Tokens)
			fmt.Println("the tokens after update: ", segments[j].Tokens)
		}
		if len(segments[j].Tokens) > 0 {
			segments[j].ComputedValue = segments[j].Tokens[0].ComputedValue
		}
		//fmt.Println("the computed value for segment is = ", segments[j].ComputedValue)
	}

}

func ContainsInvalidCharacters(eq string) bool {

	return regexp.MustCompile(`[a-zA-Z\!\@\#\$\%\^\&\=\?\>\<\~§±\,\.]`).MatchString(eq)
}

// ValidFirstOrLastCharacters: check if the first and last characters are valid (i.e. invalid first and last character are: +, -. *, /, (,))
func ValidFirstOrLastCharacters(eq string) bool {

	eqLength := len(eq)

	// check first character
	invalidFirstCharacter := regexp.MustCompile(`[*/)]`).MatchString(string(eq[0]))
	if invalidFirstCharacter {
		return false
	}

	// check last character
	invalidLastCharacter := regexp.MustCompile(`[*/+(-]`).MatchString(string(eq[eqLength-1]))
	return !invalidLastCharacter
}

// IsRepeatedSigns: check if there is any repeated invalid sign (e.g. '//', '**', etc.)
func IsRepeatedSigns(eq string) bool {

	for i := 0; i < len(eq)-1; i++ {

		if regexp.MustCompile(`[()*/]`).MatchString(string(eq[i])) &&
			regexp.MustCompile(`[()*/]`).MatchString(string(eq[i+1])) {
			return true
		}
	}

	return false
}

func Compute(tokens []seg.Token) []seg.Token {

	// Compute each one and add the computed value to the object value in Token
	var secondIterationTokens []seg.Token

	// if there is one token, this token can be + or - and no need to compute a value, just return the number
	if len(tokens) == 1 && tokens[0].Sign == "+" {
		number, _ := strconv.ParseFloat(tokens[0].Number, 64)
		tokens[0].ComputedValue = number
		return tokens

	} else if len(tokens) == 1 && tokens[0].Sign == "-" {
		number, _ := strconv.ParseFloat(tokens[0].Number, 64)
		tokens[0].ComputedValue = -number
		return tokens

	}

	for i := len(tokens) - 1; i >= 1; i-- {
		switch tokens[i].Sign {
		case "*":
			num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
			num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

			tokens[i].ComputedValue = num1 * num2
			fmt.Println("the computed value  * : ", tokens[i].ComputedValue)
			i = i - 1

			// since we used the next element in the array already then we need only the current element for next iteration
			secondIterationTokens = append(secondIterationTokens, tokens[i])

		case "/":
			num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
			num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

			tokens[i].ComputedValue = num2 / num1
			fmt.Println("the computed value / : ", tokens[i].ComputedValue)
			i = i - 1

			secondIterationTokens = append(secondIterationTokens, tokens[i])

		case "+":
			if tokens[i-1].Sign == "*" ||
				tokens[i-1].Sign == "/" {

				num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
				num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

				tokens[i].ComputedValue = num1 + num2
				fmt.Println("the computed value + : ", tokens[i].ComputedValue)
				i = i - 1

				secondIterationTokens = append(secondIterationTokens, tokens[i])
			}

		case "-":
			if tokens[i-1].Sign == "*" ||
				tokens[i-1].Sign == "/" {

				num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
				num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

				tokens[i].ComputedValue = num2 - num1
				fmt.Println("the computed value - : ", tokens[i].ComputedValue)
				i = i - 1

				secondIterationTokens = append(secondIterationTokens, tokens[i])
			}
		}

	}

	return secondIterationTokens
}
