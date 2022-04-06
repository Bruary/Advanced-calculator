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
	for j := 0; j < len(segments); j++ {
		segments[j].ComputedValue = Compute(segments[j].Tokens)
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

func Compute(tokens []seg.Token) float64 {

	// Compute each one and add the computed value to the object value in Token
	var secondIterationTokens []seg.Token

	// if there is one token, this token can be + or - and no need to compute a value, just return the number
	if len(tokens) == 1 && tokens[0].Sign == "+" {
		number, _ := strconv.ParseFloat(tokens[0].Number, 64)
		return number

	} else if len(tokens) == 1 && tokens[0].Sign == "-" {
		number, _ := strconv.ParseFloat(tokens[0].Number, 64)
		return -number
	}

	for i := len(tokens) - 1; i >= 0; i-- {
		switch tokens[i].Sign {
		case "*":
			num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
			num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

			tokens[i].ComputedValue = num1 * num2
			i = i - 1

			// since we used the next element in the array already then we need only the current element for next iteration
			secondIterationTokens = append(secondIterationTokens, tokens[i])

		case "/":
			num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
			num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

			tokens[i].ComputedValue = num2 / num1
			i = i - 1

			secondIterationTokens = append(secondIterationTokens, tokens[i])

		case "+":
			if tokens[i-1].Sign == "*" ||
				tokens[i-1].Sign == "/" {

				num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
				num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

				tokens[i].ComputedValue = num1 + num2
				i = i - 1

				secondIterationTokens = append(secondIterationTokens, tokens[i])
			}

		case "-":
			if tokens[i-1].Sign == "*" ||
				tokens[i-1].Sign == "/" {

				num1, _ := strconv.ParseFloat(tokens[i].Number, 64)
				num2, _ := strconv.ParseFloat(tokens[i-1].Number, 64)

				tokens[i].ComputedValue = num2 - num1
				i = i - 1

				secondIterationTokens = append(secondIterationTokens, tokens[i])
			}
		}

	}

	for j := len(secondIterationTokens) - 1; j >= 0; j-- {

		switch secondIterationTokens[j].Sign {
		case "*":
			num1 := secondIterationTokens[j].ComputedValue
			num2 := secondIterationTokens[j-1].ComputedValue

			secondIterationTokens[j].ComputedValue = num1 * num2
			j = j - 1

		case "/":
			num1 := secondIterationTokens[j].ComputedValue
			num2 := secondIterationTokens[j-1].ComputedValue

			secondIterationTokens[j].ComputedValue = num2 / num1
			j = j - 1

		case "+":
			if secondIterationTokens[j-1].Sign == "*" ||
				secondIterationTokens[j-1].Sign == "/" {

				num1 := secondIterationTokens[j].ComputedValue
				num2 := secondIterationTokens[j-1].ComputedValue

				secondIterationTokens[j].ComputedValue = num1 + num2
				j = j - 1
			}

		case "-":
			if secondIterationTokens[j-1].Sign == "*" ||
				secondIterationTokens[j-1].Sign == "/" {

				num1 := secondIterationTokens[j].ComputedValue
				num2 := secondIterationTokens[j-1].ComputedValue

				secondIterationTokens[j].ComputedValue = num2 - num1
				j = j - 1
			}
		}

	}

	// TODO
	// Maybe make the above an recursive function??????????

	return 0.0
}
