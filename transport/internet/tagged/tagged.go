package tagged

import (
	"context"

	"github.com/hkobir1/xray_core/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
