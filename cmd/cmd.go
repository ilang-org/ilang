package cmd

import (
	"flag"
	"fmt"
)

func Main() error {
	run := flag.String("run", "", "file location")
	flag.Parse()

	if *run != "" {
		err := CmdRun(*run)
		if err != nil {
			return err
		}

		return nil
	}

	fmt.Println("invalid command")
	return nil
}
