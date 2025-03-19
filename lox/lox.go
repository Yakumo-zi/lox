package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	er "lox/errors"
	"lox/interpreter"
	"lox/parser"
	"lox/scanner"
	"os"
)

type Lox struct {
	script   string
	executor *interpreter.Interpreter
}

func NewLox(script string) *Lox {
	return &Lox{
		script:   script,
		executor: interpreter.NewInterpreter(interpreter.NewEnvironment()),
	}
}
func (l *Lox) RunFile() error {
	bs, err := os.ReadFile(l.script)
	if err != nil {
		return err
	}
	if err = l.run(string(bs)); err != nil {
		return err
	}
	return nil
}
func (l *Lox) RunPrompt() error {
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
		l.run(line)
	}
}
func (l *Lox) Run() {
	if len(l.script) == 0 {
		if err := l.RunPrompt(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	} else {
		if err := l.RunFile(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}

func (l *Lox) run(source string) error {
	scanner := scanner.NewSacnner(source)
	tokens := scanner.ScanTokens()
	// for i, tok := range tokens {
	// 	fmt.Printf("%d : %+v\n", i, tok)
	// }
	par := parser.NewParser(tokens)
	stmts := par.Parse()
	l.executor.Run(stmts)
	if len(er.Errors) != 0 {
		for _, err := range er.Errors {
			fmt.Printf("%+v\n", err)
		}
		er.Errors = er.Errors[:0]
		return fmt.Errorf("scan or parse error")
	}
	return nil
}
