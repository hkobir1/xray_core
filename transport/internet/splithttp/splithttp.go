package splithttp

import (
	"context"

	"github.com/hkobir1/xray_core/common"
	"github.com/hkobir1/xray_core/common/errors"
)

//go:generate go run github.com/hkobir1/xray_core/common/errors/errorgen

const protocolName = "splithttp"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, errors.New("splithttp is a transport protocol.")
	}))
}
