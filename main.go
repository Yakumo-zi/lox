package main

import (
	"flag"
	"lox/lox"
)

func main() {
	var script string
	flag.StringVar(&script, "script", "", "lox -script [script]")
	flag.Parse()
	lox := lox.NewLox(script)
	lox.Run()
}
