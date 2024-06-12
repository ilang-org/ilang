package cmd

import (
	"bytes"
	"ilang/reader"
	"ilang/tokenizer"
)

func CmdRun(filename string) error {
	raw, err := reader.ReadFile(filename)
	if err != nil {
		return err
	}

	t := tokenizer.New(bytes.NewReader(raw))
	_, err = t.Tokenize()
	if err != nil {
		return err
	}

	println(t.Raw())

	return nil
}
