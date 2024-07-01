package tcp

import (
	"github.com/hkobir1/xray_core/common"
	"github.com/hkobir1/xray_core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
