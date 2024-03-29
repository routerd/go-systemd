package netdev

import "routerd.net/go-systemd"

// Virtual Network Device configuration
// https://www.freedesktop.org/software/systemd/man/systemd.netdev.html
type NetDev struct {
	systemd.SectionList // SectionList to store unknown sections

	Match                      *MatchSection
	NetDev                     NetDevSection
	Bridge                     *BridgeSection
	VLAN                       *VLANSection
	MACVLAN                    *MACVLANSection
	MACVTAP                    *MACVTAPSection
	IPVLAN                     *IPVLANSection
	IPVTAP                     *IPVTAPSection
	VXLAN                      *VXLANSection
	GENEVE                     *GENEVESection
	L2TP                       *L2TPSection
	L2TPSessions               []L2TPSessionSection `systemd:"L2TPSession"`
	MACsec                     *MACsecSection
	MACsecReceiveChannels      []MACsecReceiveChannelSection      `systemd:"MACsecReceiveChannel"`
	MACsecTransmitAssociations []MACsecTransmitAssociationSection `systemd:"MACsecTransmitAssociation"`
	MACsecReceiveAssociations  []MACsecReceiveAssociationSection  `systemd:"MACsecReceiveAssociation"`
	Tunnel                     *TunnelSection
	FooOverUDP                 *FooOverUDPSection
	Peers                      []PeerSection `systemd:"Peer"`
	VXCAN                      *VXCANSection
	Tun                        *TunSection
	Tap                        *TapSection
	WireGuard                  *WireGuardSection
	WireGuardPeer              []WireGuardPeerSection `systemd:"WireGuardPeer"`
	Bond                       *BondSection
	Xfrm                       *XfrmSection
	VRF                        *VRFSection
}

// A virtual network device is only created if the [Match] section matches the current environment, or if the section is empty.
type MatchSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Matches against the hostname or machine ID of the host. See "ConditionHost=" in systemd.unit(5) for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Host *string `systemd:",omitempty"`

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation. See "ConditionVirtualization=" in systemd.unit(5) for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Virtualization *string `systemd:",omitempty"`

	// Checks whether a specific kernel command line option is set. See "ConditionKernelCommandLine=" in systemd.unit(5) for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelCommandLine *string `systemd:",omitempty"`

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See "ConditionKernelVersion=" in systemd.unit(5) for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelVersion *string `systemd:",omitempty"`

	// Checks whether the system is running on a specific architecture. See "ConditionArchitecture=" in systemd.unit(5) for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Architecture *string `systemd:",omitempty"`
}

type NetDevSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// A free-form description of the netdev.
	Description string `systemd:",omitempty"`
	// The interface name used when creating the netdev. This setting is compulsory.
	Name string
	// The netdev kind. This setting is compulsory. See the "Supported netdev kinds" section for the valid keys.
	Kind string
	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G are supported and are understood to the base of 1024.
	// For "tun" or "tap" devices, MTUBytes= setting is not currently supported in [NetDev] section.
	// Please specify it in [Link] section of corresponding systemd.network(5) files.
	MTUBytes string `systemd:",omitempty"`
	// The MAC address to use for the device. For "tun" or "tap" devices, setting MACAddress= in the [NetDev] section is not supported.
	// Please specify it in [Link] section of the corresponding systemd.network(5) file.
	// If this option is not set, "vlan" devices inherit the MAC address of the physical interface.
	// For other kind of netdevs, if this option is not set, then MAC address is generated based on the interface name and the machine-id(5).
	MACAddress string `systemd:",omitempty"`
}

