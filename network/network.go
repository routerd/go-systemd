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

import "routerd.net/go-systemd"

// Network configuration
// https://www.freedesktop.org/software/systemd/man/systemd.network.html
type Network struct {
	systemd.SectionList // SectionList to store unknown sections

	Match                           *MatchSection
	Link                            *LinkSection
	SRIOVs                          []SRIOVSection `systemd:"SR-IOV"`
	Network                         *NetworkSection
	Addresses                       []AddressSection           `systemd:"Address"`
	Neighbors                       []NeighborSection          `systemd:"Neighbor"`
	IPv6AddressLabels               []IPv6AddressLabelSection  `systemd:"IPv6AddressLabel"`
	RoutingPolicyRules              []RoutingPolicyRuleSection `systemd:"RoutingPolicyRule"`
	NextHops                        []NextHopSection           `systemd:"NextHop"`
	Routes                          []RouteSection             `systemd:"Route"`
	DHCPv4                          *DHCPv4Section
	DHCPv6                          *DHCPv6Section
	DHCPv6PrefixDelegation          *DHCPv6PrefixDelegationSection
	IPv6AcceptRA                    *IPv6AcceptRASection
	DHCPServer                      *DHCPServerSection
	IPv6PrefixDelegation            *IPv6PrefixDelegationSection
	IPv6Prefixes                    []IPv6PrefixSection      `systemd:"IPv6Prefix"`
	IPv6RoutePrefixes               []IPv6RoutePrefixSection `systemd:"IPv6RoutePrefix"`
	Bridge                          *BridgeSection
	BridgeFDBs                      []BridgeFDBSection `systemd:"BridgeFDB"`
	LLDP                            *LLDPSection
	CAN                             *CANSection
	QDisc                           *QDiscSection
	NetworkEmulator                 *NetworkEmulatorSection
	TokenBucketFilters              []TokenBucketFilterSection          `systemd:"TokenBucketFilter"`
	PIEs                            []PIESection                        `systemd:"PIE"`
	StochasticFairBlue              []StochasticFairBlueSection         `systemd:"StochasticFairBlue"`
	StochasticFairnessQueueing      []StochasticFairnessQueueingSection `systemd:"StochasticFairnessQueueing"`
	BFIFO                           []BFIFOSection
	PFIFO                           []PFIFOSection
	PFIFOHeadDrop                   []PFIFOHeadDropSection
	PFIFOFast                       []PFIFOFastSection
	CAKE                            []CAKESection
	ControlledDelay                 []ControlledDelaySection
	DeficitRoundRobinScheduler      []DeficitRoundRobinSchedulerSection
	DeficitRoundRobinSchedulerClass []DeficitRoundRobinSchedulerClassSection
	EnhancedTransmissionSelection   []EnhancedTransmissionSelectionSection
	GenericRandomEarlyDetection     []GenericRandomEarlyDetectionSection
	FairQueueingControlledDelay     []FairQueueingControlledDelaySection
	FairQueueing                    []FairQueueingSection
	TrivialLinkEqualizer            []TrivialLinkEqualizerSection
	HierarchyTokenBucket            []HierarchyTokenBucketSection
	HierarchyTokenBucketClass       []HierarchyTokenBucketClassSection
	HeavyHitterFilter               []HeavyHitterFilterSection
	QuickFairQueueing               []QuickFairQueueingSection
	QuickFairQueueingClass          []QuickFairQueueingClassSection
	BridgeVLAN                      []BridgeVLANSection
}

// The network file contains a [Match] section, which determines if a given network file may be applied to a given device; and a [Network] section specifying how the device should be configured. The first (in lexical order) of the network files that matches a given device is applied, all later files are ignored, even if they match as well.
//
// A network file is said to match a network interface if all matches specified by the [Match] section are satisfied. When a network file does not contain valid settings in [Match] section, then the file will match all interfaces and systemd-networkd warns about that. Hint: to avoid the warning and to make it clear that all interfaces shall be matched, add the following:
// Name=*
type MatchSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// A whitespace-separated list of hardware addresses. Use full colon-, hyphen- or dot-delimited hexadecimal. See the example below. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	// Example:
	// MACAddress=01:23:45:67:89:ab 00-11-22-33-44-55 AABB.CCDD.EEFF
	PermanentMACAddress string `systemd:",omitempty"`

	// A whitespace-separated list of shell-style globs matching the persistent path, as exposed by the udev property ID_PATH.
	Path string `systemd:",omitempty"`

	// A whitespace-separated list of shell-style globs matching the driver currently bound to the device, as exposed by the udev property ID_NET_DRIVER of its parent device, or if that is not set, the driver as exposed by ethtool -i of the device itself. If the list is prefixed with a "!", the test is inverted.
	Driver string `systemd:",omitempty"`

	// A whitespace-separated list of shell-style globs matching the device type, as exposed by networkctl status. If the list is prefixed with a "!", the test is inverted.
	Type string `systemd:",omitempty"`

	// A whitespace-separated list of udev property name with its value after a equal ("="). If multiple properties are specified, the test results are ANDed. If the list is prefixed with a "!", the test is inverted. If a value contains white spaces, then please quote whole key and value pair. If a value contains quotation, then please escape the quotation with "\".
	// Example: if a .link file has the following:
	// Property=ID_MODEL_ID=9999 "ID_VENDOR_FROM_DATABASE=vendor name" "KEY=with \"quotation\""
	// then, the .link file matches only when an interface has all the above three properties.
	Property string `systemd:",omitempty"`

	// A whitespace-separated list of shell-style globs matching the device name, as exposed by the udev property "INTERFACE", or device's alternative names. If the list is prefixed with a "!", the test is inverted.
	Name string `systemd:",omitempty"`

	// A whitespace-separated list of wireless network type. Supported values are "ad-hoc", "station", "ap", "ap-vlan", "wds", "monitor", "mesh-point", "p2p-client", "p2p-go", "p2p-device", "ocb", and "nan". If the list is prefixed with a "!", the test is inverted.
	WLANInterfaceType string `systemd:",omitempty"`

	// A whitespace-separated list of shell-style globs matching the SSID of the currently connected wireless LAN. If the list is prefixed with a "!", the test is inverted.
	SSID string `systemd:",omitempty"`

	// A whitespace-separated list of hardware address of the currently connected wireless LAN. Use full colon-, hyphen- or dot-delimited hexadecimal. See the example in MACAddress=. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list is reset.
	BSSID string `systemd:",omitempty"`

	// Matches against the hostname or machine ID of the host. See ConditionHost= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Host string `systemd:",omitempty"`

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation. See ConditionVirtualization= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Virtualization string `systemd:",omitempty"`

	// Checks whether a specific kernel command line option is set. See ConditionKernelCommandLine= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelCommandLine string `systemd:",omitempty"`

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See ConditionKernelVersion= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelVersion string `systemd:",omitempty"`

	// Checks whether the system is running on a specific architecture. See ConditionArchitecture= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Architecture string `systemd:",omitempty"`
}

type LinkSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The hardware address to set for the device.
	MACAddress string `systemd:",omitempty"`

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G, are supported and are understood to the base of 1024.
	//
	// Note that if IPv6 is enabled on the interface, and the MTU is chosen below 1280 (the minimum MTU for IPv6) it will automatically be increased to this value.
	MTUBytes string `systemd:",omitempty"`

	// Takes a boolean. If set to true, the ARP (low-level Address Resolution Protocol) for this interface is enabled. When unset, the kernel's default will be used.
	//
	// For example, disabling ARP is useful when creating multiple MACVLAN or VLAN virtual interfaces atop a single lower-level physical interface, which will then only serve as a link/"bridge" device aggregating traffic to the same physical link and not participate in the network otherwise.
	ARP string `systemd:",omitempty"`

	// Takes a boolean. If set to true, the multicast flag on the device is enabled.
	Multicast string `systemd:",omitempty"`

	// Takes a boolean. If set to true, the driver retrieves all multicast packets from the network. This happens when multicast routing is enabled.
	AllMulticast string `systemd:",omitempty"`

	// Takes a boolean. When "yes", no attempts are made to bring up or configure matching links, equivalent to when there are no matching network files. Defaults to "no".
	//
	// This is useful for preventing later matching network files from interfering with certain interfaces that are fully controlled by other applications.
	Unmanaged string `systemd:",omitempty"`

	// Link groups are similar to port ranges found in managed switches. When network interfaces are added to a numbered group, operations on all the interfaces from that group can be performed at once. An unsigned integer in the range 0—4294967294. Defaults to unset.
	Group string `systemd:",omitempty"`

	// Takes a boolean or a minimum operational state and an optional maximum operational state. Please see networkctl(1) for possible operational states. When "yes", the network is deemed required when determining whether the system is online when running systemd-networkd-wait-online. When "no", the network is ignored when checking for online state. When a minimum operational state and an optional maximum operational state are set, "yes" is implied, and this controls the minimum and maximum operational state required for the network interface to be considered online. Defaults to "yes".
	//
	// The network will be brought up normally in all cases, but in the event that there is no address being assigned by DHCP or the cable is not plugged in, the link will simply remain offline and be skipped automatically by systemd-networkd-wait-online if "RequiredForOnline=no".
	RequiredForOnline string `systemd:",omitempty"`
}

// The [SR-IOV] section accepts the following keys. Specify several [SR-IOV] sections to configure several SR-IOVs. SR-IOV provides the ability to partition a single physical PCI resource into virtual PCI functions which can then be injected into a VM. In the case of network VFs, SR-IOV improves north-south network performance (that is, traffic with endpoints outside the host machine) by allowing traffic to bypass the host machine’s network stack.
type SRIOVSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies a Virtual Function (VF), lightweight PCIe function designed solely to move data in and out. Takes an unsigned integer in the range 0..2147483646. This option is compulsory.
	VirtualFunction string `systemd:",omitempty"`

	// Specifies VLAN ID of the virtual function. Takes an unsigned integer in the range 1..4095.
	VLANId string `systemd:",omitempty"`

	// Specifies quality of service of the virtual function. Takes an unsigned integer in the range 1..4294967294.
	QualityOfService string `systemd:",omitempty"`

	// Specifies VLAN protocol of the virtual function. Takes "802.1Q" or "802.1ad".
	VLANProtocol string `systemd:",omitempty"`

	// Takes a boolean. Controls the MAC spoof checking. When unset, the kernel's default will be used.
	MACSpoofCheck string `systemd:",omitempty"`

	// Takes a boolean. Toggle the ability of querying the receive side scaling (RSS) configuration of the virtual function (VF). The VF RSS information like RSS hash key may be considered sensitive on some devices where this information is shared between VF and the physical function (PF). When unset, the kernel's default will be used.
	QueryReceiveSideScaling string `systemd:",omitempty"`

	// Takes a boolean. Allows to set trust mode of the virtual function (VF). When set, VF users can set a specific feature which may impact security and/or performance. When unset, the kernel's default will be used.
	Trust string `systemd:",omitempty"`

	// Allows to set the link state of the virtual function (VF). Takes a boolean or a special value "auto". Setting to "auto" means a reflection of the physical function (PF) link state, "yes" lets the VF to communicate with other VFs on this host even if the PF link state is down, "no" causes the hardware to drop any packets sent by the VF. When unset, the kernel's default will be used.
	LinkState string `systemd:",omitempty"`

	// Specifies the MAC address for the virtual function.
	MACAddress string `systemd:",omitempty"`
}

type NetworkSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// A description of the device. This is only used for presentation purposes.
	Description string `systemd:",omitempty"`

	// Enables DHCPv4 and/or DHCPv6 client support. Accepts "yes", "no", "ipv4", or "ipv6". Defaults to "no".
	//
	// Note that DHCPv6 will by default be triggered by Router Advertisement, if that is enabled, regardless of this parameter. By enabling DHCPv6 support explicitly, the DHCPv6 client will be started regardless of the presence of routers on the link, or what flags the routers pass. See "IPv6AcceptRA=".
	//
	// Furthermore, note that by default the domain name specified through DHCP is not used for name resolution. See option UseDomains= below.
	//
	// See the [DHCPv4] or [DHCPv6] sections below for further configuration options for the DHCP client support.
	DHCP string `systemd:",omitempty"`

	// Takes a boolean. If set to "yes", DHCPv4 server will be started. Defaults to "no". Further settings for the DHCP server may be set in the [DHCPServer] section described below.
	DHCPServer string `systemd:",omitempty"`

	// Enables link-local address autoconfiguration. Accepts "yes", "no", "ipv4", "ipv6", "fallback", or "ipv4-fallback". If "fallback" or "ipv4-fallback" is specified, then an IPv4 link-local address is configured only when DHCPv4 fails. If "fallback", an IPv6 link-local address is always configured, and if "ipv4-fallback", the address is not configured. Note that, the fallback mechanism works only when DHCPv4 client is enabled, that is, it requires "DHCP=yes" or "DHCP=ipv4". If Bridge= is set, defaults to "no", and if not, defaults to "ipv6".
	LinkLocalAddressing string `systemd:",omitempty"`

	// Specifies how IPv6 link local address is generated. Takes one of "eui64", "none", "stable-privacy" and "random". When unset, the kernel's default will be used. Note that if LinkLocalAdressing= not configured as "ipv6" then IPv6LinkLocalAddressGenerationMode= is ignored.
	IPv6LinkLocalAddressGenerationMode string `systemd:",omitempty"`

	// Takes a boolean. If set to true, sets up the route needed for non-IPv4LL hosts to communicate with IPv4LL-only hosts. Defaults to false.
	IPv4LLRoute string `systemd:",omitempty"`

	// Takes a boolean. If set to true, sets up the default route bound to the interface. Defaults to false. This is useful when creating routes on point-to-point interfaces. This is equivalent to e.g. the following.
	//
	// ip route add default dev veth99
	DefaultRouteOnDevice string `systemd:",omitempty"`
	// Specifies an optional address generation mode and a required IPv6 address. If the mode is present, the two parts must be separated with a colon "mode:address". The address generation mode may be either prefixstable or static. If not specified, static is assumed.
	//
	// When the mode is set to static, or unspecified, the lower bits of the supplied address are combined with the upper bits of a prefix received in a Router Advertisement message to form a complete address. Note that if multiple prefixes are received in an RA message, or in multiple RA messages, addresses will be formed from each of them using the supplied address. This mode implements SLAAC but uses a static interface identifier instead of an identifier generated using the EUI-64 algorithm. Because the interface identifier is static, if Duplicate Address Detection detects that the computed address is a duplicate (in use by another node on the link), then this mode will fail to provide an address for that prefix.
	//
	// When the mode is set to "prefixstable" the RFC 7217 algorithm for generating interface identifiers will be used, but only when a prefix received in an RA message matches the supplied address. See RFC 7217. Prefix matching will be attempted against each prefixstable IPv6Token variable provided in the configuration; if a received prefix does not match any of the provided addresses, then the EUI-64 algorithm will be used to form an interface identifier for that prefix. This mode is also SLAAC, but with a potentially stable interface identifier which does not directly map to the interface's hardware address. Note that the prefixstable algorithm includes both the interface's name and MAC address in the hash used to compute the interface identifier, so if either of those are changed the resulting interface identifier (and address) will change, even if the prefix received in the RA message has not changed. Note that if multiple prefixstable IPv6Token variables are supplied with addresses that match a prefix received in an RA message, only the first one will be used to generate addresses.
	IPv6Token string `systemd:",omitempty"`

	// Takes a boolean or "resolve". When true, enables Link-Local Multicast Name Resolution on the link. When set to "resolve", only resolution is enabled, but not host registration and announcement. Defaults to true. This setting is read by systemd-resolved.service(8).
	LLMNR string `systemd:",omitempty"`

	// Takes a boolean or "resolve". When true, enables Multicast DNS support on the link. When set to "resolve", only resolution is enabled, but not host or service registration and announcement. Defaults to false. This setting is read by systemd-resolved.service(8).
	MulticastDNS string `systemd:",omitempty"`

	// Takes a boolean or "opportunistic". When true, enables DNS-over-TLS support on the link. When set to "opportunistic", compatibility with non-DNS-over-TLS servers is increased, by automatically turning off DNS-over-TLS servers in this case. This option defines a per-interface setting for resolved.conf(5)'s global DNSOverTLS= option. Defaults to false. This setting is read by systemd-resolved.service(8).
	DNSOverTLS string `systemd:",omitempty"`

	// Takes a boolean or "allow-downgrade". When true, enables DNSSEC DNS validation support on the link. When set to "allow-downgrade", compatibility with non-DNSSEC capable networks is increased, by automatically turning off DNSSEC in this case. This option defines a per-interface setting for resolved.conf(5)'s global DNSSEC= option. Defaults to false. This setting is read by systemd-resolved.service(8).
	DNSSEC string `systemd:",omitempty"`

	// A space-separated list of DNSSEC negative trust anchor domains. If specified and DNSSEC is enabled, look-ups done via the interface's DNS server will be subject to the list of negative trust anchors, and not require authentication for the specified domains, or anything below it. Use this to disable DNSSEC authentication for specific private domains, that cannot be proven valid using the Internet DNS hierarchy. Defaults to the empty list. This setting is read by systemd-resolved.service(8).
	DNSSECNegativeTrustAnchors string `systemd:",omitempty"`

	// Controls support for Ethernet LLDP packet reception. LLDP is a link-layer protocol commonly implemented on professional routers and bridges which announces which physical port a system is connected to, as well as other related data. Accepts a boolean or the special value "routers-only". When true, incoming LLDP packets are accepted and a database of all LLDP neighbors maintained. If "routers-only" is set only LLDP data of various types of routers is collected and LLDP data about other types of devices ignored (such as stations, telephones and others). If false, LLDP reception is disabled. Defaults to "routers-only". Use networkctl(1) to query the collected neighbor data. LLDP is only available on Ethernet links. See EmitLLDP= below for enabling LLDP packet emission from the local system.
	LLDP string `systemd:",omitempty"`

	// Controls support for Ethernet LLDP packet emission. Accepts a boolean parameter or the special values "nearest-bridge", "non-tpmr-bridge" and "customer-bridge". Defaults to false, which turns off LLDP packet emission. If not false, a short LLDP packet with information about the local system is sent out in regular intervals on the link. The LLDP packet will contain information about the local hostname, the local machine ID (as stored in machine-id(5)) and the local interface name, as well as the pretty hostname of the system (as set in machine-info(5)). LLDP emission is only available on Ethernet links. Note that this setting passes data suitable for identification of host to the network and should thus not be enabled on untrusted networks, where such identification data should not be made available. Use this option to permit other systems to identify on which interfaces they are connected to this system. The three special values control propagation of the LLDP packets. The "nearest-bridge" setting permits propagation only to the nearest connected bridge, "non-tpmr-bridge" permits propagation across Two-Port MAC Relays, but not any other bridges, and "customer-bridge" permits propagation until a customer bridge is reached. For details about these concepts, see IEEE 802.1AB-2016. Note that configuring this setting to true is equivalent to "nearest-bridge", the recommended and most restricted level of propagation. See LLDP= above for an option to enable LLDP reception.
	EmitLLDP string `systemd:",omitempty"`

	// A link name or a list of link names. When set, controls the behavior of the current link. When all links in the list are in an operational down state, the current link is brought down. When at least one link has carrier, the current interface is brought up.
	BindCarrier string `systemd:",omitempty"`

	// A static IPv4 or IPv6 address and its prefix length, separated by a "/" character. Specify this key more than once to configure several addresses. The format of the address must be as described in inet_pton(3). This is a short-hand for an [Address] section only containing an Address key (see below). This option may be specified more than once.
	//
	// If the specified address is "0.0.0.0" (for IPv4) or "::" (for IPv6), a new address range of the requested size is automatically allocated from a system-wide pool of unused ranges. Note that the prefix length must be equal or larger than 8 for IPv4, and 64 for IPv6. The allocated range is checked against all current network interfaces and all known network configuration files to avoid address range conflicts. The default system-wide pool consists of 192.168.0.0/16, 172.16.0.0/12 and 10.0.0.0/8 for IPv4, and fd00::/8 for IPv6. This functionality is useful to manage a large number of dynamically created network interfaces with the same network configuration and automatic address range assignment.
	Address string `systemd:",omitempty"`

	// The gateway address, which must be in the format described in inet_pton(3). This is a short-hand for a [Route] section only containing a Gateway key. This option may be specified more than once.
	Gateway string `systemd:",omitempty"`

	// A DNS server address, which must be in the format described in inet_pton(3). This option may be specified more than once. Each address can optionally take a port number separated with ":", a network interface name or index separated with "%", and a Server Name Indication (SNI) separated with "#". When IPv6 address is specified with a port number, then the address must be in the square brackets. That is, the acceptable full formats are "111.222.333.444:9953%ifname#example.com" for IPv4 and "[1111:2222::3333]:9953%ifname#example.com" for IPv6. This setting can be specified multiple times. If an empty string is assigned, then the all previous assignments are cleared. This setting is read by systemd-resolved.service(8).
	DNS string `systemd:",omitempty"`

	// A whitespace-separated list of domains which should be resolved using the DNS servers on this link. Each item in the list should be a domain name, optionally prefixed with a tilde ("~"). The domains with the prefix are called "routing-only domains". The domains without the prefix are called "search domains" and are first used as search suffixes for extending single-label hostnames (hostnames containing no dots) to become fully qualified domain names (FQDNs). If a single-label hostname is resolved on this interface, each of the specified search domains are appended to it in turn, converting it into a fully qualified domain name, until one of them may be successfully resolved.
	//
	// Both "search" and "routing-only" domains are used for routing of DNS queries: look-ups for hostnames ending in those domains (hence also single label names, if any "search domains" are listed), are routed to the DNS servers configured for this interface. The domain routing logic is particularly useful on multi-homed hosts with DNS servers serving particular private DNS zones on each interface.
	//
	// The "routing-only" domain "~." (the tilde indicating definition of a routing domain, the dot referring to the DNS root domain which is the implied suffix of all valid DNS names) has special effect. It causes all DNS traffic which does not match another configured domain routing entry to be routed to DNS servers specified for this interface. This setting is useful to prefer a certain set of DNS servers if a link on which they are connected is available.
	//
	// This setting is read by systemd-resolved.service(8). "Search domains" correspond to the domain and search entries in resolv.conf(5). Domain name routing has no equivalent in the traditional glibc API, which has no concept of domain name servers limited to a specific link.
	Domains string `systemd:",omitempty"`

	// Takes a boolean argument. If true, this link's configured DNS servers are used for resolving domain names that do not match any link's configured Domains= setting. If false, this link's configured DNS servers are never used for such domains, and are exclusively used for resolving names that match at least one of the domains configured on this link. If not specified defaults to an automatic mode: queries not matching any link's configured domains will be routed to this link if it has no routing-only domains configured.
	DNSDefaultRoute string `systemd:",omitempty"`

	// An NTP server address (either an IP address, or a hostname). This option may be specified more than once. This setting is read by systemd-timesyncd.service(8).
	NTP string `systemd:",omitempty"`

	// Configures IP packet forwarding for the system. If enabled, incoming packets on any network interface will be forwarded to any other interfaces according to the routing table. Takes a boolean, or the values "ipv4" or "ipv6", which only enable IP packet forwarding for the specified address family. This controls the net.ipv4.ip_forward and net.ipv6.conf.all.forwarding sysctl options of the network interface (see ip-sysctl.txt for details about sysctl options). Defaults to "no".
	//
	// Note: this setting controls a global kernel option, and does so one way only: if a network that has this setting enabled is set up the global setting is turned on. However, it is never turned off again, even after all networks with this setting enabled are shut down again.
	//
	// To allow IP packet forwarding only between specific network interfaces use a firewall.
	IPForward string `systemd:",omitempty"`

	// Configures IP masquerading for the network interface. If enabled, packets forwarded from the network interface will be appear as coming from the local host. Takes a boolean argument. Implies IPForward=ipv4. Defaults to "no".
	IPMasquerade string `systemd:",omitempty"`

	// Configures use of stateless temporary addresses that change over time (see RFC 4941, Privacy Extensions for Stateless Address Autoconfiguration in IPv6). Takes a boolean or the special values "prefer-public" and "kernel". When true, enables the privacy extensions and prefers temporary addresses over public addresses. When "prefer-public", enables the privacy extensions, but prefers public addresses over temporary addresses. When false, the privacy extensions remain disabled. When "kernel", the kernel's default setting will be left in place. Defaults to "no".
	IPv6PrivacyExtensions string `systemd:",omitempty"`

	// Takes a boolean. Controls IPv6 Router Advertisement (RA) reception support for the interface. If true, RAs are accepted; if false, RAs are ignored. When RAs are accepted, they may trigger the start of the DHCPv6 client if the relevant flags are set in the RA data, or if no routers are found on the link. The default is to disable RA reception for bridge devices or when IP forwarding is enabled, and to enable it otherwise. Cannot be enabled on bond devices and when link local addressing is disabled.
	//
	// Further settings for the IPv6 RA support may be configured in the [IPv6AcceptRA] section, see below.
	//
	// Also see ip-sysctl.txt in the kernel documentation regarding "accept_ra", but note that systemd's setting of 1 (i.e. true) corresponds to kernel's setting of 2.
	//
	// Note that kernel's implementation of the IPv6 RA protocol is always disabled, regardless of this setting. If this option is enabled, a userspace implementation of the IPv6 RA protocol is used, and the kernel's own implementation remains disabled, since systemd-networkd needs to know all details supplied in the advertisements, and these are not available from the kernel if the kernel's own implementation is used.
	IPv6AcceptRA string `systemd:",omitempty"`

	// Configures the amount of IPv6 Duplicate Address Detection (DAD) probes to send. When unset, the kernel's default will be used.
	IPv6DuplicateAddressDetection string `systemd:",omitempty"`

	// Configures IPv6 Hop Limit. For each router that forwards the packet, the hop limit is decremented by 1. When the hop limit field reaches zero, the packet is discarded. When unset, the kernel's default will be used.
	IPv6HopLimit string `systemd:",omitempty"`

	// Takes a boolean. Accept packets with local source addresses. In combination with suitable routing, this can be used to direct packets between two local interfaces over the wire and have them accepted properly. When unset, the kernel's default will be used.
	IPv4AcceptLocal string `systemd:",omitempty"`

	// Takes a boolean. Configures proxy ARP for IPv4. Proxy ARP is the technique in which one host, usually a router, answers ARP requests intended for another machine. By "faking" its identity, the router accepts responsibility for routing packets to the "real" destination. See RFC 1027. When unset, the kernel's default will be used.
	IPv4ProxyARP string `systemd:",omitempty"`

	// Takes a boolean. Configures proxy NDP for IPv6. Proxy NDP (Neighbor Discovery Protocol) is a technique for IPv6 to allow routing of addresses to a different destination when peers expect them to be present on a certain physical link. In this case a router answers Neighbour Advertisement messages intended for another machine by offering its own MAC address as destination. Unlike proxy ARP for IPv4, it is not enabled globally, but will only send Neighbour Advertisement messages for addresses in the IPv6 neighbor proxy table, which can also be shown by ip -6 neighbour show proxy. systemd-networkd will control the per-interface `proxy_ndp` switch for each configured interface depending on this option. When unset, the kernel's default will be used.
	IPv6ProxyNDP string `systemd:",omitempty"`

	// An IPv6 address, for which Neighbour Advertisement messages will be proxied. This option may be specified more than once. systemd-networkd will add the IPv6ProxyNDPAddress= entries to the kernel's IPv6 neighbor proxy table. This option implies IPv6ProxyNDP=yes but has no effect if IPv6ProxyNDP has been set to false. When unset, the kernel's default will be used.
	IPv6ProxyNDPAddress string `systemd:",omitempty"`

	// Whether to enable or disable Router Advertisement sending on a link. Allowed values are "static" which distributes prefixes as defined in the [IPv6PrefixDelegation] and any [IPv6Prefix] sections, "dhcpv6" which requests prefixes using a DHCPv6 client configured for another link and any values configured in the [IPv6PrefixDelegation] section while ignoring all static prefix configuration sections, "yes" which uses both static configuration and DHCPv6, and "false" which turns off IPv6 prefix delegation altogether. Defaults to "false". See the [IPv6PrefixDelegation] and the [IPv6Prefix] sections for more configuration options.
	IPv6PrefixDelegation string `systemd:",omitempty"`

	// Configures IPv6 maximum transmission unit (MTU). An integer greater than or equal to 1280 bytes. When unset, the kernel's default will be used.
	IPv6MTUBytes string `systemd:",omitempty"`

	// The name of the bridge to add the link to. See systemd.netdev(5).
	Bridge string `systemd:",omitempty"`

	// The name of the bond to add the link to. See systemd.netdev(5).
	Bond string `systemd:",omitempty"`

	// The name of the VRF to add the link to. See systemd.netdev(5).
	VRF string `systemd:",omitempty"`

	// The name of a VLAN to create on the link. See systemd.netdev(5). This option may be specified more than once.
	VLAN string `systemd:",omitempty"`

	// The name of a IPVLAN to create on the link. See systemd.netdev(5). This option may be specified more than once.
	IPVLAN string `systemd:",omitempty"`

	// The name of a MACVLAN to create on the link. See systemd.netdev(5). This option may be specified more than once.
	MACVLAN string `systemd:",omitempty"`

	// The name of a VXLAN to create on the link. See systemd.netdev(5). This option may be specified more than once.
	VXLAN string `systemd:",omitempty"`

	// The name of a Tunnel to create on the link. See systemd.netdev(5). This option may be specified more than once.
	Tunnel string `systemd:",omitempty"`

	// The name of a MACsec device to create on the link. See systemd.netdev(5). This option may be specified more than once.
	MACsec string `systemd:",omitempty"`

	ActiveSlave string `systemd:",omitempty"`
	// Takes a boolean. Specifies the new active slave. The "ActiveSlave=" option is only valid for following modes: "active-backup", "balance-alb" and "balance-tlb". Defaults to false.

	// Takes a boolean. Specifies which slave is the primary device. The specified device will always be the active slave while it is available. Only when the primary is off-line will alternate devices be used. This is useful when one slave is preferred over another, e.g. when one slave has higher throughput than another. The "PrimarySlave=" option is only valid for following modes: "active-backup", "balance-alb" and "balance-tlb". Defaults to false.
	PrimarySlave string `systemd:",omitempty"`

	// Takes a boolean. Allows networkd to configure a specific link even if it has no carrier. Defaults to false. If IgnoreCarrierLoss= is not explicitly set, it will default to this value.
	ConfigureWithoutCarrier string `systemd:",omitempty"`

	// Takes a boolean. Allows networkd to retain both the static and dynamic configuration of the interface even if its carrier is lost. When unset, the value specified with ConfigureWithoutCarrier= is used.
	IgnoreCarrierLoss string `systemd:",omitempty"`

	// The name of the xfrm to create on the link. See systemd.netdev(5). This option may be specified more than once.
	Xfrm string `systemd:",omitempty"`

	// Takes a boolean or one of "static", "dhcp-on-stop", "dhcp". When "static", systemd-networkd will not drop static addresses and routes on starting up process. When set to "dhcp-on-stop", systemd-networkd will not drop addresses and routes on stopping the daemon. When "dhcp", the addresses and routes provided by a DHCP server will never be dropped even if the DHCP lease expires. This is contrary to the DHCP specification, but may be the best choice if, e.g., the root filesystem relies on this connection. The setting "dhcp" implies "dhcp-on-stop", and "yes" implies "dhcp" and "static". Defaults to "no".
	KeepConfiguration string `systemd:",omitempty"`
}

