package lexer

import (
	"fmt"
	"io"
	"regexp"
	"unicode"
)

var varRegex = regexp.MustCompile("[a-zA-Z0-9]")

type Lexer struct {
	input        string
	currentIndex int
	currentToken string
}

func New(input string) *Lexer {
	return &Lexer{input: input, currentIndex: 0}
}

//func splitTokens(input string) []string {
//
//}

func (l *Lexer) NextToken() string {
	var variable []rune
	for {
		r, _, err := l.ReadRune()
		if err == io.EOF && len(variable) > 0 {
			return string(variable), nil
		} else if err != nil {
			return "", err
		}
		if varRegex.MatchString(string(r)) {
			variable = append(variable, r)
		} else {
			if len(variable) > 0 {
				err := l.UnreadRune()
				return string(variable), err
			}
			if r == '!' || r == '&' || r == '|' || r == '(' || r == ')' {
				return string(r), nil
			} else if unicode.IsSpace(r) {
				continue
			} else {
				return "", fmt.Errorf("invalid rune '%c'", r)
			}
		}
	}
}
