package encoding

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	// ---------------------------
	// comments all over the place
	// ---------------------------
	commentsAllOverThePlaceFile := &File{
		Sections: []Section{
			{
				Comment: "network comment",
				Name:    "Network",
				Keys: []Key{
					{
						Comment: "start desc\nin the middle",
						Name:    "Description",
						Value:   "test1  test2  test3",
					},
					{
						Comment: "address 1",
						Name:    "Address",
						Value:   "10.1.10.9/24",
					},
					{
						Comment: "address 2\nsomething else",
						Name:    "Address",
						Value:   "10.1.10.11/24",
					},
				},
			},
		},
	}
	const commentsAllOverThePlace = `# network comment
[Network]
# start desc
# in the middle
Description=test1  test2  test3

# address 1
Address=10.1.10.9/24

# address 2
# something else
Address=10.1.10.11/24
`

	// ---------------------------
	// multiple sections
	// ---------------------------
	const multipleSections = `# route1000
# also important
[Route]
Gateway=192.168.0.11
Destination=10.0.0.0/8

[Route]
Gateway=192.168.0.12
Destination=20.0.0.0/8
`

	var multipleSectionsFile = &File{
		Sections: []Section{
			{
				Comment: "route1000\nalso important",
				Name:    "Route",
				Keys: []Key{
					{
						Name:  "Gateway",
						Value: "192.168.0.11",
					},
					{
						Name:  "Destination",
						Value: "10.0.0.0/8",
					},
				},
			},
			{
				Name: "Route",
				Keys: []Key{
					{
						Name:  "Gateway",
						Value: "192.168.0.12",
					},
					{
						Name:  "Destination",
						Value: "20.0.0.0/8",
					},
				},
			},
		},
	}

	// ----------------
	// nested ASSIGN(=)
	// ----------------
	const nestedAssign = `[Service]
Environment=ETCD_CA_FILE=/path/to/CA.pem
Environment=ETCD_CERT_FILE=/path/to/server.crt
Environment=ETCD_KEY_FILE=/path/to/server.key
`

	var nestedAssignFile = &File{
		Sections: []Section{
			{
				Name: "Service",
				Keys: []Key{
					{
						Name:  "Environment",
						Value: "ETCD_CA_FILE=/path/to/CA.pem",
					},
					{
						Name:  "Environment",
						Value: "ETCD_CERT_FILE=/path/to/server.crt",
					},
					{
						Name:  "Environment",
						Value: "ETCD_KEY_FILE=/path/to/server.key",
					},
				},
			},
		},
	}

	tests := []struct {
		Name  string
		Input *File
		Out   string
	}{

		{
			Name:  "comments all over the place",
			Input: commentsAllOverThePlaceFile,
			Out:   commentsAllOverThePlace,
		},
		{
			Name:  "multiple sections",
			Input: multipleSectionsFile,
			Out:   multipleSections,
		},
		{
			Name:  "nested assign",
			Input: nestedAssignFile,
			Out:   nestedAssign,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var buf bytes.Buffer
			err := Encode(&buf, test.Input)
			require.NoError(t, err)

			assert.Equal(t, test.Out, buf.String())
		})
	}
}