// An [Address] section accepts the following keys. Specify several [Address] sections to configure several addresses.
type AddressSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// As in the [Network] section. This key is mandatory. Each [Address] section can contain one Address= setting.
	Address string `systemd:",omitempty"`

	// The peer address in a point-to-point connection. Accepts the same format as the Address= key.
	Peer string `systemd:",omitempty"`

	// The broadcast address, which must be in the format described in inet_pton(3). This key only applies to IPv4 addresses. If it is not given, it is derived from the Address= key.
	Broadcast string `systemd:",omitempty"`

	// An address label.
	Label string `systemd:",omitempty"`

	// Allows the default "preferred lifetime" of the address to be overridden. Only three settings are accepted: "forever" or "infinity" which is the default and means that the address never expires, and "0" which means that the address is considered immediately "expired" and will not be used, unless explicitly requested. A setting of PreferredLifetime=0 is useful for addresses which are added to be used only by a specific application, which is then configured to use them explicitly.
	PreferredLifetime string `systemd:",omitempty"`

	// The scope of the address, which can be "global", "link" or "host" or an unsigned integer in the range 0—255. Defaults to "global".
	Scope string `systemd:",omitempty"`

	// Takes a boolean. Designates this address the "home address" as defined in RFC 6275. Supported only on IPv6. Defaults to false.
	HomeAddress string `systemd:",omitempty"`

	// Takes one of "ipv4", "ipv6", "both", "none". When "ipv4", performs IPv4 Duplicate Address Detection. See RFC 5224. When "ipv6", performs IPv6 Duplicate Address Detection. See RFC 4862. Defaults to "ipv6".
	DuplicateAddressDetection string `systemd:",omitempty"`

	// Takes a boolean. If true the kernel manage temporary addresses created from this one as template on behalf of Privacy Extensions RFC 3041. For this to become active, the use_tempaddr sysctl setting has to be set to a value greater than zero. The given address needs to have a prefix length of 64. This flag allows using privacy extensions in a manually configured network, just like if stateless auto-configuration was active. Defaults to false.
	ManageTemporaryAddress string `systemd:",omitempty"`

	// Takes a boolean. When true, the prefix route for the address is automatically added. Defaults to true.
	AddPrefixRoute string `systemd:",omitempty"`

	// Takes a boolean. Joining multicast group on ethernet level via ip maddr command would not work if we have an Ethernet switch that does IGMP snooping since the switch would not replicate multicast packets on ports that did not have IGMP reports for the multicast addresses. Linux vxlan interfaces created via ip link add vxlan or networkd's netdev kind vxlan have the group option that enables then to do the required join. By extending ip address command with option "autojoin" we can get similar functionality for openvswitch (OVS) vxlan interfaces as well as other tunneling mechanisms that need to receive multicast traffic. Defaults to "no".
	AutoJoin string `systemd:",omitempty"`
}