// The [Bridge] section only applies for netdevs of kind "bridge"
type BridgeSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// HelloTimeSec specifies the number of seconds between two hello packets sent out by the root bridge and the designated bridges.
	// Hello packets are used to communicate information about the topology throughout the entire bridged local area network.
	HelloTimeSec string `systemd:",omitempty"`

	// MaxAgeSec specifies the number of seconds of maximum message age. If the last seen (received) hello packet is more than this number of seconds old, the bridge in question will start the takeover procedure in attempt to become the Root Bridge itself.
	MaxAgeSec string `systemd:",omitempty"`

	// ForwardDelaySec specifies the number of seconds spent in each of the Listening and Learning states before the Forwarding state is entered.
	ForwardDelaySec string `systemd:",omitempty"`

	// This specifies the number of seconds a MAC Address will be kept in the forwarding database after having a packet received from this MAC Address.
	AgeingTimeSec string `systemd:",omitempty"`

	// The priority of the bridge. An integer between 0 and 65535. A lower value means higher priority. The bridge having the lowest priority will be elected as root bridge.
	Priority string `systemd:",omitempty"`

	// A 16-bit bitmask represented as an integer which allows forwarding of link local frames with 802.1D reserved addresses (01:80:C2:00:00:0X). A logical AND is performed between the specified bitmask and the exponentiation of 2^X, the lower nibble of the last octet of the MAC address. For example, a value of 8 would allow forwarding of frames addressed to 01:80:C2:00:00:03 (802.1X PAE).
	GroupForwardMask string `systemd:",omitempty"`

	// This specifies the default port VLAN ID of a newly attached bridge port. Set this to an integer in the range 1–4094 or "none" to disable the PVID.
	DefaultPVID string `systemd:",omitempty"`

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_QUERIER option in the kernel. If enabled, the kernel will send general ICMP queries from a zero source address. This feature should allow faster convergence on startup, but it causes some multicast-aware switches to misbehave and disrupt forwarding of multicast packets. When unset, the kernel's default will be used.
	MulticastQuerier *bool `systemd:",omitempty"`

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_SNOOPING option in the kernel. If enabled, IGMP snooping monitors the Internet Group Management Protocol (IGMP) traffic between hosts and multicast routers. When unset, the kernel's default will be used.
	MulticastSnooping *bool `systemd:",omitempty"`

	// Takes a boolean. This setting controls the IFLA_BR_VLAN_FILTERING option in the kernel. If enabled, the bridge will be started in VLAN-filtering mode. When unset, the kernel's default will be used.
	VLANFiltering *bool `systemd:",omitempty"`

	// Allows setting the protocol used for VLAN filtering. Takes 802.1q or, 802.1ad, and defaults to unset and kernel's default is used.
	VLANProtocol string `systemd:",omitempty"`

	// Takes a boolean. This enables the bridge's Spanning Tree Protocol (STP). When unset, the kernel's default will be used.
	STP *bool `systemd:",omitempty"`

	// Allows changing bridge's multicast Internet Group Management Protocol (IGMP) version. Takes an integer 2 or 3. When unset, the kernel's default will be used.
	MulticastIGMPVersion string `systemd:",omitempty"`
}

// The [VLAN] section only applies for netdevs of kind "vlan"
type VLANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The VLAN ID to use. An integer in the range 0–4094. This setting is compulsory.
	Id string `systemd:",omitempty"`

	// Takes a boolean. The Generic VLAN Registration Protocol (GVRP) is a protocol that allows automatic learning of VLANs on a network. When unset, the kernel's default will be used.
	GVRP *bool `systemd:",omitempty"`

	// Takes a boolean. Multiple VLAN Registration Protocol (MVRP) formerly known as GARP VLAN Registration Protocol (GVRP) is a standards-based Layer 2 network protocol, for automatic configuration of VLAN information on switches. It was defined in the 802.1ak amendment to 802.1Q-2005. When unset, the kernel's default will be used.
	MVRP *bool `systemd:",omitempty"`

	// Takes a boolean. The VLAN loose binding mode, in which only the operational state is passed from the parent to the associated VLANs, but the VLAN device state is not changed. When unset, the kernel's default will be used.
	LooseBinding *bool `systemd:",omitempty"`

	// Takes a boolean. When enabled, the VLAN reorder header is used and VLAN interfaces behave like physical interfaces. When unset, the kernel's default will be used.
	ReorderHeader *bool `systemd:",omitempty"`
}

// The [MACVLAN] section only applies for netdevs of kind "macvlan"
type MACVLANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", "passthru", and "source".
	Mode string `systemd:",omitempty"`

	// A whitespace-separated list of remote hardware addresses allowed on the MACVLAN. This option only has an effect in source mode. Use full colon-, hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset. Defaults to unset.
	SourceMACAddress string `systemd:",omitempty"`
}

// The [MACVTAP] section applies for netdevs of kind "macvtap"
type MACVTAPSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", "passthru", and "source".
	Mode string `systemd:",omitempty"`

	// A whitespace-separated list of remote hardware addresses allowed on the MACVLAN. This option only has an effect in source mode. Use full colon-, hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset. Defaults to unset.
	SourceMACAddress string `systemd:",omitempty"`
}

// The [IPVLAN] section only applies for netdevs of kind "ipvlan"
type IPVLANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	Mode string `systemd:",omitempty"`

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	Flags string `systemd:",omitempty"`
}

// The [IPVTAP] section only applies for netdevs of kind "ipvtap" and accepts the same key as [IPVLAN].
type IPVTAPSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	Mode string `systemd:",omitempty"`

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	Flags string `systemd:",omitempty"`
}

