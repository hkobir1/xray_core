package command_test

import (
	"context"
	"testing"

	"github.com/hkobir1/xray_core/app/dispatcher"
	"github.com/hkobir1/xray_core/app/log"
	. "github.com/hkobir1/xray_core/app/log/command"
	"github.com/hkobir1/xray_core/app/proxyman"
	_ "github.com/hkobir1/xray_core/app/proxyman/inbound"
	_ "github.com/hkobir1/xray_core/app/proxyman/outbound"
	"github.com/hkobir1/xray_core/common"
	"github.com/hkobir1/xray_core/common/serial"
	"github.com/hkobir1/xray_core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