// A [Neighbor] section accepts the following keys. The neighbor section adds a permanent, static entry to the neighbor table (IPv6) or ARP table (IPv4) for the given hardware address on the links matched for the network. Specify several [Neighbor] sections to configure several static neighbors.
type NeighborSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The IP address of the neighbor.
	Address string `systemd:",omitempty"`

	// The link layer address (MAC address or IP address) of the neighbor.
	LinkLayerAddress string `systemd:",omitempty"`
}

// An [IPv6AddressLabel] section accepts the following keys. Specify several [IPv6AddressLabel] sections to configure several address labels. IPv6 address labels are used for address selection. See RFC 3484. Precedence is managed by userspace, and only the label itself is stored in the kernel
type IPv6AddressLabelSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The label for the prefix, an unsigned integer in the range 0–4294967294. 0xffffffff is reserved. This setting is mandatory.
	Label string `systemd:",omitempty"`

	// IPv6 prefix is an address with a prefix length, separated by a slash "/" character. This key is mandatory.
	Prefix string `systemd:",omitempty"`
}

// An [RoutingPolicyRule] section accepts the following keys. Specify several [RoutingPolicyRule] sections to configure several rules.
type RoutingPolicyRuleSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a number between 0 and 255 that specifies the type of service to match.
	TypeOfService string `systemd:",omitempty"`

	// Specifies the source address prefix to match. Possibly followed by a slash and the prefix length.
	From string `systemd:",omitempty"`

	// Specifies the destination address prefix to match. Possibly followed by a slash and the prefix length.
	To string `systemd:",omitempty"`

	// Specifies the iptables firewall mark value to match (a number between 1 and 4294967295).
	FirewallMark string `systemd:",omitempty"`

	// Specifies the routing table identifier to lookup if the rule selector matches. Takes one of "default", "main", and "local", or a number between 1 and 4294967295. Defaults to "main".
	Table string `systemd:",omitempty"`

	// Specifies the priority of this rule. Priority= is an unsigned integer. Higher number means lower priority, and rules get processed in order of increasing number.
	Priority string `systemd:",omitempty"`

	// Specifies incoming device to match. If the interface is loopback, the rule only matches packets originating from this host.
	IncomingInterface string `systemd:",omitempty"`

	// Specifies the outgoing device to match. The outgoing interface is only available for packets originating from local sockets that are bound to a device.
	OutgoingInterface string `systemd:",omitempty"`

	// Specifies the source IP port or IP port range match in forwarding information base (FIB) rules. A port range is specified by the lower and upper port separated by a dash. Defaults to unset.
	SourcePort string `systemd:",omitempty"`

	// Specifies the destination IP port or IP port range match in forwarding information base (FIB) rules. A port range is specified by the lower and upper port separated by a dash. Defaults to unset.
	DestinationPort string `systemd:",omitempty"`

	// Specifies the IP protocol to match in forwarding information base (FIB) rules. Takes IP protocol name such as "tcp", "udp" or "sctp", or IP protocol number such as "6" for "tcp" or "17" for "udp". Defaults to unset.
	IPProtocol string `systemd:",omitempty"`

	// A boolean. Specifies whether the rule is to be inverted. Defaults to false.
	InvertRule string `systemd:",omitempty"`

	// Takes a special value "ipv4", "ipv6", or "both". By default, the address family is determined by the address specified in To= or From=. If neither To= nor From= are specified, then defaults to "ipv4".
	Family string `systemd:",omitempty"`

	// Takes a username, a user ID, or a range of user IDs separated by a dash. Defaults to unset.
	User string `systemd:",omitempty"`

	// Takes a number N in the range 0-128 and rejects routing decisions that have a prefix length of N or less. Defaults to unset.
	SuppressPrefixLength string `systemd:",omitempty"`
}

// The [NextHop] section is used to manipulate entries in the kernel's "nexthop" tables. The [NextHop] section accepts the following keys. Specify several [NextHop] sections to configure several hops.
type NextHopSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// As in the [Network] section. This is mandatory.
	Gateway string `systemd:",omitempty"`

	// The id of the nexthop (an unsigned integer). If unspecified or '0' then automatically chosen by kernel.
	Id string `systemd:",omitempty"`
}

// The [Route] section accepts the following keys. Specify several [Route] sections to configure several routes.
type RouteSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes the gateway address or special value "_dhcp". If "_dhcp", then the gateway address provided by DHCP (or in the IPv6 case, provided by IPv6 RA) is used.
	Gateway string `systemd:",omitempty"`

	// Takes a boolean. If set to true, the kernel does not have to check if the gateway is reachable directly by the current machine (i.e., the kernel does not need to check if the gateway is attached to the local network), so that we can insert the route in the kernel table without it being complained about. Defaults to "no".
	GatewayOnLink string `systemd:",omitempty"`

	// The destination prefix of the route. Possibly followed by a slash and the prefix length. If omitted, a full-length host route is assumed.
	Destination string `systemd:",omitempty"`

	// The source prefix of the route. Possibly followed by a slash and the prefix length. If omitted, a full-length host route is assumed.
	Source string `systemd:",omitempty"`

	// The metric of the route (an unsigned integer).
	Metric string `systemd:",omitempty"`

	// Specifies the route preference as defined in RFC 4191 for Router Discovery messages. Which can be one of "low" the route has a lowest priority, "medium" the route has a default priority or "high" the route has a highest priority.
	IPv6Preference string `systemd:",omitempty"`

	// The scope of the route, which can be "global", "site", "link", "host", or "nowhere". For IPv4 route, defaults to "host" if Type= is "local" or "nat", and "link" if Type= is "broadcast", "multicast", or "anycast". In other cases, defaults to "global".
	Scope string `systemd:",omitempty"`

	// The preferred source address of the route. The address must be in the format described in inet_pton(3).
	PreferredSource string `systemd:",omitempty"`

	// The table identifier for the route. Takes "default", "main", "local" or a number between 1 and 4294967295. The table can be retrieved using ip route show table num. If unset and Type= is "local", "broadcast", "anycast", or "nat", then "local" is used. In other cases, defaults to "main".
	Table string `systemd:",omitempty"`

	// The protocol identifier for the route. Takes a number between 0 and 255 or the special values "kernel", "boot", "static", "ra" and "dhcp". Defaults to "static".
	Protocol string `systemd:",omitempty"`

	// Specifies the type for the route. Takes one of "unicast", "local", "broadcast", "anycast", "multicast", "blackhole", "unreachable", "prohibit", "throw", "nat", and "xresolve". If "unicast", a regular route is defined, i.e. a route indicating the path to take to a destination network address. If "blackhole", packets to the defined route are discarded silently. If "unreachable", packets to the defined route are discarded and the ICMP message "Host Unreachable" is generated. If "prohibit", packets to the defined route are discarded and the ICMP message "Communication Administratively Prohibited" is generated. If "throw", route lookup in the current routing table will fail and the route selection process will return to Routing Policy Database (RPDB). Defaults to "unicast".
	Type string `systemd:",omitempty"`

	// The TCP initial congestion window is used during the start of a TCP connection. During the start of a TCP session, when a client requests a resource, the server's initial congestion window determines how many data bytes will be sent during the initial burst of data. Takes a size in bytes between 1 and 4294967295 (2^32 - 1). The usual suffixes K, M, G are supported and are understood to the base of 1024. When unset, the kernel's default will be used.
	InitialCongestionWindow string `systemd:",omitempty"`

	// The TCP initial advertised receive window is the amount of receive data (in bytes) that can initially be buffered at one time on a connection. The sending host can send only that amount of data before waiting for an acknowledgment and window update from the receiving host. Takes a size in bytes between 1 and 4294967295 (2^32 - 1). The usual suffixes K, M, G are supported and are understood to the base of 1024. When unset, the kernel's default will be used.
	InitialAdvertisedReceiveWindow string `systemd:",omitempty"`

	// Takes a boolean. When true enables TCP quick ack mode for the route. When unset, the kernel's default will be used.
	QuickAck string `systemd:",omitempty"`

	// Takes a boolean. When true enables TCP fastopen without a cookie on a per-route basis. When unset, the kernel's default will be used.
	FastOpenNoCookie string `systemd:",omitempty"`

	// Takes a boolean. When true enables TTL propagation at Label Switched Path (LSP) egress. When unset, the kernel's default will be used.
	TTLPropagate string `systemd:",omitempty"`

	// The maximum transmission unit in bytes to set for the route. The usual suffixes K, M, G, are supported and are understood to the base of 1024.
	//
	// Note that if IPv6 is enabled on the interface, and the MTU is chosen below 1280 (the minimum MTU for IPv6) it will automatically be increased to this value.
	MTUBytes string `systemd:",omitempty"`

	// Takes string; "CS6" or "CS4". Used to set IP service type to CS6 (network control) or CS4 (Realtime). Defaults to CS6.
	IPServiceType string `systemd:",omitempty"`

	// address[@name] [weight]
	// Configures multipath route. Multipath routing is the technique of using multiple alternative paths through a network. Takes gateway address. Optionally, takes a network interface name or index separated with "@", and a weight in 1..256 for this multipath route separated with whitespace. This setting can be specified multiple times. If an empty string is assigned, then the all previous assignments are cleared.
	MultiPathRoute string `systemd:",omitempty"`
}