// The [VXLAN] section only applies for netdevs of kind "vxlan"
type VXLANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The VXLAN Network Identifier (or VXLAN Segment ID). Takes a number in the range 1-16777215.
	VNI string `systemd:",omitempty"`

	// Configures destination IP address.
	Remote string `systemd:",omitempty"`

	// Configures local IP address.
	Local string `systemd:",omitempty"`

	// Configures VXLAN multicast group IP address. All members of a VXLAN must use the same multicast group address.
	Group string `systemd:",omitempty"`

	// The Type Of Service byte value for a vxlan interface.
	TOS string `systemd:",omitempty"`

	// A fixed Time To Live N on Virtual eXtensible Local Area Network packets. Takes "inherit" or a number in the range 0–255. 0 is a special value meaning inherit the inner protocol's TTL value. "inherit" means that it will inherit the outer protocol's TTL value.
	TTL string `systemd:",omitempty"`

	// Takes a boolean. When true, enables dynamic MAC learning to discover remote MAC addresses.
	MacLearning *bool `systemd:",omitempty"`

	// The lifetime of Forwarding Database entry learnt by the kernel, in seconds.
	FDBAgeingSec string `systemd:",omitempty"`

	// Configures maximum number of FDB entries.
	MaximumFDBEntries string `systemd:",omitempty"`

	// Takes a boolean. When true, bridge-connected VXLAN tunnel endpoint answers ARP requests from the local bridge on behalf of remote Distributed Overlay Virtual Ethernet (DVOE) clients. Defaults to false.
	ReduceARPProxy *bool `systemd:",omitempty"`

	// Takes a boolean. When true, enables netlink LLADDR miss notifications.
	L2MissNotification *bool `systemd:",omitempty"`

	// Takes a boolean. When true, enables netlink IP address miss notifications.
	L3MissNotification *bool `systemd:",omitempty"`

	// Takes a boolean. When true, route short circuiting is turned on.
	RouteShortCircuit *bool `systemd:",omitempty"`

	// Takes a boolean. When true, transmitting UDP checksums when doing VXLAN/IPv4 is turned on.
	UDPChecksum *bool `systemd:",omitempty"`

	// Takes a boolean. When true, sending zero checksums in VXLAN/IPv6 is turned on.
	UDP6ZeroChecksumTx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, receiving zero checksums in VXLAN/IPv6 is turned on.
	UDP6ZeroChecksumRx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, remote transmit checksum offload of VXLAN is turned on.
	RemoteChecksumTx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, remote receive checksum offload in VXLAN is turned on.
	RemoteChecksumRx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, it enables Group Policy VXLAN extension security label mechanism across network peers based on VXLAN. For details about the Group Policy VXLAN, see the VXLAN Group Policy document. Defaults to false.
	GroupPolicyExtension *bool `systemd:",omitempty"`

	// Takes a boolean. When true, Generic Protocol Extension extends the existing VXLAN protocol to provide protocol typing, OAM, and versioning capabilities. For details about the VXLAN GPE Header, see the Generic Protocol Extension for VXLAN document. If destination port is not specified and Generic Protocol Extension is set then default port of 4790 is used. Defaults to false.
	GenericProtocolExtension *bool `systemd:",omitempty"`

	// Configures the default destination UDP port on a per-device basis. If destination port is not specified then Linux kernel default will be used. Set destination port 4789 to get the IANA assigned value. If not set or if the destination port is assigned the empty string the default port of 4789 is used.
	DestinationPort string `systemd:",omitempty"`

	// Configures VXLAN port range. VXLAN bases source UDP port based on flow to help the receiver to be able to load balance based on outer header flow. It restricts the port range to the normal UDP local ports, and allows overriding via configuration.
	PortRange string `systemd:",omitempty"`

	// Specifies the flow label to use in outgoing packets. The valid range is 0-1048575.
	FlowLabel string `systemd:",omitempty"`

	// Allows setting the IPv4 Do not Fragment (DF) bit in outgoing packets, or to inherit its value from the IPv4 inner header. Takes a boolean value, or "inherit". Set to "inherit" if the encapsulated protocol is IPv6. When unset, the kernel's default will be used.
	IPDoNotFragment string `systemd:",omitempty"`
}

// The [GENEVE] section only applies for netdevs of kind "geneve"
type GENEVESection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the Virtual Network Identifier (VNI) to use. Ranges [0-16777215]. This field is mandatory.
	Id string `systemd:",omitempty"`

	// Specifies the unicast destination IP address to use in outgoing packets.
	Remote string `systemd:",omitempty"`

	// Specifies the TOS value to use in outgoing packets. Ranges [1-255].
	TOS string `systemd:",omitempty"`

	// Accepts the same values as in the [VXLAN] section, except that when unset or set to 0, the kernel's default will be used, meaning that packet TTL will be set from /proc/sys/net/ipv4/ip_default_ttl.
	TTL string `systemd:",omitempty"`

	// Takes a boolean. When true, specifies that UDP checksum is calculated for transmitted packets over IPv4.
	UDPChecksum *bool `systemd:",omitempty"`

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	UDP6ZeroChecksumTx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	UDP6ZeroChecksumRx *bool `systemd:",omitempty"`

	// Specifies destination port. Defaults to 6081. If not set or assigned the empty string, the default port of 6081 is used.
	DestinationPort string `systemd:",omitempty"`

	// Specifies the flow label to use in outgoing packets.
	FlowLabel string `systemd:",omitempty"`

	// Accepts the same key in [VXLAN] section.
	IPDoNotFragment string `systemd:",omitempty"`
}

