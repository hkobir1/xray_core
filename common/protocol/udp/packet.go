package udp

import (
	"github.com/hkobir1/xray_core/common/buf"
	"github.com/hkobir1/xray_core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
