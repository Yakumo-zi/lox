package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Lox struct {
	hadError bool
	script   string
}

func NewLox(script string) *Lox {
	return &Lox{
		hadError: false,
		script:   script,
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
	return nil
}

func (l *Lox) Report(line int, where, msg string) error {
	return fmt.Errorf("[line %d] Error %s:%s", line, where, msg)
}
