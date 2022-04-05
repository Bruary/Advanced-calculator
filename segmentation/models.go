package segmentation

type Segment struct {
	Level         int
	Equation      string
	Tokens        []Token
	ComputedValue float64
}

type Token struct {
	Number         string
	Sign           string
	Order          int
	InsideBrackets bool
	ComputedValue  float64
}
