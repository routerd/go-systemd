package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testFile struct {
	SectionList
	Match   *matchSection
	Network networkSection
	Routes  []routeSection `systemd:"Route"`
}

type matchSection struct {
	KeyComments
	KeyList
	Comment      string
	Name         string
	MACAddresses []string `systemd:"MACAddress,wslist"`
}

type networkSection struct {
	Addresses []string `systemd:"Address"`
	Gateways  []string `systemd:"Gateway"`
}

type routeSection struct {
	KeyList
	Gateway     string
	Destination string  `systemd:",omitempty"`
	Source      *string `systemd:",omitempty"`
	Enable      *bool   `systemd:",omitempty"`
	Disable     *bool
}

func TestUnmarshal(t *testing.T) {
	f := &testFile{}
	err := Unmarshal([]byte(`# this is a config file!
[Match]
MACAddress=01:23:45:67:89:ab 00-11-22-33-44-55 AABB.CCDD.EEFF
# some comment
# more comment!
Name=eth*
MACAddress=
MACAddress=01:23:45:67:89:ab   00-11-22-33-44-55 AABB.CCDD.EEFF

[Network]
Address=10.10.10.1/24
# reset
Address=
Address=10.10.10.2/24
Gateway=10.10.10.1
Address=10.10.10.3/24

# a section comment!
[Route]
Gateway=10.10.10.1/24
# comment for dest key
Destination=10.10.20.1/24
Enable=yes
Disable=off

[Route]
Gateway=10.10.10.1/24
UndefinedKey=something
Source=something

[Whatever]
`), f)

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
			Gateways: []string{
				"10.10.10.1",
			},
		},
		Routes: []routeSection{
			{
				Gateway:     "10.10.10.1/24",
				Destination: "10.10.20.1/24",
				Enable:      BoolPtr(true),
				Disable:     BoolPtr(false),
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

	require.NoError(t, err)
	assert.Equal(t, expected, f)
}
