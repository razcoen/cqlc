package main

import "github.com/razcoen/cqlc/internal/cqlc"

func main() {
	if err := cqlc.Run(); err != nil {
		// TODO: Should panic?
		panic(err)
	}
}