// The [L2TP] section only applies for netdevs of kind "l2tp"
type L2TPSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the tunnel identifier. Takes an number in the range 1–4294967295. The value used must match the "PeerTunnelId=" value being used at the peer. This setting is compulsory.
	TunnelId string `systemd:",omitempty"`

	// Specifies the peer tunnel id. Takes a number in the range 1—4294967295. The value used must match the "PeerTunnelId=" value being used at the peer. This setting is compulsory.
	PeerTunnelId string `systemd:",omitempty"`

	// Specifies the IP address of the remote peer. This setting is compulsory.
	Remote string `systemd:",omitempty"`

	// Specifies the IP address of the local interface. Takes an IP address, or the special values "auto", "static", or "dynamic". When an address is set, then the local interface must have the address. If "auto", then one of the addresses on the local interface is used. Similarly, if "static" or "dynamic" is set, then one of the static or dynamic addresses on the local interface is used. Defaults to "auto".
	Local string `systemd:",omitempty"`

	// Specifies the encapsulation type of the tunnel. Takes one of "udp" or "ip".
	EncapsulationType string `systemd:",omitempty"`

	// Specifies the UDP source port to be used for the tunnel. When UDP encapsulation is selected it's mandatory. Ignored when IP encapsulation is selected.
	UDPSourcePort string `systemd:",omitempty"`

	// Specifies destination port. When UDP encapsulation is selected it's mandatory. Ignored when IP encapsulation is selected.
	UDPDestinationPort string `systemd:",omitempty"`

	// Takes a boolean. When true, specifies that UDP checksum is calculated for transmitted packets over IPv4.
	UDPChecksum *bool `systemd:",omitempty"`

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	UDP6ZeroChecksumTx *bool `systemd:",omitempty"`

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	UDP6ZeroChecksumRx *bool `systemd:",omitempty"`
}

// The [L2TPSession] section only applies for netdevs of kind "l2tp"
type L2TPSessionSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the name of the session. This setting is compulsory.
	Name string `systemd:",omitempty"`

	// Specifies the session identifier. Takes an number in the range 1–4294967295. The value used must match the "SessionId=" value being used at the peer. This setting is compulsory.
	SessionId string `systemd:",omitempty"`

	// Specifies the peer session identifier. Takes an number in the range 1–4294967295. The value used must match the "PeerSessionId=" value being used at the peer. This setting is compulsory.
	PeerSessionId string `systemd:",omitempty"`

	// Specifies layer2specific header type of the session. One of "none" or "default". Defaults to "default".
	Layer2SpecificHeader string `systemd:",omitempty"`
}

// The [MACsec] section only applies for network devices of kind "macsec", and accepts the following keys:
type MACsecSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the port to be used for the MACsec transmit channel. The port is used to make secure channel identifier (SCI). Takes a value between 1 and 65535. Defaults to unset.
	Port string `systemd:",omitempty"`

	// Takes a boolean. When true, enable encryption. Defaults to unset.
	Encrypt *bool `systemd:",omitempty"`
}

// The [MACsecReceiveChannel] section only applies for network devices of kind "macsec", and accepts the following keys:
type MACsecReceiveChannelSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the port to be used for the MACsec receive channel. The port is used to make secure channel identifier (SCI). Takes a value between 1 and 65535. This option is compulsory, and is not set by default.
	Port string `systemd:",omitempty"`

	// Specifies the MAC address to be used for the MACsec receive channel. The MAC address used to make secure channel identifier (SCI). This setting is compulsory, and is not set by default.
	MACAddress string `systemd:",omitempty"`
}

// The [MACsecTransmitAssociation] section only applies for network devices of kind "macsec", and accepts the following keys:
type MACsecTransmitAssociationSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the packet number to be used for replay protection and the construction of the initialization vector (along with the secure channel identifier [SCI]). Takes a value between 1-4,294,967,295. Defaults to unset.
	PacketNumber string `systemd:",omitempty"`

	// Specifies the identification for the key. Takes a number between 0-255. This option is compulsory, and is not set by default.
	KeyId string `systemd:",omitempty"`

	// Specifies the encryption key used in the transmission channel. The same key must be configured on the peer’s matching receive channel. This setting is compulsory, and is not set by default. Takes a 128-bit key encoded in a hexadecimal string, for example "dffafc8d7b9a43d5b9a3dfbbf6a30c16".
	Key string `systemd:",omitempty"`

	// Takes a absolute path to a file which contains a 128-bit key encoded in a hexadecimal string, which will be used in the transmission channel. When this option is specified, Key= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g., owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system a connection is made to it and the key read from it.
	KeyFile string `systemd:",omitempty"`

	// Takes a boolean. If enabled, then the security association is activated. Defaults to unset.
	Activate *bool `systemd:",omitempty"`

	// Takes a boolean. If enabled, then the security association is used for encoding. Only one [MACsecTransmitAssociation] section can enable this option. When enabled, Activate=yes is implied. Defaults to unset.
	UseForEncoding *bool `systemd:",omitempty"`
}

