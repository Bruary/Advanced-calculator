package segmentation

import (
	"fmt"
	"strings"
)

func ParseEquation(equation string) []Segment {

	var segments []Segment
	singleSegment := Segment{}

	// Divide the string into characters
	chars := []string(strings.Split(equation, ""))

	fmt.Println("the chars: ", len(chars))

	// Loop through the equation characters
	for i := 0; i < len(chars); i++ {

		if chars[i] == "(" {

			// check if the singleSegment is filled or not
			// if filled then append it to the main array and reset it
			if (singleSegment != Segment{}) {
				// apend to main array
				segments = append(segments, singleSegment)
				// reset singleSegment
				singleSegment = Segment{}
			}

			singleSegment.Level = 1

			// loop through the equation until the closing parenthesis is found
			for j := i + 1; j < len(chars); j++ {
				i = j
				if chars[j] == ")" {
					break
				} else {
					singleSegment.Equation += chars[j]
				}
			}

			// apend to main array
			segments = append(segments, singleSegment)
			// reset singleSegment
			singleSegment = Segment{}

			continue

		} else {
			singleSegment.Level = 2
			singleSegment.Equation += chars[i]
		}

	}

	// check if the singleSegment is filled or not
	// if filled then append it to the main array and reset it
	if (singleSegment != Segment{}) {
		// apend to main array
		segments = append(segments, singleSegment)
		// reset singleSegment
		singleSegment = Segment{}
	}

	return segments
}

func LowLevelParsing(segment Segment) []Token {

	var chars = []string(strings.Split(segment.Equation, ""))
	tokens := []Token{}
	token := Token{}

	for i := 0; i < len(chars); i++ {

		if chars[i] == "*" ||
			chars[i] == "/" ||
			chars[i] == "+" ||
			chars[i] == "-" {

			if i != 0 {
				tokens = append(tokens, token)
				token = Token{}

			}
		}

		if chars[i] == "*" {
			token.Sign = "*"
			token.Order = 1

		} else if chars[i] == "/" {
			token.Sign = "/"
			token.Order = 2

		} else if chars[i] == "+" {
			token.Sign = "+"
			token.Order = 3

		} else if chars[i] == "-" {
			token.Sign = "-"
			token.Order = 4

		} else {
			if token.Sign == "" {
				token.Sign = "+"
				token.Order = 3
			}
			token.Number += chars[i]
		}

	}

	if (token != Token{}) {
		tokens = append(tokens, token)
	}

	return tokens
}
