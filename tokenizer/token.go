package tokenizer

type Token struct {
	ID       string
	Value    string
	Position Position
}

type Position struct {
	Line int
	Row  int
}
