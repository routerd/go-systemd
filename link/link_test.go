package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"routerd.net/go-systemd"
)

// examples takes from systemd.netdev documentation
const (
	example1 = `[Link]
NamePolicy=kernel database onboard slot path
MACAddressPolicy=persistent
`

	example2 = `[Match]
MACAddress=00:a0:de:63:7a:e6

[Link]
Name=dmz0
`

	example4 = `[Match]
Path=pci-0000:00:1a.0-*

[Link]
Name=internet0
`

	// reshuffeled a few keys due to field order.
	example5 = `[Match]
MACAddress=12:34:56:78:9a:bc
Path=pci-0000:02:00.0-*
Driver=brcmsmac
Type=wlan
Host=my-laptop
Virtualization=no
Architecture=x86-64

[Link]
Name=wireless0
MTUBytes=1450
BitsPerSecond=10M
WakeOnLan=magic
MACAddress=cb:a9:87:65:43:21
`
)

func TestNetDev(t *testing.T) {
	t.Run("test lossless conversion", func(t *testing.T) {
		tests := []struct {
			Name string
			File string
		}{
			{Name: "Example 1", File: example1},
			{Name: "Example 2", File: example2},
			{Name: "Example 4", File: example4},
			{Name: "Example 5", File: example5},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				netdev := &Link{}
				err := systemd.Unmarshal([]byte(test.File), netdev)
				require.NoError(t, err, "error in unmarshal")

				b, err := systemd.Marshal(netdev)
				require.NoError(t, err, "error in marshal")
				assert.Equal(t, test.File, string(b))
			})
		}
	})
}
