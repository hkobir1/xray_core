package encoding

import (
	"github.com/hkobir1/xray_core/common/net"
	"github.com/hkobir1/xray_core/common/protocol"
)

//go:generate go run github.com/hkobir1/xray_core/common/errors/errorgen

const (
	Version = byte(1)
)

var addrParser = protocol.NewAddressParser(
	protocol.AddressFamilyByte(byte(protocol.AddressTypeIPv4), net.AddressFamilyIPv4),
	protocol.AddressFamilyByte(byte(protocol.AddressTypeDomain), net.AddressFamilyDomain),
	protocol.AddressFamilyByte(byte(protocol.AddressTypeIPv6), net.AddressFamilyIPv6),
	protocol.PortThenAddress(),
)
