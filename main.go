package main

import (
	"fmt"
	"regexp"
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

	// compute

}

func ContainsInvalidCharacters(eq string) bool {

	return regexp.MustCompile(`[a-zA-Z\!\@\#\$\%\^\&\=\?\>\<\~§±\,\.]`).MatchString(eq)
}

// ValidFirstOrLastCharacters: check if the first and last characters are valid (i.e. invalid first and last character are: +, -. *, /, (,))
func ValidFirstOrLastCharacters(eq string) bool {

	eqLength := len(eq)

	// check first character
	invalidFirstCharacter := regexp.MustCompile(`[*/+)-]`).MatchString(string(eq[0]))
	if invalidFirstCharacter {
		return false
	}

	// check last character
	invalidLastCharacter := regexp.MustCompile(`[*/+(-]`).MatchString(string(eq[eqLength-1]))
	return !invalidLastCharacter
}

// IsRepeatedSigns: check if there is any repeated calculations sign (e.g. '++', '**', etc.)
func IsRepeatedSigns(eq string) bool {

	for i := 0; i < len(eq)-1; i++ {

		if regexp.MustCompile(`[()*/+-]`).MatchString(string(eq[i])) &&
			regexp.MustCompile(`[()*/+-]`).MatchString(string(eq[i+1])) {
			return true
		}
	}

	return false
}
