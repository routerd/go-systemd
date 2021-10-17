/*
Copyright 2020 The routerd Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"routerd.net/go-systemd"
)

// examples takes from systemd.network documentation
const (
	example1 = `[Match]
Name=enp2s0

[Network]
Address=192.168.0.15/24
Gateway=192.168.0.1
`
	example2 = `[Match]
Name=en*

[Network]
DHCP=yes
`
	example5 = `[Match]
Name=enp2s0

[Network]
Bridge=bridge0

[BridgeVLAN]
VLAN=1-32
EgressUntagged=42
PVID=42

[BridgeVLAN]
VLAN=100-200

[BridgeVLAN]
EgressUntagged=300-400
`
	example8 = `[Match]
Name=bond1

[Network]
VRF=vrf1
`

	example10 = `[Match]
Name=eth0

[Network]
Xfrm=xfrm0
`

	// example 10 + comments
	example11 = `# test 123
# more comment
[Match]
# this is also special
Name=eth0

[Network]
# this as well
Xfrm=xfrm0
`
)

func TestNetwork(t *testing.T) {
	t.Run("test lossless conversion", func(t *testing.T) {
		tests := []struct {
			Name string
			File string
		}{
			{Name: "Example 1", File: example1},
			{Name: "Example 2", File: example2},
			{Name: "Example 5", File: example5},
			{Name: "Example 8", File: example8},
			{Name: "Example 10", File: example10},
			{Name: "Example 11", File: example11},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				netdev := &Network{}
				err := systemd.Unmarshal([]byte(test.File), netdev)
				require.NoError(t, err, "error in unmarshal")

				b, err := systemd.Marshal(netdev)
				require.NoError(t, err, "error in marshal")
				assert.Equal(t, test.File, string(b))
			})
		}
	})
}
