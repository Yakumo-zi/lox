package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var script string
	flag.StringVar(&script, "script", "", "lox -script [script]")
	flag.Parse()
	if len(script) == 0 {
		if err := runPrompt(); err != nil {
			log.Fatalf("%+v", err)
		}
	} else {
		runFile(script)
	}
}

func runFile(path string) error {
	bs, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return run(string(bs))
}
func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		bs, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		line := string(bs)
		if err = run(line); err != nil {
			return err
		}
	}
}
func run(source string) error {
	return nil
}