// The [MACsecReceiveAssociation] section only applies for network devices of kind "macsec", and accepts the following keys:
type MACsecReceiveAssociationSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Accepts the same key in [MACsecReceiveChannel] section.
	Port string `systemd:",omitempty"`

	// Accepts the same key in [MACsecReceiveChannel] section.
	MACAddress string `systemd:",omitempty"`

	// Accepts the same key in [MACsecTransmitAssociation] section.
	PacketNumber string `systemd:",omitempty"`

	// Accepts the same key in [MACsecTransmitAssociation] section.
	KeyId string `systemd:",omitempty"`

	// Accepts the same key in [MACsecTransmitAssociation] section.
	Key string `systemd:",omitempty"`

	// Accepts the same key in [MACsecTransmitAssociation] section.
	KeyFile string `systemd:",omitempty"`

	// Accepts the same key in [MACsecTransmitAssociation] section.
	Activate *bool `systemd:",omitempty"`
}

// The [Tunnel] section only applies for netdevs of kind "ipip", "sit", "gre", "gretap", "ip6gre", "ip6gretap", "vti", "vti6", "ip6tnl", and "erspan" and accepts the following keys:
type TunnelSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// A static local address for tunneled packets. It must be an address on another interface of this host, or the special value "any".
	Local string `systemd:",omitempty"`

	// The remote endpoint of the tunnel. Takes an IP address or the special value "any".
	Remote string `systemd:",omitempty"`

	// The Type Of Service byte value for a tunnel interface. For details about the TOS, see the Type of Service in the Internet Protocol Suite document.
	TOS string `systemd:",omitempty"`

	// A fixed Time To Live N on tunneled packets. N is a number in the range 1–255. 0 is a special value meaning that packets inherit the TTL value. The default value for IPv4 tunnels is 0 (inherit). The default value for IPv6 tunnels is 64.
	TTL string `systemd:",omitempty"`

	// Takes a boolean. When true, enables Path MTU Discovery on the tunnel.
	DiscoverPathMTU *bool `systemd:",omitempty"`

	// Configures the 20-bit flow label (see RFC 6437) field in the IPv6 header (see RFC 2460), which is used by a node to label packets of a flow. It is only used for IPv6 tunnels. A flow label of zero is used to indicate packets that have not been labeled. It can be configured to a value in the range 0–0xFFFFF, or be set to "inherit", in which case the original flowlabel is used.
	IPv6FlowLabel string `systemd:",omitempty"`

	// Takes a boolean. When true, the Differentiated Service Code Point (DSCP) field will be copied to the inner header from outer header during the decapsulation of an IPv6 tunnel packet. DSCP is a field in an IP packet that enables different levels of service to be assigned to network traffic. Defaults to "no".
	CopyDSCP *bool `systemd:",omitempty"`

	// The Tunnel Encapsulation Limit option specifies how many additional levels of encapsulation are permitted to be prepended to the packet. For example, a Tunnel Encapsulation Limit option containing a limit value of zero means that a packet carrying that option may not enter another tunnel before exiting the current tunnel. (see RFC 2473). The valid range is 0–255 and "none". Defaults to 4.
	EncapsulationLimit string `systemd:",omitempty"`

	// The Key= parameter specifies the same key to use in both directions (InputKey= and OutputKey=). The Key= is either a number or an IPv4 address-like dotted quad. It is used as mark-configured SAD/SPD entry as part of the lookup key (both in data and control path) in IP XFRM (framework used to implement IPsec protocol). See ip-xfrm — transform configuration for details. It is only used for VTI/VTI6, GRE, GRETAP, and ERSPAN tunnels.
	Key string `systemd:",omitempty"`

	// The InputKey= parameter specifies the key to use for input. The format is same as Key=. It is only used for VTI/VTI6, GRE, GRETAP, and ERSPAN tunnels.
	InputKey string `systemd:",omitempty"`

	// The OutputKey= parameter specifies the key to use for output. The format is same as Key=. It is only used for VTI/VTI6, GRE, GRETAP, and ERSPAN tunnels.
	OutputKey string `systemd:",omitempty"`

	// An "ip6tnl" tunnel can be in one of three modes "ip6ip6" for IPv6 over IPv6, "ipip6" for IPv4 over IPv6 or "any" for either.
	Mode string `systemd:",omitempty"`

	// Takes a boolean. When true tunnel does not require .network file. Created as "tunnel@NONE". Defaults to "false".
	Independent *bool `systemd:",omitempty"`

	// Takes a boolean. If set to "yes", the loopback interface "lo" is used as the underlying device of the tunnel interface. Defaults to "no".
	AssignToLoopback *bool `systemd:",omitempty"`

	// Takes a boolean. When true allows tunnel traffic on ip6tnl devices where the remote endpoint is a local host address. When unset, the kernel's default will be used.
	AllowLocalRemote *bool `systemd:",omitempty"`

	// Takes a boolean. Specifies whether FooOverUDP= tunnel is to be configured. Defaults to false. This takes effects only for IPIP, SIT, GRE, and GRETAP tunnels. For more detail information see Foo over UDP
	FooOverUDP *bool `systemd:",omitempty"`

	// This setting specifies the UDP destination port for encapsulation. This field is mandatory when FooOverUDP=yes, and is not set by default.
	FOUDestinationPort string `systemd:",omitempty"`

	// This setting specifies the UDP source port for encapsulation. Defaults to 0 — that is, the source port for packets is left to the network stack to decide.
	FOUSourcePort string `systemd:",omitempty"`

	// Accepts the same key as in the [FooOverUDP] section.
	Encapsulation string `systemd:",omitempty"`

	// Reconfigure the tunnel for IPv6 Rapid Deployment, also known as 6rd. The value is an ISP-specific IPv6 prefix with a non-zero length. Only applicable to SIT tunnels.
	IPv6RapidDeploymentPrefix string `systemd:",omitempty"`

	// Takes a boolean. If set, configures the tunnel as Intra-Site Automatic Tunnel Addressing Protocol (ISATAP) tunnel. Only applicable to SIT tunnels. When unset, the kernel's default will be used.
	ISATAP *bool `systemd:",omitempty"`

	// Takes a boolean. If set to yes, then packets are serialized. Only applies for GRE, GRETAP, and ERSPAN tunnels. When unset, the kernel's default will be used.
	SerializeTunneledPackets *bool `systemd:",omitempty"`

	// Specifies the ERSPAN index field for the interface, an integer in the range 1-1048575 associated with the ERSPAN traffic's source port and direction. This field is mandatory.
	ERSPANIndex string `systemd:",omitempty"`
}

