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
