package link

import "routerd.net/go-systemd"

type Link struct {
	systemd.SectionList // SectionList to store unknown sections

	Match       *MatchSection
	LinkSection *LinkSection
}

// A link file is said to match a device if all matches specified by the [Match] section are satisfied. When a link file does not contain valid settings in [Match] section, then the file will match all devices and systemd-udevd warns about that.
// Hint: to avoid the warning and to make it clear that all interfaces shall be matched, add the following:
// OriginalName=*
type MatchSection struct {
	systemd.KeyList        // KeyList to store unknown keys
	Comment         string // Section Comment
	systemd.KeyComments

	// A whitespace-separated list of hardware addresses. Use full colon-, hyphen- or dot-delimited hexadecimal. See the example below. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	// Example:
	// MACAddress=01:23:45:67:89:ab 00-11-22-33-44-55 AABB.CCDD.EEFF
	MACAddresses []string `systemd:"MACAddress,omitempty,wslist"`

	// A whitespace-separated list of hardware's permanent addresses. While MACAddress= matches the device's current MAC address, this matches the device's permanent MAC address, which may be different from the current one. Use full colon-, hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	PermanentMACAddresses []string `systemd:"PermanentMACAddress,omitempty,wslist"`

	// A whitespace-separated list of shell-style globs matching the persistent path, as exposed by the udev property ID_PATH.
	Paths []string `systemd:"Path,omitempty,wslist"`

	// A whitespace-separated list of shell-style globs matching the driver currently bound to the device, as exposed by the udev property ID_NET_DRIVER of its parent device, or if that is not set, the driver as exposed by ethtool -i of the device itself. If the list is prefixed with a "!", the test is inverted.
	Drivers []string `systemd:"Driver,omitempty,wslist"`

	// A whitespace-separated list of shell-style globs matching the device type, as exposed by networkctl list. If the list is prefixed with a "!", the test is inverted. Some valid values are "ether", "loopback", "wlan", "wwan". Valid types are named either from the udev "DEVTYPE" attribute, or "ARPHRD_" macros in linux/if_arp.h, so this is not comprehensive.
	Types []string `systemd:"Type,omitempty,wslist"`

	// A whitespace-separated list of udev property names with their values after equals sign ("="). If multiple properties are specified, the test results are ANDed. If the list is prefixed with a "!", the test is inverted. If a value contains white spaces, then please quote whole key and value pair. If a value contains quotation, then please escape the quotation with "\".
	// Example: if a .link file has the following:
	// Property=ID_MODEL_ID=9999 "ID_VENDOR_FROM_DATABASE=vendor name" "KEY=with \"quotation\""
	// then, the .link file matches only when an interface has all the above three properties.
	Properties []string `systemd:"Property,omitempty,wslist"`

	// A whitespace-separated list of shell-style globs matching the device name, as exposed by the udev property "INTERFACE". This cannot be used to match on names that have already been changed from userspace. Caution is advised when matching on kernel-assigned names, as they are known to be unstable between reboots.
	OriginalNames []string `systemd:"OriginalName,omitempty,wslist"`

	// Matches against the hostname or machine ID of the host. See ConditionHost= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Host *string `systemd:",omitempty"`

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation. See ConditionVirtualization= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Virtualization *string `systemd:",omitempty"`

	// Checks whether a specific kernel command line option is set. See ConditionKernelCommandLine= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelCommandLine *string `systemd:",omitempty"`

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See ConditionKernelVersion= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	KernelVersion *string `systemd:",omitempty"`

	// Checks whether the system is running on a specific architecture. See ConditionArchitecture= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Architecture *string `systemd:",omitempty"`

	// Checks whether the system is running on a machine with the specified firmware. See ConditionFirmware= in systemd.unit(5) for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is cleared.
	Firmware *string `systemd:",omitempty"`
}

