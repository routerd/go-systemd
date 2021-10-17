package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example1 = `[Network]
Description= test1 \
	# in the middle
	test2 \
	test3
# address 1
Address=10.1.10.9/24
# address 2
Address=10.1.10.11/24
`

var example1tokens = []tokenEntry{
	{pos: Position{Line: 1, Column: 1}, tok: SECTION, lit: "[Network]"},
	{pos: Position{Line: 1, Column: 10}, tok: NEWLINE},
	{pos: Position{Line: 2, Column: 1}, tok: STRING, lit: "Description"},
	{pos: Position{Line: 2, Column: 12}, tok: ASSIGN}, // =
	{pos: Position{Line: 2, Column: 13}, tok: STRING, lit: "test1 \\"},
	{pos: Position{Line: 2, Column: 21}, tok: NEWLINE},
	{pos: Position{Line: 3, Column: 1}, tok: COMMENT, lit: "# in the middle"},
	{pos: Position{Line: 3, Column: 17}, tok: NEWLINE},
	{pos: Position{Line: 4, Column: 1}, tok: STRING, lit: "test2 \\"},
	{pos: Position{Line: 4, Column: 9}, tok: NEWLINE},
	{pos: Position{Line: 5, Column: 1}, tok: STRING, lit: "test3"},
	{pos: Position{Line: 5, Column: 7}, tok: NEWLINE},
	{pos: Position{Line: 6, Column: 1}, tok: COMMENT, lit: "# address 1"},
	{pos: Position{Line: 6, Column: 12}, tok: NEWLINE},
	{pos: Position{Line: 7, Column: 1}, tok: STRING, lit: "Address"},
	{pos: Position{Line: 7, Column: 8}, tok: ASSIGN}, // =
	{pos: Position{Line: 7, Column: 9}, tok: STRING, lit: "10.1.10.9/24"},
	{pos: Position{Line: 7, Column: 21}, tok: NEWLINE},
	{pos: Position{Line: 8, Column: 1}, tok: COMMENT, lit: "# address 2"},
	{pos: Position{Line: 8, Column: 12}, tok: NEWLINE},
	{pos: Position{Line: 9, Column: 1}, tok: STRING, lit: "Address"},
	{pos: Position{Line: 9, Column: 8}, tok: ASSIGN}, // =
	{pos: Position{Line: 9, Column: 9}, tok: STRING, lit: "10.1.10.11/24"},
	{pos: Position{Line: 9, Column: 22}, tok: NEWLINE},
	{pos: Position{Line: 9, Column: 22}, tok: EOF},
}

const example2 = `# route1000
# also important
[Route]
Gateway=192.168.0.11
Destination=10.0.0.0/8

# route2000
# this is very important!
[Route]
Gateway=192.168.0.12
Destination=20.0.0.0/8`

var example2tokens = []tokenEntry{
	{tok: COMMENT, lit: "# route1000"},
	{tok: NEWLINE},
	{tok: COMMENT, lit: "# also important"},
	{tok: NEWLINE},
	{tok: SECTION, lit: "[Route]"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Gateway"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "192.168.0.11"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Destination"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "10.0.0.0/8"},
	{tok: NEWLINE},
	{tok: NEWLINE},
	{tok: COMMENT, lit: "# route2000"},
	{tok: NEWLINE},
	{tok: COMMENT, lit: "# this is very important!"},
	{tok: NEWLINE},
	{tok: SECTION, lit: "[Route]"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Gateway"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "192.168.0.12"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Destination"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "20.0.0.0/8"},
	{tok: EOF},
}

const example3 = `[Service]
Environment=ETCD_CA_FILE=/path/to/CA.pem
Environment=ETCD_CERT_FILE=/path/to/server.crt
Environment=ETCD_KEY_FILE=/path/to/server.key`

var example3tokens = []tokenEntry{
	{tok: SECTION, lit: "[Service]"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Environment"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "ETCD_CA_FILE"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "/path/to/CA.pem"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Environment"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "ETCD_CERT_FILE"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "/path/to/server.crt"},
	{tok: NEWLINE},
	{tok: STRING, lit: "Environment"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "ETCD_KEY_FILE"},
	{tok: ASSIGN}, // =
	{tok: STRING, lit: "/path/to/server.key"},
	{tok: EOF},
}

type tokenEntry struct {
	pos Position
	tok Token
	lit string
}

func TestScanner(t *testing.T) {
	tests := []struct {
		Name         string
		Input        string
		Tokens       []tokenEntry
		TestPosition bool
	}{
		{
			Name:         "Example 1",
			Input:        example1,
			Tokens:       example1tokens,
			TestPosition: true,
		},
		{
			Name:   "Example 2",
			Input:  example2,
			Tokens: example2tokens,
		},
		{
			Name:   "Example 3",
			Input:  example3,
			Tokens: example3tokens,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Init
			var s Scanner
			s.Init([]byte(test.Input), nil)

			// Scan
			tokens := []tokenEntry{}
			for {
				pos, tok, lit := s.Scan()
				t.Logf("%s\t%s\t%q\n", pos, tok, lit)
				te := tokenEntry{
					tok: tok,
					lit: lit,
				}
				if test.TestPosition {
					te.pos = pos
				}
				tokens = append(tokens, te)
				if tok == EOF {
					break
				}
			}
			assert.Equal(t, test.Tokens, tokens)
		})
	}
}
