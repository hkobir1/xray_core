package blackhole_test

import (
	"context"
	"testing"

	"github.com/hkobir1/xray_core/common"
	"github.com/hkobir1/xray_core/common/buf"
	"github.com/hkobir1/xray_core/common/serial"
	"github.com/hkobir1/xray_core/common/session"
	"github.com/hkobir1/xray_core/proxy/blackhole"
	"github.com/hkobir1/xray_core/transport"
	"github.com/hkobir1/xray_core/transport/pipe"
)

func TestBlackholeHTTPResponse(t *testing.T) {
	ctx := session.ContextWithOutbounds(context.Background(), []*session.Outbound{{}})
	handler, err := blackhole.New(ctx, &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	var mb buf.MultiBuffer
	var rerr error
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		rerr = e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(ctx, &link, nil))
	common.Must(rerr)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
