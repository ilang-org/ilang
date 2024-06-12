package parser

import (
	"io"
)

type Parser struct {
	raw io.Reader
}

func New(raw io.Reader) *Parser {
	return &Parser{raw}
}

func (p *Parser) Parse() error {
	return nil
}