type LinkSection struct {
	// A description of the device.
	Description *string `systemd:",omitempty"`

	// The ifalias interface property is set to this value.
	Alias *string `systemd:",omitempty"`

	// The policy by which the MAC address should be set. The available policies are:
	// persistent
	// If the hardware has a persistent MAC address, as most hardware should, and if it is used by the kernel, nothing is done. Otherwise, a new MAC address is generated which is guaranteed to be the same on every boot for the given machine and the given device, but which is otherwise random. This feature depends on ID_NET_NAME_* properties to exist for the link. On hardware where these properties are not set, the generation of a persistent MAC address will fail.
	//
	// random
	// If the kernel is using a random MAC address, nothing is done. Otherwise, a new address is randomly generated each time the device appears, typically at boot. Either way, the random address will have the "unicast" and "locally administered" bits set.
	//
	// none
	// Keeps the MAC address assigned by the kernel. Or use the MAC address specified in MACAddress=.
	//
	// An empty string assignment is equivalent to setting "none".
	//
	// MACAddress=
	// The interface MAC address to use. For this setting to take effect, MACAddressPolicy= must either be unset, empty, or "none".
	//
	// NamePolicy=
	// An ordered, space-separated list of policies by which the interface name should be set. NamePolicy= may be disabled by specifying net.ifnames=0 on the kernel command line. Each of the policies may fail, and the first successful one is used. The name is not set directly, but is exported to udev as the property ID_NET_NAME, which is, by default, used by a udev(7), rule to set NAME. The available policies are:
	//
	// kernel
	// If the kernel claims that the name it has set for a device is predictable, then no renaming is performed.
	//
	// database
	// The name is set based on entries in the udev's Hardware Database with the key ID_NET_NAME_FROM_DATABASE.
	//
	// onboard
	// The name is set based on information given by the firmware for on-board devices, as exported by the udev property ID_NET_NAME_ONBOARD. See systemd.net-naming-scheme(7).
	//
	// slot
	// The name is set based on information given by the firmware for hot-plug devices, as exported by the udev property ID_NET_NAME_SLOT. See systemd.net-naming-scheme(7).
	//
	// path
	// The name is set based on the device's physical location, as exported by the udev property ID_NET_NAME_PATH. See systemd.net-naming-scheme(7).
	//
	// mac
	// The name is set based on the device's persistent MAC address, as exported by the udev property ID_NET_NAME_MAC. See systemd.net-naming-scheme(7).
	//
	// keep
	// If the device already had a name given by userspace (as part of creation of the device or a rename), keep it.
	MACAddressPolicy *string `systemd:",omitempty"`

	// The interface name to use. This option has lower precedence than NamePolicy=, so for this setting to take effect, NamePolicy= must either be unset, empty, disabled, or all policies configured there must fail. Also see the example below with "Name=dmz0".
	// Note that specifying a name that the kernel might use for another interface (for example "eth0") is dangerous because the name assignment done by udev will race with the assignment done by the kernel, and only one interface may use the name. Depending on the order of operations, either udev or the kernel will win, making the naming unpredictable. It is best to use some different prefix, for example "internal0"/"external0" or "lan0"/"lan1"/"lan3".
	Name *string `systemd:",omitempty"`

	// A space-separated list of policies by which the interface's alternative names should be set. Each of the policies may fail, and all successful policies are used. The available policies are "database", "onboard", "slot", "path", and "mac". If the kernel does not support the alternative names, then this setting will be ignored.
	AlternativeNamesPolicies []string `systemd:"AlternativeNamesPolicy,omitempty,wslist"`

	// The alternative interface name to use. This option can be specified multiple times. If the empty string is assigned to this option, the list is reset, and all prior assignments have no effect. If the kernel does not support the alternative names, then this setting will be ignored.
	AlternativeNames []string `systemd:"AlternativeName,omitempty"`

	// Specifies the device's number of transmit queues. An integer in the range 1…4096. When unset, the kernel's default will be used.
	TransmitQueues uint `systemd:",omitempty"`

	// Specifies the device's number of receive queues. An integer in the range 1…4096. When unset, the kernel's default will be used.
	ReceiveQueues uint `systemd:",omitempty"`

	// Specifies the transmit queue length of the device in number of packets. An unsigned integer in the range 0…4294967294. When unset, the kernel's default will be used.
	TransmitQueueLength *uint `systemd:",omitempty"`

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G are supported and are understood to the base of 1024.
	MTUBytes *string `systemd:",omitempty"`

	// The speed to set for the device, the value is rounded down to the nearest Mbps. The usual suffixes K, M, G are supported and are understood to the base of 1000.
	BitsPerSecond *string `systemd:",omitempty"`

	// The duplex mode to set for the device. The accepted values are half and full.
	Duplex *string `systemd:",omitempty"`

	// Takes a boolean. If set to yes, automatic negotiation of transmission parameters is enabled. Autonegotiation is a procedure by which two connected ethernet devices choose common transmission parameters, such as speed, duplex mode, and flow control. When unset, the kernel's default will be used.
	// Note that if autonegotiation is enabled, speed and duplex settings are read-only. If autonegotiation is disabled, speed and duplex settings are writable if the driver supports multiple link modes.
	AutoNegotiation *bool `systemd:",omitempty"`

	// The Wake-on-LAN policy to set for the device. Takes the special value "off" which disables Wake-on-LAN, or space separated list of the following words:
	// phy
	// Wake on PHY activity.
	//
	// unicast
	// Wake on unicast messages.
	//
	// multicast
	// Wake on multicast messages.
	//
	// broadcast
	// Wake on broadcast messages.
	//
	// arp
	// Wake on ARP.
	//
	// magic
	// Wake on receipt of a magic packet.
	//
	// secureon
	// Enable secureon(tm) password for MagicPacket(tm).
	//
	// Defaults to unset, and the device's default will be used. This setting can be specified multiple times. If an empty string is assigned, then the all previous assignments are cleared.
	WakeOnLan *string `systemd:",omitempty"`

	// The port option is used to select the device port. The supported values are:
	// tp
	// An Ethernet interface using Twisted-Pair cable as the medium.
	//
	// aui
	// Attachment Unit Interface (AUI). Normally used with hubs.
	//
	// bnc
	// An Ethernet interface using BNC connectors and co-axial cable.
	//
	// mii
	// An Ethernet interface using a Media Independent Interface (MII).
	//
	// fibre
	// An Ethernet interface using Optical Fibre as the medium.
	Port *string `systemd:",omitempty"`

	// This sets what speeds and duplex modes of operation are advertised for auto-negotiation. This implies "AutoNegotiation=yes". The supported values are:
	//
	// Table 1. Supported advertise values
	//
	// Advertise	Speed (Mbps)	Duplex Mode
	// 10baset-half	10	half
	// 10baset-full	10	full
	// 100baset-half	100	half
	// 100baset-full	100	full
	// 1000baset-half	1000	half
	// 1000baset-full	1000	full
	// 10000baset-full	10000	full
	// 2500basex-full	2500	full
	// 1000basekx-full	1000	full
	// 10000basekx4-full	10000	full
	// 10000basekr-full	10000	full
	// 10000baser-fec	10000	full
	// 20000basemld2-full	20000	full
	// 20000basekr2-full	20000	full
	//
	// By default this is unset, i.e. all possible modes will be advertised. This option may be specified more than once, in which case all specified speeds and modes are advertised. If the empty string is assigned to this option, the list is reset, and all prior assignments have no effect.
	Advertise *string `systemd:",omitempty"`

	// Takes a boolean. If set to true, hardware offload for checksumming of ingress network packets is enabled. When unset, the kernel's default will be used.
	ReceiveChecksumOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, hardware offload for checksumming of egress network packets is enabled. When unset, the kernel's default will be used.
	TransmitChecksumOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, TCP Segmentation Offload (TSO) is enabled. When unset, the kernel's default will be used.
	TCPSegmentationOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, TCP6 Segmentation Offload (tx-tcp6-segmentation) is enabled. When unset, the kernel's default will be used.
	TCP6SegmentationOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, Generic Segmentation Offload (GSO) is enabled. When unset, the kernel's default will be used.
	GenericSegmentationOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, Generic Receive Offload (GRO) is enabled. When unset, the kernel's default will be used.
	GenericReceiveOffload *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, hardware accelerated Generic Receive Offload (GRO) is enabled. When unset, the kernel's default will be used.
	GenericReceiveOffloadHardware *bool `systemd:",omitempty"`

	// Takes a boolean. If set to true, Large Receive Offload (LRO) is enabled. When unset, the kernel's default will be used.
	LargeReceiveOffload *bool `systemd:",omitempty"`

	// Specifies the number of receive, transmit, other, or combined channels, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	RxChannels       *string `systemd:",omitempty"`
	TxChannels       *string `systemd:",omitempty"`
	OtherChannels    *string `systemd:",omitempty"`
	CombinedChannels *string `systemd:",omitempty"`

	// Specifies the maximum number of pending packets in the NIC receive buffer, mini receive buffer, jumbo receive buffer, or transmit buffer, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	RxBufferSize      *string `systemd:",omitempty"`
	RxMiniBufferSize  *string `systemd:",omitempty"`
	RxJumboBufferSize *string `systemd:",omitempty"`
	TxBufferSize      *string `systemd:",omitempty"`

	// Takes a boolean. When set, enables receive flow control, also known as the ethernet receive PAUSE message (generate and send ethernet PAUSE frames). When unset, the kernel's default will be used.
	RxFlowControl *bool `systemd:",omitempty"`

	// Takes a boolean. When set, enables transmit flow control, also known as the ethernet transmit PAUSE message (respond to received ethernet PAUSE frames). When unset, the kernel's default will be used.
	TxFlowControl *bool `systemd:",omitempty"`

	// Takes a boolean. When set, auto negotiation enables the interface to exchange state advertisements with the connected peer so that the two devices can agree on the ethernet PAUSE configuration. When unset, the kernel's default will be used.
	AutoNegotiationFlowControl *bool `systemd:",omitempty"`

	// Specifies the maximum size of a Generic Segment Offload (GSO) packet the device should accept. The usual suffixes K, M, G are supported and are understood to the base of 1024. An unsigned integer in the range 1…65536. Defaults to unset.
	GenericSegmentOffloadMaxBytes *uint `systemd:",omitempty"`

	// Specifies the maximum number of Generic Segment Offload (GSO) segments the device should accept. An unsigned integer in the range 1…65535. Defaults to unset.
	GenericSegmentOffloadMaxSegments *uint `systemd:",omitempty"`

	// Boolean properties that, when set, enable/disable adaptive Rx/Tx coalescing if the hardware supports it. When unset, the kernel's default will be used.
	UseAdaptiveRxCoalesce *bool `systemd:",omitempty"`
	UseAdaptiveTxCoalesce *bool `systemd:",omitempty"`

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	RxCoalesceSec     *string `systemd:",omitempty"`
	RxCoalesceIrqSec  *string `systemd:",omitempty"`
	RxCoalesceLowSec  *string `systemd:",omitempty"`
	RxCoalesceHighSec *string `systemd:",omitempty"`
	TxCoalesceSec     *string `systemd:",omitempty"`
	TxCoalesceIrqSec  *string `systemd:",omitempty"`
	TxCoalesceLowSec  *string `systemd:",omitempty"`
	TxCoalesceHighSec *string `systemd:",omitempty"`

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated. The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	RxMaxCoalescedFrames     *string `systemd:",omitempty"`
	RxMaxCoalescedIrqFrames  *string `systemd:",omitempty"`
	RxMaxCoalescedLowFrames  *string `systemd:",omitempty"`
	RxMaxCoalescedHighFrames *string `systemd:",omitempty"`
	TxMaxCoalescedFrames     *string `systemd:",omitempty"`
	TxMaxCoalescedIrqFrames  *string `systemd:",omitempty"`
	TxMaxCoalescedLowFrames  *string `systemd:",omitempty"`
	TxMaxCoalescedHighFrames *string `systemd:",omitempty"`

	// These properties configure the low and high packet rate (expressed in packets per second) threshold respectively and are used to determine when the corresponding coalescing settings for low and high packet rates come into effect if adaptive Rx/Tx coalescing is enabled. If unset, the kernel's defaults will be used.
	CoalescePacketRateLow  *uint `systemd:",omitempty"`
	CoalescePacketRateHigh *uint `systemd:",omitempty"`

	// Configures how often to sample the packet rate used for adaptive Rx/Tx coalescing. This property cannot be zero. This lowest time granularity supported by this property is seconds. Partial seconds will be rounded up before being passed to the kernel. If unset, the kernel's default will be used.
	CoalescePacketRateSampleIntervalSec uint `systemd:",omitempty"`

	// How long to delay driver in-memory statistics block updates. If the driver does not have an in-memory statistic block, this property is ignored. This property cannot be zero. If unset, the kernel's default will be used.
	StatisticsBlockCoalesceSec uint `systemd:",omitempty"`
}
