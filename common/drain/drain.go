package drain

import "io"

//go:generate go run github.com/hkobir1/xray_core/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
