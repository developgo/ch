// Package compress implements compression support.
package compress

import (
	"encoding/binary"
	"fmt"

	"github.com/go-faster/city"
)

//go:generate go run github.com/dmarkham/enumer -transform snake_upper -type Method -output method_enum.go

// Method is compression codec.
type Method byte

// Possible compression methods.
const (
	None Method = 0x02
	LZ4  Method = 0x82
	ZSTD Method = 0x90
)

const (
	checksumSize       = 16
	compressHeaderSize = 1 + 4 + 4
	headerSize         = checksumSize + compressHeaderSize
	maxBlockSize       = 1024 * 1024 * 1   // 1MB
	maxDataSize        = 1024 * 1024 * 128 // 128MB

	hRawSize  = 17
	hDataSize = 21
	hMethod   = 16
)

var bin = binary.LittleEndian

// CorruptedDataErr means that provided hash mismatch with calculated.
type CorruptedDataErr struct {
	Actual    city.U128
	Reference city.U128
	RawSize   int
	DataSize  int
}

func (c *CorruptedDataErr) Error() string {
	return fmt.Sprintf("corrupted data: %s (actual), %s (reference), compressed size: %d, data size: %d",
		FormatU128(c.Actual), FormatU128(c.Reference), c.RawSize, c.DataSize,
	)
}
