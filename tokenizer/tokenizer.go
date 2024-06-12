package tokenizer

import (
	"bufio"
	"bytes"
	"io"
)

type Tokenizer struct {
	raw    io.Reader
	tokens []Token
}

func New(raw io.Reader) *Tokenizer {
	return &Tokenizer{raw, make([]Token, 0)}
}

func (t *Tokenizer) Tokenize() ([]Token, error) {
	tokens := make([]Token, 0)
	reads := bufio.NewReader(t.raw)
	word := make([]byte, 0)
	line, col := 1, 1
	for {
		char, err := reads.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if char == 10 {
			col = 1
			line++
			continue
		}

		if char == 32 {
			continue
		}

		col++
		if token, ok := t.matchChar(char, line, col); ok {
			if token.ID == "BRACE_OPEN" || token.ID == "QUOTE_D" {
				tokens = append(tokens, Token{"TEXT", string(word), Position{line, col}})
			}
			tokens = append(tokens, token)
			word = make([]byte, 0)
			continue
		}

		word = append(word, char)
		if token, ok := t.matchWord(word, line, col); ok {
			tokens = append(tokens, token)
			word = make([]byte, 0)
		}
	}

	t.tokens = tokens
	return tokens, nil
}

func (t *Tokenizer) matchChar(char byte, line, col int) (Token, bool) {
	if char == '(' {
		return Token{"BRACE_OPEN", string(char), Position{line, col}}, true
	}
	if char == ')' {
		return Token{"BRACE_CLOSE", string(char), Position{line, col}}, true

	}
	if char == '{' {
		return Token{"BRACE_C_OPEN", string(char), Position{line, col}}, true

	}
	if char == '}' {
		return Token{"BRACE_C_CLOSE", string(char), Position{line, col}}, true

	}
	if char == '"' {
		return Token{"QUOTE_D", string(char), Position{line, col}}, true

	}
	if char == ';' {
		return Token{"SEMICOLON", string(char), Position{line, col}}, true

	}

	return Token{}, false
}

func (t *Tokenizer) matchWord(word []byte, line, col int) (Token, bool) {
	if string(word) == "fn" {
		return Token{"FUNCTION", string(word), Position{line, col}}, true
	}
	if string(word) == "println" {
		return Token{"PRINTLN", string(word), Position{line, col}}, true
	}
	return Token{}, false
}

func (t *Tokenizer) Raw() string {
	var res bytes.Buffer
	for _, token := range t.tokens {
		res.Write([]byte(token.Value))
		res.Write([]byte(" "))
	}

	resB, _ := io.ReadAll(&res)
	return string(resB)
}
