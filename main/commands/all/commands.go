package all

import (
	"github.com/hkobir1/xray_core/main/commands/all/api"
	"github.com/hkobir1/xray_core/main/commands/all/tls"
	"github.com/hkobir1/xray_core/main/commands/base"
)

// go:generate go run github.com/hkobir1/xray_core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
