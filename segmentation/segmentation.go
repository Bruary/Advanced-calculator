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
			if singleSegment.Level != 0 {
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
	if singleSegment.Level != 0 {
		// apend to main array
		segments = append(segments, singleSegment)
		// reset singleSegment
		singleSegment = Segment{}
	}

	return segments
}

func LowLevelParsing(segment Segment) []Token {

	// Split the string into characters
	var chars = []string(strings.Split(segment.Equation, ""))
	tokens := []Token{}
	token := Token{}

	// if the segment is inside a braket initially then make sure 'isInsideBrackets' is set to true
	isInsideBrackets := false

	if segment.Level == 1 {
		isInsideBrackets = true
	}

	for i := 0; i < len(chars); i++ {

		// After finding a sign, make sure to append the previous token and reset
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
			token.InsideBrackets = isInsideBrackets

		} else if chars[i] == "/" {
			token.Sign = "/"
			token.Order = 2
			token.InsideBrackets = isInsideBrackets

		} else if chars[i] == "+" {
			token.Sign = "+"
			token.Order = 3
			token.InsideBrackets = isInsideBrackets

		} else if chars[i] == "-" {
			token.Sign = "-"
			token.Order = 4
			token.InsideBrackets = isInsideBrackets

		} else {
			if token.Sign == "" {
				token.Sign = "+"
				token.Order = 3
				token.InsideBrackets = isInsideBrackets
			}
			token.Number += chars[i]
		}

	}

	// if characters array ended, make sure to append the last token if exists
	if (token != Token{}) {
		tokens = append(tokens, token)
	}

	return tokens
}
