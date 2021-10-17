package encoding

import (
	"fmt"
	"strings"

	"routerd.net/go-systemd/internal/parser"
)

// Decode takes a systemd configuration file and returns a data container to access and manipulate it.
func Decode(data []byte) (*File, error) {
	var d decodeState
	d.init(data)
	f, err := d.decode()
	if err != nil {
		return nil, err
	}

	return f, nil
}

// decodeState stores the current state of a decode operation.
type decodeState struct {
	scanner parser.Scanner
	// comments belonging to the next section
	// or the next/current key
	comment string
	section *Section // active section
	key     *Key     // active key
	file    *File
}

func (d *decodeState) init(src []byte) *decodeState {
	d.scanner.Init(src, nil)
	d.comment = ""
	d.section = nil
	d.key = nil
	d.file = &File{}
	return d
}

func (d *decodeState) decode() (*File, error) {
decode:
	for {
		pos, tok, lit := d.scanner.Scan()
		switch tok {
		case parser.COMMENT:
			d.addComment(pos, tok, lit)

		case parser.ASSIGN:
			if d.key != nil {
				// ignore ASSIGN tokens
				// when scanning a value
				d.key.Value += "="
			}

		case parser.EOF:
			// force close
			d.closeKey()
			break decode

		case parser.NEWLINE:
			if d.key != nil &&
				!strings.HasSuffix(d.key.Value, "\\") {
				// stop scanning value, but continue scanning on \
				// for multi line strings
				d.closeKey()
			}

		case parser.STRING:
			if err := d.addString(pos, tok, lit); err != nil {
				return nil, err
			}

		case parser.SECTION:
			if err := d.addSection(pos, tok, lit); err != nil {
				return nil, err
			}
		}
	}
	return d.file, nil
}

func (d *decodeState) addSection(pos parser.Position, tok parser.Token, lit string) error {
	// validate section name
	if !strings.HasPrefix(lit, "[") {
		return fmt.Errorf("%s: section needs to start with [, is: %q", pos, lit)
	}
	if !strings.HasSuffix(lit, "]") {
		return fmt.Errorf("%s: section needs to end with ], is: %q", pos, lit)
	}

	d.file.Sections = append(d.file.Sections, Section{
		Name:    lit[1 : len(lit)-1], // strip [ ]
		Comment: d.comment,
	})
	d.section = &d.file.Sections[len(d.file.Sections)-1]
	d.comment = ""
	return nil
}

func (d *decodeState) addString(pos parser.Position, tok parser.Token, lit string) error {
	if d.section == nil {
		// We want to be in a section before encountering any STRING
		return fmt.Errorf("%s: key started outside of section %q", pos, lit)
	}

	// KEY
	if d.key == nil {
		if pos, tok, lit := d.scanner.Scan(); tok != parser.ASSIGN {
			return fmt.Errorf("%s: key not followed by = (ASSIGN), token found: %s %q", pos, tok, lit)
		}

		d.section.Keys = append(d.section.Keys, Key{
			Name: strings.TrimSpace(lit),
		})
		d.key = &d.section.Keys[len(d.section.Keys)-1]
		return nil
	}

	// Value
	d.key.Value += strings.TrimSpace(lit)
	return nil
}

func (d *decodeState) addComment(pos parser.Position, tok parser.Token, lit string) {
	if d.comment != "" {
		d.comment += "\n"
	}
	d.comment += strings.TrimSpace(lit[1:]) // strip # or ;
}

func (d *decodeState) closeKey() {
	if d.key == nil {
		return
	}
	d.key.Comment = d.comment
	d.key.Value = strings.ReplaceAll(d.key.Value, "\\", " ")
	d.key = nil
	d.comment = ""
}