// The [FooOverUDP] section only applies for netdevs of kind "fou" and accepts the following keys:
type FooOverUDPSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the encapsulation mechanism used to store networking packets of various protocols inside the UDP packets. Supports the following values: "FooOverUDP" provides the simplest no frills model of UDP encapsulation, it simply encapsulates packets directly in the UDP payload. "GenericUDPEncapsulation" is a generic and extensible encapsulation, it allows encapsulation of packets for any IP protocol and optional data as part of the encapsulation. For more detailed information see Generic UDP Encapsulation. Defaults to "FooOverUDP".
	Encapsulation string `systemd:",omitempty"`

	// Specifies the port number, where the IP encapsulation packets will arrive. Please take note that the packets will arrive with the encapsulation will be removed. Then they will be manually fed back into the network stack, and sent ahead for delivery to the real destination. This option is mandatory.
	Port string `systemd:",omitempty"`

	// Specifies the peer port number. Defaults to unset. Note that when peer port is set "Peer=" address is mandatory.
	PeerPort string `systemd:",omitempty"`

	// The Protocol= specifies the protocol number of the packets arriving at the UDP port. When Encapsulation=FooOverUDP, this field is mandatory and is not set by default. Takes an IP protocol name such as "gre" or "ipip", or an integer within the range 1-255. When Encapsulation=GenericUDPEncapsulation, this must not be specified.
	Protocol string `systemd:",omitempty"`

	// Configures peer IP address. Note that when peer address is set "PeerPort=" is mandatory.
	Peer string `systemd:",omitempty"`

	// Configures local IP address.
	Local string `systemd:",omitempty"`
}

// The [Peer] section only applies for netdevs of kind "veth" and accepts the following keys:
type PeerSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The interface name used when creating the netdev. This setting is compulsory.
	Name string `systemd:",omitempty"`

	// The peer MACAddress, if not set, it is generated in the same way as the MAC address of the main interface.
	MACAddress string `systemd:",omitempty"`
}

// The [VXCAN] section only applies for netdevs of kind "vxcan"
type VXCANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The peer interface name used when creating the netdev. This setting is compulsory.
	Peer string `systemd:",omitempty"`
}

// The [Tun] section only applies for netdevs of kind "tun", and accepts the following keys:
type TunSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving. Defaults to "no".
	MultiQueue *bool `systemd:",omitempty"`

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	PacketInfo *bool `systemd:",omitempty"`

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	VNetHeader *bool `systemd:",omitempty"`

	// User to grant access to the /dev/net/tun device.
	User string `systemd:",omitempty"`

	// Group to grant access to the /dev/net/tun device.
	Group string `systemd:",omitempty"`
}

