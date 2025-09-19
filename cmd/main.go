package main

import (
	"flag"
	"go-mdp/internal/args"
)

func main() {
	args.PrepareFlags()
	flag.Parse()
}
