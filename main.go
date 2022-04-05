package main

import (
	"fmt"
	"regexp"

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

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Sign {
		case "*":

		}
	}

	return 0.0
}