// The [DHCPv4] section configures the DHCPv4 client, if it is enabled with the DHCP= setting described above:
type DHCPv4Section struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// When true (the default), the DNS servers received from the DHCP server will be used and take precedence over any statically configured ones.
	UseDNS string `systemd:",omitempty"`
	//
	// This corresponds to the nameserver option in resolv.conf(5).

	// When true, the routes to the DNS servers received from the DHCP server will be configured. When UseDNS= is disabled, this setting is ignored. Defaults to false.
	RoutesToDNS string `systemd:",omitempty"`

	// When true (the default), the NTP servers received from the DHCP server will be used by systemd-timesyncd.service and take precedence over any statically configured ones.
	UseNTP string `systemd:",omitempty"`

	// When true (the default), the SIP servers received from the DHCP server will be collected and made available to client programs.
	UseSIP string `systemd:",omitempty"`

	// When true, the interface maximum transmission unit from the DHCP server will be used on the current link. If MTUBytes= is set, then this setting is ignored. Defaults to false.
	UseMTU string `systemd:",omitempty"`

	// Takes a boolean. When true, the options sent to the DHCP server will follow the RFC 7844 (Anonymity Profiles for DHCP Clients) to minimize disclosure of identifying information. Defaults to false.
	Anonymize string `systemd:",omitempty"`
	//
	// This option should only be set to true when MACAddressPolicy= is set to "random" (see systemd.link(5)).
	//
	// Note that this configuration will overwrite others. In concrete, the following variables will be ignored: SendHostname=, ClientIdentifier=, UseRoutes=, UseMTU=, VendorClassIdentifier=, UseTimezone=.
	//
	// With this option enabled DHCP requests will mimic those generated by Microsoft Windows, in order to reduce the ability to fingerprint and recognize installations. This means DHCP request sizes will grow and lease data will be more comprehensive than normally, though most of the requested data is not actually used.

	// When true (the default), the machine's hostname will be sent to the DHCP server. Note that the machine's hostname must consist only of 7-bit ASCII lower-case characters and no spaces or dots, and be formatted as a valid DNS domain name. Otherwise, the hostname is not sent even if this is set to true.
	SendHostname string `systemd:",omitempty"`

	// When configured, the Manufacturer Usage Descriptions (MUD) URL will be sent to the DHCPv4 server. Takes an URL of length up to 255 characters. A superficial verification that the string is a valid URL will be performed. DHCPv4 clients are intended to have at most one MUD URL associated with them. See RFC 8520.
	MUDURL string `systemd:",omitempty"`

	// When true (the default), the hostname received from the DHCP server will be set as the transient hostname of the system.
	UseHostname string `systemd:",omitempty"`

	// Use this value for the hostname which is sent to the DHCP server, instead of machine's hostname. Note that the specified hostname must consist only of 7-bit ASCII lower-case characters and no spaces or dots, and be formatted as a valid DNS domain name.
	Hostname string `systemd:",omitempty"`

	// Takes a boolean, or the special value "route". When true, the domain name received from the DHCP server will be used as DNS search domain over this link, similar to the effect of the Domains= setting. If set to "route", the domain name received from the DHCP server will be used for routing DNS queries only, but not for searching, similar to the effect of the Domains= setting when the argument is prefixed with "~". Defaults to false.
	UseDomains string `systemd:",omitempty"`
	//
	// It is recommended to enable this option only on trusted networks, as setting this affects resolution of all hostnames, in particular of single-label names. It is generally safer to use the supplied domain only as routing domain, rather than as search domain, in order to not have it affect local resolution of single-label names.
	//
	// When set to true, this setting corresponds to the domain option in resolv.conf(5).

	// When true (the default), the static routes will be requested from the DHCP server and added to the routing table with a metric of 1024, and a scope of "global", "link" or "host", depending on the route's destination and gateway. If the destination is on the local host, e.g., 127.x.x.x, or the same as the link's own address, the scope will be set to "host". Otherwise if the gateway is null (a direct route), a "link" scope will be used. For anything else, scope defaults to "global".
	UseRoutes string `systemd:",omitempty"`

	// When true, the gateway will be requested from the DHCP server and added to the routing table with a metric of 1024, and a scope of "link". When unset, the value specified with UseRoutes= is used.
	UseGateway string `systemd:",omitempty"`

	// When true, the timezone received from the DHCP server will be set as timezone of the local system. Defaults to "no".
	UseTimezone string `systemd:",omitempty"`

	// The DHCPv4 client identifier to use. Takes one of "mac", "duid" or "duid-only". If set to "mac", the MAC address of the link is used. If set to "duid", an RFC4361-compliant Client ID, which is the combination of IAID and DUID (see below), is used. If set to "duid-only", only DUID is used, this may not be RFC compliant, but some setups may require to use this. Defaults to "duid".
	ClientIdentifier string `systemd:",omitempty"`

	// The vendor class identifier used to identify vendor type and configuration.
	VendorClassIdentifier string `systemd:",omitempty"`

	// A DHCPv4 client can use UserClass option to identify the type or category of user or applications it represents. The information contained in this option is a string that represents the user class of which the client is a member. Each class sets an identifying string of information to be used by the DHCP service to classify clients. Takes a whitespace-separated list of strings.
	UserClass string `systemd:",omitempty"`

	// Specifies how many times the DHCPv4 client configuration should be attempted. Takes a number or "infinity". Defaults to "infinity". Note that the time between retries is increased exponentially, so the network will not be overloaded even if this number is high.
	MaxAttempts string `systemd:",omitempty"`

	// Override the global DUIDType setting for this network. See networkd.conf(5) for a description of possible values.
	DUIDType string `systemd:",omitempty"`

	// Override the global DUIDRawData setting for this network. See networkd.conf(5) for a description of possible values.
	DUIDRawData string `systemd:",omitempty"`

	// The DHCP Identity Association Identifier (IAID) for the interface, a 32-bit unsigned integer.
	IAID string `systemd:",omitempty"`

	// Request the server to use broadcast messages before the IP address has been configured. This is necessary for devices that cannot receive RAW packets, or that cannot receive packets at all before an IP address has been configured. On the other hand, this must not be enabled on networks where broadcasts are filtered out.
	RequestBroadcast string `systemd:",omitempty"`

	// Set the routing metric for routes specified by the DHCP server. Defaults to 1024.
	RouteMetric string `systemd:",omitempty"`
	//
	// RouteTable=num
	// The table identifier for DHCP routes (a number between 1 and 4294967295, or 0 to unset). The table can be retrieved using ip route show table num.
	//
	// When used in combination with VRF=, the VRF's routing table is used when this parameter is not specified.

	// Specifies the MTU for the DHCP routes. Please see the [Route] section for further details.
	RouteMTUBytes string `systemd:",omitempty"`

	// Allow setting custom port for the DHCP client to listen on.
	ListenPort string `systemd:",omitempty"`

	// Allows to set DHCPv4 lease lifetime when DHCPv4 server does not send the lease lifetime. Takes one of "forever" or "infinity" means that the address never expires. Defaults to unset.
	FallbackLeaseLifetimeSec string `systemd:",omitempty"`

	// When true, the DHCPv4 client sends a DHCP release packet when it stops. Defaults to true.
	SendRelease string `systemd:",omitempty"`

	// A boolean. When "true", the DHCPv4 client receives the IP address from the DHCP server. After a new IP is received, the DHCPv4 client performs IPv4 Duplicate Address Detection. If duplicate use is detected, the DHCPv4 client rejects the IP by sending a DHCPDECLINE packet and tries to obtain an IP address again. See RFC 5224. Defaults to "unset".
	SendDecline string `systemd:",omitempty"`

	// A whitespace-separated list of IPv4 addresses. DHCP offers from servers in the list are rejected. Note that if AllowList= is configured then DenyList= is ignored.
	DenyList string `systemd:",omitempty"`

	// A whitespace-separated list of IPv4 addresses. DHCP offers from servers in the list are accepted.
	AllowList string `systemd:",omitempty"`

	// When configured, allows to set arbitrary request options in the DHCPv4 request options list and will be sent to the DHCPV4 server. A whitespace-separated list of integers in the range 1..254. Defaults to unset.
	RequestOptions string `systemd:",omitempty"`

	// Send an arbitrary raw option in the DHCPv4 request. Takes a DHCP option number, data type and data separated with a colon ("option:type:value"). The option number must be an integer in the range 1..254. The type takes one of "uint8", "uint16", "uint32", "ipv4address", or "string". Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Defaults to unset.
	SendOption string `systemd:",omitempty"`

	// Send an arbitrary vendor option in the DHCPv4 request. Takes a DHCP option number, data type and data separated with a colon ("option:type:value"). The option number must be an integer in the range 1..254. The type takes one of "uint8", "uint16", "uint32", "ipv4address", or "string". Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Defaults to unset.
	SendVendorOption string `systemd:",omitempty"`
}

// The [DHCPv6] section configures the DHCPv6 client, if it is enabled with the DHCP= setting described above, or invoked by the IPv6 Router Advertisement:
type DHCPv6Section struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// As in the [DHCPv4] section.
	UseDNS string `systemd:",omitempty"`
	// As in the [DHCPv4] section.
	UseNTP string `systemd:",omitempty"`

	// Set the routing metric for routes specified by the DHCP server. Defaults to 1024.
	RouteMetric string `systemd:",omitempty"`

	// Takes a boolean. The DHCPv6 client can obtain configuration parameters from a DHCPv6 server through a rapid two-message exchange (solicit and reply). When the rapid commit option is enabled by both the DHCPv6 client and the DHCPv6 server, the two-message exchange is used, rather than the default four-message exchange (solicit, advertise, request, and reply). The two-message exchange provides faster client configuration and is beneficial in environments in which networks are under a heavy load. See RFC 3315 for details. Defaults to true.
	RapidCommit string `systemd:",omitempty"`

	// When configured, the Manufacturer Usage Descriptions (MUD) URL will be sent to the DHCPV6 server. Takes an URL of length up to 255 characters. A superficial verification that the string is a valid URL will be performed. DHCPv6 clients are intended to have at most one MUD URL associated with them. See RFC 8520.
	MUDURL string `systemd:",omitempty"`

	// When configured, allows to set arbitrary request options in the DHCPv6 request options list and will sent to the DHCPV6 server. A whitespace-separated list of integers in the range 1..254. Defaults to unset.
	RequestOptions string `systemd:",omitempty"`

	// Send an arbitrary vendor option in the DHCPv6 request. Takes an enterprise identifier, DHCP option number, data type, and data separated with a colon ("enterprise identifier:option:type: value"). Enterprise identifier is an unsigned integer in the range 1–4294967294. The option number must be an integer in the range 1–254. Data type takes one of "uint8", "uint16", "uint32", "ipv4address", "ipv6address", or "string". Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Defaults to unset.
	SendVendorOption string `systemd:",omitempty"`

	// Takes a boolean that enforces DHCPv6 stateful mode when the 'Other information' bit is set in Router Advertisement messages. By default setting only the 'O' bit in Router Advertisements makes DHCPv6 request network information in a stateless manner using a two-message Information Request and Information Reply message exchange. RFC 7084, requirement WPD-4, updates this behavior for a Customer Edge router so that stateful DHCPv6 Prefix Delegation is also requested when only the 'O' bit is set in Router Advertisements. This option enables such a CE behavior as it is impossible to automatically distinguish the intention of the 'O' bit otherwise. By default this option is set to 'false', enable it if no prefixes are delegated when the device should be acting as a CE router.
	ForceDHCPv6PDOtherInformation string `systemd:",omitempty"`

	// Takes an IPv6 address with prefix length in the same format as the Address= in the [Network] section. The DHCPv6 client will include a prefix hint in the DHCPv6 solicitation sent to the server. The prefix length must be in the range 1–128. Defaults to unset.
	PrefixDelegationHint string `systemd:",omitempty"`

	// Allows DHCPv6 client to start without router advertisements's managed or other address configuration flag. Takes one of "solicit" or "information-request". Defaults to unset.
	WithoutRA string `systemd:",omitempty"`

	// As in the [DHCPv4] section, however because DHCPv6 uses 16-bit fields to store option numbers, the option number is an integer in the range 1..65536.
	SendOption string `systemd:",omitempty"`

	// A DHCPv6 client can use User Class option to identify the type or category of user or applications it represents. The information contained in this option is a string that represents the user class of which the client is a member. Each class sets an identifying string of information to be used by the DHCP service to classify clients. Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Takes a whitespace-separated list of strings. Note that currently NUL bytes are not allowed.
	UserClass string `systemd:",omitempty"`

	// A DHCPv6 client can use VendorClass option to identify the vendor that manufactured the hardware on which the client is running. The information contained in the data area of this option is contained in one or more opaque fields that identify details of the hardware configuration. Takes a whitespace-separated list of strings.
	VendorClass string `systemd:",omitempty"`
}

// The [DHCPv6PrefixDelegation] section configures delegated prefix assigned by DHCPv6 server. The settings in this section are used only when IPv6PrefixDelegation= setting is enabled, or set to "dhcp6".
type DHCPv6PrefixDelegationSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configure a specific subnet ID on the interface from a (previously) received prefix delegation. You can either set "auto" (the default) or a specific subnet ID (as defined in RFC 4291, section 2.5.4), in which case the allowed value is hexadecimal, from 0 to 0x7fffffffffffffff inclusive. This option is only effective when used together with IPv6PrefixDelegation= and the corresponding configuration on the upstream interface.
	SubnetId string `systemd:",omitempty"`

	// Takes a boolean. Specifies whether to add an address from the delegated prefixes which are received from the WAN interface by the IPv6PrefixDelegation=. When true (on LAN interfce), the EUI-64 algorithm will be used to form an interface identifier from the delegated prefixes. Defaults to true.
	Assign string `systemd:",omitempty"`

	// Specifies an optional address generation mode for Assign=. Takes an IPv6 address. When set, the lower bits of the supplied address are combined with the upper bits of a delegatad prefix received from the WAN interface by the IPv6PrefixDelegation= prefixes to form a complete address.
	Token string `systemd:",omitempty"`
}