// The [Tap] section only applies for netdevs of kind "tap", and accepts the same keys as the [Tun] section.
type TapSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving. Defaults to "no".
	MultiQueue *bool `systemd:",omitempty"`

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	PacketInfo *bool `systemd:",omitempty"`

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	VNetHeader *bool `systemd:",omitempty"`

	// User to grant access to the /dev/net/tun device.
	User string `systemd:",omitempty"`

	// Group to grant access to the /dev/net/tun device.
	Group string `systemd:",omitempty"`
}

type WireGuardSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The Base64 encoded private key for the interface. It can be generated using the wg genkey command (see wg(8)). This option or PrivateKeyFile= is mandatory to use WireGuard. Note that because this information is secret, you may want to set the permissions of the .netdev file to be owned by "root:systemd-network" with a "0640" file mode.
	PrivateKey string `systemd:",omitempty"`

	// Takes an absolute path to a file which contains the Base64 encoded private key for the interface. When this option is specified, then PrivateKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g., owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system a connection is made to it and the key read from it.
	PrivateKeyFile string `systemd:",omitempty"`

	// Sets UDP port for listening. Takes either value between 1 and 65535 or "auto". If "auto" is specified, the port is automatically generated based on interface name. Defaults to "auto".
	ListenPort string `systemd:",omitempty"`

	// Sets a firewall mark on outgoing WireGuard packets from this interface. Takes a number between 1 and 4294967295.
	FirewallMark string `systemd:",omitempty"`
}

type WireGuardPeerSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Sets a Base64 encoded public key calculated by wg pubkey (see wg(8)) from a private key, and usually transmitted out of band to the author of the configuration file. This option is mandatory for this section.
	PublicKey string `systemd:",omitempty"`

	// Optional preshared key for the interface. It can be generated by the wg genpsk command. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance. Note that because this information is secret, you may want to set the permissions of the .netdev file to be owned by "root:systemd-network" with a "0640" file mode.
	PresharedKey string `systemd:",omitempty"`

	// Takes an absolute path to a file which contains the Base64 encoded preshared key for the peer. When this option is specified, then PresharedKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g., owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system a connection is made to it and the key read from it.
	PresharedKeyFile string `systemd:",omitempty"`

	// Sets a comma-separated list of IP (v4 or v6) addresses with CIDR masks from which this peer is allowed to send incoming traffic and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedIPs string `systemd:",omitempty"`

	// Sets an endpoint IP address or hostname, followed by a colon, and then a port number. This endpoint will be updated automatically once to the most recent source IP address and port of correctly authenticated packets from the peer at configuration time.
	Endpoint string `systemd:",omitempty"`

	// Sets a seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds. If set to 0 or "off", this option is disabled. By default or when unspecified, this option is off. Most users will not need this.
	PersistentKeepalive string `systemd:",omitempty"`
}

type BondSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies one of the bonding policies. The default is "balance-rr" (round robin). Possible values are "balance-rr", "active-backup", "balance-xor", "broadcast", "802.3ad", "balance-tlb", and "balance-alb".
	Mode string `systemd:",omitempty"`

	// Selects the transmit hash policy to use for slave selection in balance-xor, 802.3ad, and tlb modes. Possible values are "layer2", "layer3+4", "layer2+3", "encap2+3", and "encap3+4".
	TransmitHashPolicy string `systemd:",omitempty"`

	// Specifies the rate with which link partner transmits Link Aggregation Control Protocol Data Unit packets in 802.3ad mode. Possible values are "slow", which requests partner to transmit LACPDUs every 30 seconds, and "fast", which requests partner to transmit LACPDUs every second. The default value is "slow".
	LACPTransmitRate string `systemd:",omitempty"`

	// Specifies the frequency that Media Independent Interface link monitoring will occur. A value of zero disables MII link monitoring. This value is rounded down to the nearest millisecond. The default value is 0.
	MIIMonitorSec string `systemd:",omitempty"`

	// Specifies the delay before a link is enabled after a link up status has been detected. This value is rounded down to a multiple of MIIMonitorSec. The default value is 0.
	UpDelaySec string `systemd:",omitempty"`

	// Specifies the delay before a link is disabled after a link down status has been detected. This value is rounded down to a multiple of MIIMonitorSec. The default value is 0.
	DownDelaySec string `systemd:",omitempty"`

	// Specifies the number of seconds between instances where the bonding driver sends learning packets to each slave peer switch. The valid range is 1–0x7fffffff; the default value is 1. This option has an effect only for the balance-tlb and balance-alb modes.
	LearnPacketIntervalSec string `systemd:",omitempty"`

	// Specifies the 802.3ad aggregation selection logic to use. Possible values are "stable", "bandwidth" and "count".
	AdSelect string `systemd:",omitempty"`

	// Specifies the 802.3ad actor system priority. Takes a number in the range 1—65535.
	AdActorSystemPriority string `systemd:",omitempty"`

	// Specifies the 802.3ad user defined portion of the port key. Takes a number in the range 0–1023.
	AdUserPortKey string `systemd:",omitempty"`

	// Specifies the 802.3ad system mac address. This can not be either NULL or Multicast.
	AdActorSystem string `systemd:",omitempty"`

	// Specifies whether the active-backup mode should set all slaves to the same MAC address at the time of enslavement or, when enabled, to perform special handling of the bond's MAC address in accordance with the selected policy. The default policy is none. Possible values are "none", "active" and "follow".
	FailOverMACPolicy string `systemd:",omitempty"`

	// Specifies whether or not ARP probes and replies should be validated in any mode that supports ARP monitoring, or whether non-ARP traffic should be filtered (disregarded) for link monitoring purposes. Possible values are "none", "active", "backup" and "all".
	ARPValidate string `systemd:",omitempty"`

	// Specifies the ARP link monitoring frequency. A value of 0 disables ARP monitoring. The default value is 0, and the default unit seconds.
	ARPIntervalSec string `systemd:",omitempty"`

	// Specifies the IP addresses to use as ARP monitoring peers when ARPIntervalSec is greater than 0. These are the targets of the ARP request sent to determine the health of the link to the targets. Specify these values in IPv4 dotted decimal format. At least one IP address must be given for ARP monitoring to function. The maximum number of targets that can be specified is 16. The default value is no IP addresses.
	ARPIPTargets string `systemd:",omitempty"`

	// Specifies the quantity of ARPIPTargets that must be reachable in order for the ARP monitor to consider a slave as being up. This option affects only active-backup mode for slaves with ARPValidate enabled. Possible values are "any" and "all".
	ARPAllTargets string `systemd:",omitempty"`

	// Specifies the reselection policy for the primary slave. This affects how the primary slave is chosen to become the active slave when failure of the active slave or recovery of the primary slave occurs. This option is designed to prevent flip-flopping between the primary slave and other slaves. Possible values are "always", "better" and "failure".
	PrimaryReselectPolicy string `systemd:",omitempty"`

	// Specifies the number of IGMP membership reports to be issued after a failover event. One membership report is issued immediately after the failover, subsequent packets are sent in each 200ms interval. The valid range is 0–255. Defaults to 1. A value of 0 prevents the IGMP membership report from being issued in response to the failover event.
	ResendIGMP string `systemd:",omitempty"`

	// Specify the number of packets to transmit through a slave before moving to the next one. When set to 0, then a slave is chosen at random. The valid range is 0–65535. Defaults to 1. This option only has effect when in balance-rr mode.
	PacketsPerSlave string `systemd:",omitempty"`

	// Specify the number of peer notifications (gratuitous ARPs and unsolicited IPv6 Neighbor Advertisements) to be issued after a failover event. As soon as the link is up on the new slave, a peer notification is sent on the bonding device and each VLAN sub-device. This is repeated at each link monitor interval (ARPIntervalSec or MIIMonitorSec, whichever is active) if the number is greater than 1. The valid range is 0–255. The default value is 1. These options affect only the active-backup mode.
	GratuitousARP string `systemd:",omitempty"`

	// Takes a boolean. Specifies that duplicate frames (received on inactive ports) should be dropped when false, or delivered when true. Normally, bonding will drop duplicate frames (received on inactive ports), which is desirable for most users. But there are some times it is nice to allow duplicate frames to be delivered. The default value is false (drop duplicate frames received on inactive ports).
	AllSlavesActive *bool `systemd:",omitempty"`

	// Takes a boolean. Specifies if dynamic shuffling of flows is enabled. Applies only for balance-tlb mode. Defaults to unset.
	DynamicTransmitLoadBalancing *bool `systemd:",omitempty"`

	// Specifies the minimum number of links that must be active before asserting carrier. The default value is 0.
	MinLinks string `systemd:",omitempty"`
}

type XfrmSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Sets the ID/key of the xfrm interface which needs to be associated with a SA/policy. Can be decimal or hexadecimal, valid range is 0-0xffffffff, defaults to 0.
	InterfaceId string `systemd:",omitempty"`

	// Takes a boolean. If set to "no", the xfrm interface should have an underlying device which can be used for hardware offloading. Defaults to "no". See systemd.network(5) for how to configure the underlying device.
	Independent *bool `systemd:",omitempty"`
}

// The [VRF] section only applies for netdevs of kind "vrf"
type VRFSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The numeric routing table identifier. This setting is compulsory.
	Table string `systemd:",omitempty"`
}
