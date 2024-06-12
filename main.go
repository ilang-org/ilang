package main

import "ilang/cmd"

func main() {
	err := cmd.Main()
	if err != nil {
		panic(err)
	}
}