// The [IPv6AcceptRA] section configures the IPv6 Router Advertisement (RA) client, if it is enabled with the IPv6AcceptRA= setting described above:
type IPv6AcceptRASection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// When true (the default), the DNS servers received in the Router Advertisement will be used and take precedence over any statically configured ones.
	UseDNS string `systemd:",omitempty"`
	//
	// This corresponds to the nameserver option in resolv.conf(5).

	// Takes a boolean, or the special value "route". When true, the domain name received via IPv6 Router Advertisement (RA) will be used as DNS search domain over this link, similar to the effect of the Domains= setting. If set to "route", the domain name received via IPv6 RA will be used for routing DNS queries only, but not for searching, similar to the effect of the Domains= setting when the argument is prefixed with "~". Defaults to false.
	UseDomains string `systemd:",omitempty"`
	//
	// It is recommended to enable this option only on trusted networks, as setting this affects resolution of all hostnames, in particular of single-label names. It is generally safer to use the supplied domain only as routing domain, rather than as search domain, in order to not have it affect local resolution of single-label names.
	//
	// When set to true, this setting corresponds to the domain option in resolv.conf(5).

	// The table identifier for the routes received in the Router Advertisement (a number between 1 and 4294967295, or 0 to unset). The table can be retrieved using ip route show table num.
	RouteTable string `systemd:",omitempty"`

	// When true (the default), the autonomous prefix received in the Router Advertisement will be used and take precedence over any statically configured ones.
	UseAutonomousPrefix string `systemd:",omitempty"`

	// When true (the default), the onlink prefix received in the Router Advertisement will be used and take precedence over any statically configured ones.
	UseOnLinkPrefix string `systemd:",omitempty"`

	// A whitespace-separated list of IPv6 prefixes. IPv6 prefixes supplied via router advertisements in the list are ignored.
	DenyList string `systemd:",omitempty"`

	// Takes a boolean, or the special value "always". When true (the default), the DHCPv6 client will be started when the RA has the managed or other information flag. If set to "always", the DHCPv6 client will be started even if there is no managed or other information flag in the RA.
	DHCPv6Client string `systemd:",omitempty"`
}

// The [DHCPServer] section contains settings for the DHCP server, if enabled via the DHCPServer= option described above:
type DHCPServerSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the pool of addresses to hand out. The pool is a contiguous sequence of IP addresses in the subnet configured for the server address, which does not include the subnet nor the broadcast address. PoolOffset= takes the offset of the pool from the start of subnet, or zero to use the default value. PoolSize= takes the number of IP addresses in the pool or zero to use the default value. By default, the pool starts at the first address after the subnet address and takes up the rest of the subnet, excluding the broadcast address. If the pool includes the server address (the default), this is reserved and not handed out to clients.
	PoolOffset, PoolSize string `systemd:",omitempty"`

	// Control the default and maximum DHCP lease time to pass to clients. These settings take time values in seconds or another common time unit, depending on the suffix. The default lease time is used for clients that did not ask for a specific lease time. If a client asks for a lease time longer than the maximum lease time, it is automatically shortened to the specified time. The default lease time defaults to 1h, the maximum lease time to 12h. Shorter lease times are beneficial if the configuration data in DHCP leases changes frequently and clients shall learn the new settings with shorter latencies. Longer lease times reduce the generated DHCP network traffic.
	DefaultLeaseTimeSec, MaxLeaseTimeSec string `systemd:",omitempty"`

	// EmitDNS= takes a boolean. Configures whether the DHCP leases handed out to clients shall contain DNS server information. Defaults to "yes". The DNS servers to pass to clients may be configured with the DNS= option, which takes a list of IPv4 addresses. If the EmitDNS= option is enabled but no servers configured, the servers are automatically propagated from an "uplink" interface that has appropriate servers set. The "uplink" interface is determined by the default route of the system with the highest priority. Note that this information is acquired at the time the lease is handed out, and does not take uplink interfaces into account that acquire DNS server information at a later point. If no suitable uplinkg interface is found the DNS server data from /etc/resolv.conf is used. Also, note that the leases are not refreshed if the uplink network configuration changes. To ensure clients regularly acquire the most current uplink DNS server information, it is thus advisable to shorten the DHCP lease time via MaxLeaseTimeSec= described above.
	EmitDNS, DNS string `systemd:",omitempty"`

	// Similar to the EmitDNS= and DNS= settings described above, these settings configure whether and what server information for the indicate protocol shall be emitted as part of the DHCP lease. The same syntax, propagation semantics and defaults apply as for EmitDNS= and DNS=.
	EmitNTP, NTP, EmitSIP, SIP, EmitPOP3, POP3, EmitSMTP, SMTP, EmitLPR, LPR string `systemd:",omitempty"`

	// Similar to the EmitDNS= setting described above, this setting configures whether the DHCP lease should contain the router option. The same syntax, propagation semantics and defaults apply as for EmitDNS=.
	EmitRouter string `systemd:",omitempty"`

	// Takes a boolean. Configures whether the DHCP leases handed out to clients shall contain timezone information. Defaults to "yes". The Timezone= setting takes a timezone string (such as "Europe/Berlin" or "UTC") to pass to clients. If no explicit timezone is set, the system timezone of the local host is propagated, as determined by the /etc/localtime symlink.
	EmitTimezone, Timezone string `systemd:",omitempty"`

	// Send a raw option with value via DHCPv4 server. Takes a DHCP option number, data type and data ("option:type:value"). The option number is an integer in the range 1..254. The type takes one of "uint8", "uint16", "uint32", "ipv4address", "ipv6address", or "string". Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Defaults to unset.
	SendOption string `systemd:",omitempty"`

	// Send a vendor option with value via DHCPv4 server. Takes a DHCP option number, data type and data ("option:type:value"). The option number is an integer in the range 1..254. The type takes one of "uint8", "uint16", "uint32", "ipv4address", or "string". Special characters in the data string may be escaped using C-style escapes. This setting can be specified multiple times. If an empty string is specified, then all options specified earlier are cleared. Defaults to unset.
	SendVendorOption string `systemd:",omitempty"`
}

// The [IPv6PrefixDelegation] section contains settings for sending IPv6 Router Advertisements and whether to act as a router, if enabled via the IPv6PrefixDelegation= option described above. IPv6 network prefixes are defined with one or more [IPv6Prefix] sections.
type IPv6PrefixDelegationSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a boolean. Controls whether a DHCPv6 server is used to acquire IPv6 addresses on the network link when Managed= is set to "true" or if only additional network information can be obtained via DHCPv6 for the network link when OtherInformation= is set to "true". Both settings default to "false", which means that a DHCPv6 server is not being used.
	Managed, OtherInformation string `systemd:",omitempty"`

	// Takes a timespan. Configures the IPv6 router lifetime in seconds. If set, this host also announces itself in Router Advertisements as an IPv6 router for the network link. When unset, the host is not acting as a router.
	RouterLifetimeSec string `systemd:",omitempty"`

	// Configures IPv6 router preference if RouterLifetimeSec= is non-zero. Valid values are "high", "medium" and "low", with "normal" and "default" added as synonyms for "medium" just to make configuration easier. See RFC 4191 for details. Defaults to "medium".
	RouterPreference string `systemd:",omitempty"`

	// DNS= specifies a list of recursive DNS server IPv6 addresses that are distributed via Router Advertisement messages when EmitDNS= is true. DNS= also takes special value "_link_local"; in that case the IPv6 link local address is distributed. If DNS= is empty, DNS servers are read from the [Network] section. If the [Network] section does not contain any DNS servers either, DNS servers from the uplink with the highest priority default route are used. When EmitDNS= is false, no DNS server information is sent in Router Advertisement messages. EmitDNS= defaults to true.
	EmitDNS, DNS string `systemd:",omitempty"`

	// A list of DNS search domains distributed via Router Advertisement messages when EmitDomains= is true. If Domains= is empty, DNS search domains are read from the [Network] section. If the [Network] section does not contain any DNS search domains either, DNS search domains from the uplink with the highest priority default route are used. When EmitDomains= is false, no DNS search domain information is sent in Router Advertisement messages. EmitDomains= defaults to true.
	EmitDomains, Domains string `systemd:",omitempty"`

	// Lifetime in seconds for the DNS server addresses listed in DNS= and search domains listed in Domains=.
	DNSLifetimeSec string `systemd:",omitempty"`
}

// One or more [IPv6Prefix] sections contain the IPv6 prefixes that are announced via Router Advertisements. See RFC 4861 for further details.
type IPv6PrefixSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a boolean to specify whether IPv6 addresses can be autoconfigured with this prefix and whether the prefix can be used for onlink determination. Both settings default to "true" in order to ease configuration.
	AddressAutoconfiguration, OnLink string `systemd:",omitempty"`

	// The IPv6 prefix that is to be distributed to hosts. Similarly to configuring static IPv6 addresses, the setting is configured as an IPv6 prefix and its prefix length, separated by a "/" character. Use multiple [IPv6Prefix] sections to configure multiple IPv6 prefixes since prefix lifetimes, address autoconfiguration and onlink status may differ from one prefix to another.
	Prefix string `systemd:",omitempty"`

	// Preferred and valid lifetimes for the prefix measured in seconds. PreferredLifetimeSec= defaults to 604800 seconds (one week) and ValidLifetimeSec= defaults to 2592000 seconds (30 days).
	PreferredLifetimeSec, ValidLifetimeSec string `systemd:",omitempty"`

	// Takes a boolean. When true, adds an address from the prefix. Default to false.
	Assign string `systemd:",omitempty"`
}

// One or more [IPv6RoutePrefix] sections contain the IPv6 prefix routes that are announced via Router Advertisements. See RFC 4191 for further details.
type IPv6RoutePrefixSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The IPv6 route that is to be distributed to hosts. Similarly to configuring static IPv6 routes, the setting is configured as an IPv6 prefix routes and its prefix route length, separated by a "/" character. Use multiple [IPv6PrefixRoutes] sections to configure multiple IPv6 prefix routes.
	Route string `systemd:",omitempty"`

	// Lifetime for the route prefix measured in seconds. LifetimeSec= defaults to 604800 seconds (one week).
	LifetimeSec string `systemd:",omitempty"`
}

type BridgeSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Takes a boolean. Controls whether the bridge should flood traffic for which an FDB entry is missing and the destination is unknown through this port. When unset, the kernel's default will be used.
	UnicastFlood string `systemd:",omitempty"`

	// Takes a boolean. Controls whether the bridge should flood traffic for which an MDB entry is missing and the destination is unknown through this port. When unset, the kernel's default will be used.
	MulticastFlood string `systemd:",omitempty"`

	// Takes a boolean. Multicast to unicast works on top of the multicast snooping feature of the bridge. Which means unicast copies are only delivered to hosts which are interested in it. When unset, the kernel's default will be used.
	MulticastToUnicast string `systemd:",omitempty"`

	// Takes a boolean. Configures whether ARP and ND neighbor suppression is enabled for this port. When unset, the kernel's default will be used.
	NeighborSuppression string `systemd:",omitempty"`

	// Takes a boolean. Configures whether MAC address learning is enabled for this port. When unset, the kernel's default will be used.
	Learning string `systemd:",omitempty"`

	// Takes a boolean. Configures whether traffic may be sent back out of the port on which it was received. When this flag is false, then the bridge will not forward traffic back out of the receiving port. When unset, the kernel's default will be used.
	HairPin string `systemd:",omitempty"`

	// Takes a boolean. Configures whether STP Bridge Protocol Data Units will be processed by the bridge port. When unset, the kernel's default will be used.
	UseBPDU string `systemd:",omitempty"`

	// Takes a boolean. This flag allows the bridge to immediately stop multicast traffic on a port that receives an IGMP Leave message. It is only used with IGMP snooping if enabled on the bridge. When unset, the kernel's default will be used.
	FastLeave string `systemd:",omitempty"`

	// Takes a boolean. Configures whether a given port is allowed to become a root port. Only used when STP is enabled on the bridge. When unset, the kernel's default will be used.
	AllowPortToBeRoot string `systemd:",omitempty"`

	// Takes a boolean. Configures whether proxy ARP to be enabled on this port. When unset, the kernel's default will be used.
	ProxyARP string `systemd:",omitempty"`

	// Takes a boolean. Configures whether proxy ARP to be enabled on this port which meets extended requirements by IEEE 802.11 and Hotspot 2.0 specifications. When unset, the kernel's default will be used.
	ProxyARPWiFi string `systemd:",omitempty"`

	// Configures this port for having multicast routers attached. A port with a multicast router will receive all multicast traffic. Takes one of "no" to disable multicast routers on this port, "query" to let the system detect the presence of routers, "permanent" to permanently enable multicast traffic forwarding on this port, or "temporary" to enable multicast routers temporarily on this port, not depending on incoming queries. When unset, the kernel's default will be used.
	MulticastRouter string `systemd:",omitempty"`

	// Sets the "cost" of sending packets of this interface. Each port in a bridge may have a different speed and the cost is used to decide which link to use. Faster interfaces should have lower costs. It is an integer value between 1 and 65535.
	Cost string `systemd:",omitempty"`

	// Sets the "priority" of sending packets on this interface. Each port in a bridge may have a different priority which is used to decide which link to use. Lower value means higher priority. It is an integer value between 0 to 63. Networkd does not set any default, meaning the kernel default value of 32 is used.
	Priority string `systemd:",omitempty"`
}

// The [BridgeFDB] section manages the forwarding database table of a port and accepts the following keys. Specify several [BridgeFDB] sections to configure several static MAC table entries.
type BridgeFDBSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// As in the [Network] section. This key is mandatory.
	MACAddress string `systemd:",omitempty"`

	// Takes an IP address of the destination VXLAN tunnel endpoint.
	Destination string `systemd:",omitempty"`

	// The VLAN ID for the new static MAC table entry. If omitted, no VLAN ID information is appended to the new static MAC table entry.
	VLANId string `systemd:",omitempty"`

	// The VXLAN Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint. Takes a number in the range 1-16777215. Defaults to unset.
	VNI string `systemd:",omitempty"`

	// Specifies where the address is associated with. Takes one of "use", "self", "master" or "router". "use" means the address is in use. User space can use this option to indicate to the kernel that the fdb entry is in use. "self" means the address is associated with the port drivers fdb. Usually hardware. "master" means the address is associated with master devices fdb. "router" means the destination address is associated with a router. Note that it's valid if the referenced device is a VXLAN type device and has route shortcircuit enabled. Defaults to "self".
	AssociatedWith string `systemd:",omitempty"`
}

// The [LLDP] section manages the Link Layer Discovery Protocol (LLDP) and accepts the following keys.
type LLDPSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Controls support for Ethernet LLDP packet's Manufacturer Usage Description (MUD). MUD is an embedded software standard defined by the IETF that allows IoT Device makers to advertise device specifications, including the intended communication patterns for their device when it connects to the network. The network can then use this intent to author a context-specific access policy, so the device functions only within those parameters. Takes an URL of length up to 255 characters. A superficial verification that the string is a valid URL will be performed. See RFC 8520 for details. The MUD URL received from the LLDP packets will be saved at the state files and can be read via sd_lldp_neighbor_get_mud_url() function.
	MUDURL string `systemd:",omitempty"`
}

// The [CAN] section manages the Controller Area Network (CAN bus) and accepts the following keys:
type CANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The bitrate of CAN device in bits per second. The usual SI prefixes (K, M) with the base of 1000 can be used here. Takes a number in the range 1..4294967295.
	BitRate string `systemd:",omitempty"`

	// Optional sample point in percent with one decimal (e.g. "75%", "87.5%") or permille (e.g. "875‰").
	SamplePoint string `systemd:",omitempty"`

	// The bitrate and sample point for the data phase, if CAN-FD is used. These settings are analogous to the BitRate= and SamplePoint= keys.
	DataBitRate, DataSamplePoint string `systemd:",omitempty"`

	// Takes a boolean. When "yes", CAN-FD mode is enabled for the interface. Note, that a bitrate and optional sample point should also be set for the CAN-FD data phase using the DataBitRate= and DataSamplePoint= keys.
	FDMode string `systemd:",omitempty"`

	// Takes a boolean. When "yes", non-ISO CAN-FD mode is enabled for the interface. When unset, the kernel's default will be used.
	FDNonISO string `systemd:",omitempty"`

	// Automatic restart delay time. If set to a non-zero value, a restart of the CAN controller will be triggered automatically in case of a bus-off condition after the specified delay time. Subsecond delays can be specified using decimals (e.g. "0.1s") or a "ms" or "us" postfix. Using "infinity" or "0" will turn the automatic restart off. By default automatic restart is disabled.
	RestartSec string `systemd:",omitempty"`

	// Takes a boolean. When "yes", the termination resistor will be selected for the bias network. When unset, the kernel's default will be used.
	Termination string `systemd:",omitempty"`

	// Takes a boolean. When "yes", three samples (instead of one) are used to determine the value of a received bit by majority rule. When unset, the kernel's default will be used.
	TripleSampling string `systemd:",omitempty"`

	// Takes a boolean. When "yes", listen-only mode is enabled. When the interface is in listen-only mode, the interface neither transmit CAN frames nor send ACK bit. Listen-only mode is important to debug CAN networks without interfering with the communication or acknowledge the CAN frame. When unset, the kernel's default will be used.
	ListenOnly string `systemd:",omitempty"`
}

// The [QDisc] section manages the traffic control queueing discipline (qdisc).
type QDiscSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Specifies the parent Queueing Discipline (qdisc). Takes one of "clsact" or "ingress". This is mandatory.
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`
}

// The [NetworkEmulator] section manages the queueing discipline (qdisc) of the network emulator. It can be used to configure the kernel packet scheduler and simulate packet delay and loss for UDP or TCP applications, or limit the bandwidth usage of a particular service to simulate internet connections.
type NetworkEmulatorSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the fixed amount of delay to be added to all packets going out of the interface. Defaults to unset.
	DelaySec string `systemd:",omitempty"`

	// Specifies the chosen delay to be added to the packets outgoing to the network interface. Defaults to unset.
	DelayJitterSec string `systemd:",omitempty"`

	// Specifies the maximum number of packets the qdisc may hold queued at a time. An unsigned integer in the range 0–4294967294. Defaults to 1000.
	PacketLimit string `systemd:",omitempty"`

	// Specifies an independent loss probability to be added to the packets outgoing from the network interface. Takes a percentage value, suffixed with "%". Defaults to unset.
	LossRate string `systemd:",omitempty"`

	// Specifies that the chosen percent of packets is duplicated before queuing them. Takes a percentage value, suffixed with "%". Defaults to unset.
	DuplicateRate string `systemd:",omitempty"`
}

// The [TokenBucketFilter] section manages the queueing discipline (qdisc) of token bucket filter (tbf).
type TokenBucketFilterSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the latency parameter, which specifies the maximum amount of time a packet can sit in the Token Bucket Filter (TBF). Defaults to unset.
	LatencySec string `systemd:",omitempty"`

	// Takes the number of bytes that can be queued waiting for tokens to become available. When the size is suffixed with K, M, or G, it is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset.
	LimitBytes string `systemd:",omitempty"`

	// Specifies the size of the bucket. This is the maximum amount of bytes that tokens can be available for instantaneous transfer. When the size is suffixed with K, M, or G, it is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset.
	BurstBytes string `systemd:",omitempty"`

	// Specifies the device specific bandwidth. When suffixed with K, M, or G, the specified bandwidth is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. Defaults to unset.
	Rate string `systemd:",omitempty"`

	// The Minimum Packet Unit (MPU) determines the minimal token usage (specified in bytes) for a packet. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to zero.
	MPUBytes string `systemd:",omitempty"`

	// Takes the maximum depletion rate of the bucket. When suffixed with K, M, or G, the specified size is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. Defaults to unset.
	PeakRate string `systemd:",omitempty"`

	// Specifies the size of the peakrate bucket. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset.
	MTUBytes string `systemd:",omitempty"`
}

// The [PIE] section manages the queueing discipline (qdisc) of Proportional Integral controller-Enhanced (PIE).
type PIESection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the queue size in number of packets. When this limit is reached, incoming packets are dropped. An unsigned integer in the range 1–4294967294. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`
}

// The [StochasticFairBlue] section manages the queueing discipline (qdisc) of stochastic fair blue (sfb).
type StochasticFairBlueSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the queue size in number of packets. When this limit is reached, incoming packets are dropped. An unsigned integer in the range 0–4294967294. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`
}

// The [StochasticFairnessQueueing] section manages the queueing discipline (qdisc) of stochastic fairness queueing (sfq).
type StochasticFairnessQueueingSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the interval in seconds for queue algorithm perturbation. Defaults to unset.
	PerturbPeriodSec string `systemd:",omitempty"`
}

// The [BFIFO] section manages the queueing discipline (qdisc) of Byte limited Packet First In First Out (bfifo).
type BFIFOSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the FIFO size in bytes. The size limit (a buffer size) to prevent it from overflowing in case it is unable to dequeue packets as quickly as it receives them. When this limit is reached, incoming packets are dropped. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset and kernel's default is used.
	LimitBytes string `systemd:",omitempty"`
}

// The [PFIFO] section manages the queueing discipline (qdisc) of Packet First In First Out (pfifo).
type PFIFOSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the FIFO size in number of packets. The size limit (a buffer size) to prevent it from overflowing in case it is unable to dequeue packets as quickly as it receives them. When this limit is reached, incoming packets are dropped. An unsigned integer in the range 0–4294967294. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`
}

// The [PFIFOHeadDrop] section manages the queueing discipline (qdisc) of Packet First In First Out Head Drop (pfifo_head_drop).
type PFIFOHeadDropSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// As in [PFIFO] section.
	PacketLimit string `systemd:",omitempty"`
}

// The [PFIFOFast] section manages the queueing discipline (qdisc) of Packet First In First Out Fast (pfifo_fast).
type PFIFOFastSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`
}

// The [CAKE] section manages the queueing discipline (qdisc) of Common Applications Kept Enhanced (CAKE).
type CAKESection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies that bytes to be addeded to the size of each packet. Bytes may be negative. Takes an integer in the range from -64 to 256. Defaults to unset and kernel's default is used.
	OverheadBytes string `systemd:",omitempty"`

	// Specifies the shaper bandwidth. When suffixed with K, M, or G, the specified size is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. Defaults to unset and kernel's default is used.
	Bandwidth string `systemd:",omitempty"`
}

// The [ControlledDelay] section manages the queueing discipline (qdisc) of controlled delay (CoDel).
type ControlledDelaySection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the queue size in number of packets. When this limit is reached, incoming packets are dropped. An unsigned integer in the range 0–4294967294. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`

	// Takes a timespan. Specifies the acceptable minimum standing/persistent queue delay. Defaults to unset and kernel's default is used.
	TargetSec string `systemd:",omitempty"`

	// Takes a timespan. This is used to ensure that the measured minimum delay does not become too stale. Defaults to unset and kernel's default is used.
	IntervalSec string `systemd:",omitempty"`

	// Takes a boolean. This can be used to mark packets instead of dropping them. Defaults to unset and kernel's default is used.
	ECN string `systemd:",omitempty"`

	// Takes a timespan. This sets a threshold above which all packets are marked with ECN Congestion Experienced (CE). Defaults to unset and kernel's default is used.
	CEThresholdSec string `systemd:",omitempty"`
}

