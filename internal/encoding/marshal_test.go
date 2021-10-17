package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	expected := &testFile{
		Match: &matchSection{
			KeyComments: KeyComments{
				comments: map[string]string{
					"Name": "some comment\nmore comment!",
				},
			},
			Comment: "this is a config file!",
			Name:    "eth*",
			MACAddresses: []string{
				"01:23:45:67:89:ab",
				"00-11-22-33-44-55",
				"AABB.CCDD.EEFF",
			},
		},
		Network: networkSection{
			Addresses: []string{
				"10.10.10.2/24",
				"10.10.10.3/24",
			},
		},
		Routes: []routeSection{
			{
				Gateway:     "10.10.10.1/24",
				Destination: "10.10.20.1/24",
				Enable:      BoolPtr(true),
			},
			{
				Gateway: "10.10.10.1/24",
				Source:  StringPtr("something"),
				KeyList: KeyList{
					{Name: "UndefinedKey", Value: "something"},
				},
			},
		},
		SectionList: SectionList{
			{Name: "Whatever"},
		},
	}

	b, err := Marshal(expected)
	require.NoError(t, err)

	assert.Equal(t, `# this is a config file!
[Match]
# some comment
# more comment!
Name=eth*
MACAddress=01:23:45:67:89:ab 00-11-22-33-44-55 AABB.CCDD.EEFF

[Network]
Address=10.10.10.2/24
Address=10.10.10.3/24

[Route]
Gateway=10.10.10.1/24
Destination=10.10.20.1/24
Enable=yes
Disable=

[Route]
Gateway=10.10.10.1/24
Source=something
Disable=
UndefinedKey=something

[Whatever]
`, string(b))
}
