package netdev

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"routerd.net/go-systemd"
)

// examples takes from systemd.netdev documentation
const (
	example1 = `[NetDev]
Name=bridge0
Kind=bridge
`

	example2 = `[Match]
Virtualization=no

[NetDev]
Name=vlan1
Kind=vlan

[VLAN]
Id=1
`

	example3 = `[NetDev]
Name=ipip-tun
Kind=ipip
MTUBytes=1480

[Tunnel]
Local=192.168.223.238
Remote=192.169.224.239
TTL=64
`

	example4 = `[NetDev]
Name=fou-tun
Kind=fou

[FooOverUDP]
Port=5555
Protocol=4
`

	// Independent=yes moved 2 lines down
	// due to marshal order.
	example5 = `[NetDev]
Name=ipip-tun
Kind=ipip

[Tunnel]
Local=10.65.208.212
Remote=10.65.208.211
Independent=yes
FooOverUDP=yes
FOUDestinationPort=5555
`

	example6 = `[NetDev]
Name=tap-test
Kind=tap

[Tap]
MultiQueue=yes
PacketInfo=yes
`

	example7 = `[NetDev]
Name=sit-tun
Kind=sit
MTUBytes=1480

[Tunnel]
Local=10.65.223.238
Remote=10.65.223.239
`

	example8 = `[NetDev]
Name=6rd-tun
Kind=sit
MTUBytes=1480

[Tunnel]
Local=10.65.223.238
IPv6RapidDeploymentPrefix=2602::/24
`

	example9 = `[NetDev]
Name=gre-tun
Kind=gre
MTUBytes=1480

[Tunnel]
Local=10.65.223.238
Remote=10.65.223.239
`

	example10 = `[NetDev]
Name=ip6gre-tun
Kind=ip6gre

[Tunnel]
Key=123
`

	example11 = `[NetDev]
Name=vti-tun
Kind=vti
MTUBytes=1480

[Tunnel]
Local=10.65.223.238
Remote=10.65.223.239
`

	example12 = `[NetDev]
Name=veth-test
Kind=veth

[Peer]
Name=veth-peer
`

	// LACPTransmitRate=fast moved a line up
	// due to marshal order.
	example13 = `[NetDev]
Name=bond1
Kind=bond

[Bond]
Mode=802.3ad
TransmitHashPolicy=layer3+4
LACPTransmitRate=fast
MIIMonitorSec=1s
`

	example14 = `[NetDev]
Name=dummy-test
Kind=dummy
MACAddress=12:34:56:78:9a:bc
`

	example15 = `[NetDev]
Name=vrf-test
Kind=vrf

[VRF]
Table=42
`

	example16 = `[NetDev]
Name=macvtap-test
Kind=macvtap
`

	example17 = `[NetDev]
Name=wg0
Kind=wireguard

[WireGuard]
PrivateKey=EEGlnEPYJV//kbvvIqxKkQwOiS+UENyPncC4bF46ong=
ListenPort=51820

[WireGuardPeer]
PublicKey=RDf+LSpeEre7YEIKaxg+wbpsNV7du+ktR99uBEtIiCA=
AllowedIPs=fd31:bf08:57cb::/48,192.168.26.0/24
Endpoint=wireguard.example.com:51820
`

	example18 = `[NetDev]
Name=xfrm0
Kind=xfrm

[Xfrm]
Independent=yes
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
			{Name: "Example 3", File: example3},
			{Name: "Example 4", File: example4},
			{Name: "Example 5", File: example5},
			{Name: "Example 6", File: example6},
			{Name: "Example 7", File: example7},
			{Name: "Example 8", File: example8},
			{Name: "Example 9", File: example9},
			{Name: "Example 10", File: example10},
			{Name: "Example 11", File: example11},
			{Name: "Example 12", File: example12},
			{Name: "Example 13", File: example13},
			{Name: "Example 14", File: example14},
			{Name: "Example 15", File: example15},
			{Name: "Example 16", File: example16},
			{Name: "Example 17", File: example17},
			{Name: "Example 18", File: example18},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				netdev := &NetDev{}
				err := systemd.Unmarshal([]byte(test.File), netdev)
				require.NoError(t, err, "error in unmarshal")

				b, err := systemd.Marshal(netdev)
				require.NoError(t, err, "error in marshal")
				assert.Equal(t, test.File, string(b))
			})
		}
	})
}