// The [DeficitRoundRobinScheduler] section manages the queueing discipline (qdisc) of Deficit Round Robin Scheduler (DRR).
type DeficitRoundRobinSchedulerSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`
}

// The [DeficitRoundRobinSchedulerClass] section manages the traffic control class of Deficit Round Robin Scheduler (DRR).
type DeficitRoundRobinSchedulerClassSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", or a qdisc identifier. The qdisc identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configues the unique identifier of the class. It is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to unset.
	ClassId string `systemd:",omitempty"`

	// Specifies the amount of bytes a flow is allowed to dequeue before the scheduler moves to the next class. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to the MTU of the interface.
	QuantumBytes string `systemd:",omitempty"`
}

// The [EnhancedTransmissionSelection] section manages the queueing discipline (qdisc) of Enhanced Transmission Selection (ETS).
type EnhancedTransmissionSelectionSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the number of bands. An unsigned integer in the range 1–16. This value has to be at least large enough to cover the strict bands specified through the StrictBands= and bandwidth-sharing bands specified in QuantumBytes=.
	Bands string `systemd:",omitempty"`

	// Specifies the number of bands that should be created in strict mode. An unsigned integer in the range 1–16.
	StrictBands string `systemd:",omitempty"`

	// Specifies the white-space separated list of quantum used in band-sharing bands. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. This setting can be specified multiple times. If an empty string is assigned, then the all previous assignments are cleared.
	QuantumBytes string `systemd:",omitempty"`

	// The priority map maps the priority of a packet to a band. The argument is a white-space separated list of numbers. The first number indicates which band the packets with priority 0 should be put to, the second is for priority 1, and so on. There can be up to 16 numbers in the list. If there are fewer, the default band that traffic with one of the unmentioned priorities goes to is the last one. Each band number must be 0..255. This setting can be specified multiple times. If an empty string is assigned, then the all previous assignments are cleared.
	PriorityMap string `systemd:",omitempty"`
}

// The [GenericRandomEarlyDetection] section manages the queueing discipline (qdisc) of Generic Random Early Detection (GRED).
type GenericRandomEarlyDetectionSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the number of virtual queues. Takes a integer in the range 1-16. Defaults to unset and kernel's default is used.
	VirtualQueues string `systemd:",omitempty"`

	// Specifies the number of default virtual queue. This must be less than VirtualQueue=. Defaults to unset and kernel's default is used.
	DefaultVirtualQueue string `systemd:",omitempty"`

	// Takes a boolean. It turns on the RIO-like buffering scheme. Defaults to unset and kernel's default is used.
	GenericRIO string `systemd:",omitempty"`
}

// The [FairQueueingControlledDelay] section manages the queueing discipline (qdisc) of fair queuing controlled delay (FQ-CoDel).
type FairQueueingControlledDelaySection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the real queue size. When this limit is reached, incoming packets are dropped. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`

	// Specifies the limit on the total number of bytes that can be queued in this FQ-CoDel instance. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset and kernel's default is used.
	MemoryLimitBytes string `systemd:",omitempty"`

	// Specifies the number of flows into which the incoming packets are classified. Defaults to unset and kernel's default is used.
	Flows string `systemd:",omitempty"`

	// Takes a timespan. Specifies the acceptable minimum standing/persistent queue delay. Defaults to unset and kernel's default is used.
	TargetSec string `systemd:",omitempty"`

	// Takes a timespan. This is used to ensure that the measured minimum delay does not become too stale. Defaults to unset and kernel's default is used.
	IntervalSec string `systemd:",omitempty"`

	// Specifies the number of bytes used as the "deficit" in the fair queuing algorithm timespan. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset and kernel's default is used.
	QuantumBytes string `systemd:",omitempty"`

	// Takes a boolean. This can be used to mark packets instead of dropping them. Defaults to unset and kernel's default is used.
	ECN string `systemd:",omitempty"`

	// Takes a timespan. This sets a threshold above which all packets are marked with ECN Congestion Experienced (CE). Defaults to unset and kernel's default is used.
	CEThresholdSec string `systemd:",omitempty"`
}

// The [FairQueueing] section manages the queueing discipline (qdisc) of fair queue traffic policing (FQ).
type FairQueueingSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the real queue size. When this limit is reached, incoming packets are dropped. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`

	// Specifies the hard limit on the maximum number of packets queued per flow. Defaults to unset and kernel's default is used.
	FlowLimit string `systemd:",omitempty"`

	// Specifies the credit per dequeue RR round, i.e. the amount of bytes a flow is allowed to dequeue at once. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset and kernel's default is used.
	QuantumBytes string `systemd:",omitempty"`

	// Specifies the initial sending rate credit, i.e. the amount of bytes a new flow is allowed to dequeue initially. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. Defaults to unset and kernel's default is used.
	InitialQuantumBytes string `systemd:",omitempty"`

	// Specifies the maximum sending rate of a flow. When suffixed with K, M, or G, the specified size is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. Defaults to unset and kernel's default is used.
	MaximumRate string `systemd:",omitempty"`

	// Specifies the size of the hash table used for flow lookups. Defaults to unset and kernel's default is used.
	Buckets string `systemd:",omitempty"`

	// Takes an unsigned integer. For packets not owned by a socket, fq is able to mask a part of hash and reduce number of buckets associated with the traffic. Defaults to unset and kernel's default is used.
	OrphanMask string `systemd:",omitempty"`

	// Takes a boolean, and enables or disables flow pacing. Defaults to unset and kernel's default is used.
	Pacing string `systemd:",omitempty"`

	// Takes a timespan. This sets a threshold above which all packets are marked with ECN Congestion Experienced (CE). Defaults to unset and kernel's default is used.
	CEThresholdSec string `systemd:",omitempty"`
}

// The [TrivialLinkEqualizer] section manages the queueing discipline (qdisc) of trivial link equalizer (teql).
type TrivialLinkEqualizerSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the interface ID "N" of teql. Defaults to "0". Note that when teql is used, currently, the module sch_teql with max_equalizers=N+1 option must be loaded before systemd-networkd is started.
	Id string `systemd:",omitempty"`
}

// The [HierarchyTokenBucket] section manages the queueing discipline (qdisc) of hierarchy token bucket (htb).
type HierarchyTokenBucketSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Takes the minor id in hexadecimal of the default class. Unclassified traffic gets sent to the class. Defaults to unset.
	DefaultClass string `systemd:",omitempty"`

	// Takes an unsigned integer. The DRR quantums are calculated by dividing the value configured in Rate= by RateToQuantum=.
	RateToQuantum string `systemd:",omitempty"`
}

// The [HierarchyTokenBucketClass] section manages the traffic control class of hierarchy token bucket (htb).
type HierarchyTokenBucketClassSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", or a qdisc identifier. The qdisc identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configues the unique identifier of the class. It is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to unset.
	ClassId string `systemd:",omitempty"`

	// Specifies the priority of the class. In the round-robin process, classes with the lowest priority field are tried for packets first.
	Priority string `systemd:",omitempty"`

	// Specifies how many bytes to serve from leaf at once. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024.
	QuantumBytes string `systemd:",omitempty"`

	// Specifies the maximum packet size we create. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024.
	MTUBytes string `systemd:",omitempty"`

	// Takes an unsigned integer which specifies per-packet size overhead used in rate computations. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024.
	OverheadBytes string `systemd:",omitempty"`

	// Specifies the maximum rate this class and all its children are guaranteed. When suffixed with K, M, or G, the specified size is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. This setting is mandatory.
	Rate string `systemd:",omitempty"`

	// Specifies the maximum rate at which a class can send, if its parent has bandwidth to spare. When suffixed with K, M, or G, the specified size is parsed as Kilobits, Megabits, or Gigabits, respectively, to the base of 1000. When unset, the value specified with Rate= is used.
	CeilRate string `systemd:",omitempty"`

	// Specifies the maximum bytes burst which can be accumulated during idle period. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024.
	BufferBytes string `systemd:",omitempty"`

	// Specifies the maximum bytes burst for ceil which can be accumulated during idle period. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024.
	CeilBufferBytes string `systemd:",omitempty"`
}

// The [HeavyHitterFilter] section manages the queueing discipline (qdisc) of Heavy Hitter Filter (hhf).
type HeavyHitterFilterSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`

	// Specifies the hard limit on the queue size in number of packets. When this limit is reached, incoming packets are dropped. An unsigned integer in the range 0–4294967294. Defaults to unset and kernel's default is used.
	PacketLimit string `systemd:",omitempty"`
}

// The [QuickFairQueueing] section manages the queueing discipline (qdisc) of Quick Fair Queueing (QFQ).
type QuickFairQueueingSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", "clsact", "ingress" or a class identifier. The class identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configures the major number of unique identifier of the qdisc, known as the handle. Takes a hexadecimal number in the range 0x1–0xffff. Defaults to unset.
	Handle string `systemd:",omitempty"`
}

// The [QuickFairQueueingClass] section manages the traffic control class of Quick Fair Queueing (qfq).
type QuickFairQueueingClassSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// Configures the parent Queueing Discipline (qdisc). Takes one of "root", or a qdisc identifier. The qdisc identifier is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to "root".
	Parent string `systemd:",omitempty"`

	// Configues the unique identifier of the class. It is specified as the major and minor numbers in hexadecimal in the range 0x1–Oxffff separated with a colon ("major:minor"). Defaults to unset.
	ClassId string `systemd:",omitempty"`

	// Specifies the weight of the class. Takes an integer in the range 1..1023. Defaults to unset in which case the kernel default is used.
	Weight string `systemd:",omitempty"`

	// Specifies the maximum packet size in bytes for the class. When suffixed with K, M, or G, the specified size is parsed as Kilobytes, Megabytes, or Gigabytes, respectively, to the base of 1024. When unset, the kernel default is used.
	MaxPacketBytes string `systemd:",omitempty"`
}

// The [BridgeVLAN] section manages the VLAN ID configuration of a bridge port and accepts the following keys. Specify several [BridgeVLAN] sections to configure several VLAN entries. The VLANFiltering= option has to be enabled, see the [Bridge] section in systemd.netdev(5).
type BridgeVLANSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// The VLAN ID allowed on the port. This can be either a single ID or a range M-N. VLAN IDs are valid from 1 to 4094.
	VLAN string `systemd:",omitempty"`

	// The VLAN ID specified here will be used to untag frames on egress. Configuring EgressUntagged= implicates the use of VLAN= above and will enable the VLAN ID for ingress as well. This can be either a single ID or a range M-N.
	EgressUntagged string `systemd:",omitempty"`

	// The Port VLAN ID specified here is assigned to all untagged frames at ingress. PVID= can be used only once. Configuring PVID= implicates the use of VLAN= above and will enable the VLAN ID for ingress as well.
	PVID string `systemd:",omitempty"`
}
