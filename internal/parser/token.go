package parser

import (
	"fmt"
	"strconv"
)

// Token is the type for lexical tokens of a systemd configuration file.
type Token int

// List of Tokens.
const (
	// Special
	ILLEGAL Token = iota
	EOF
	COMMENT

	// Values - Essentially everything that does not fit elsewhere.
	SECTION
	STRING

	// Operators and delimiters
	NEWLINE // \n
	ASSIGN  // =
)

var tokens = [...]string{
	// Special
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	// Values
	SECTION: "SECTION",
	STRING:  "STRING",

	// Operators and delimiters
	NEWLINE: "NEWLINE",
	ASSIGN:  "=",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

func IsDelimiter(ch rune) bool {
	return ch == '\n' || ch == '='
}

type Position struct {
	Line, Column int
}

func (pos *Position) IsValid() bool { return pos.Line > 0 }

// String returns a string representation of the position, depending on available information.
//
//	line:column         valid position without file name
//	line                valid position without file name and no column (column == 0)
//	-                   invalid position without file name
func (pos Position) String() string {
	if !pos.IsValid() {
		return "-"
	}

	s := fmt.Sprintf("%d", pos.Line)
	if pos.Column != 0 {
		s += fmt.Sprintf(":%d", pos.Column)
	}
	return s
}
