package lzma

// #cgo CFLAGS: -O3
// #include "src/Lzma86.h"
import "C"

// FilterMode type of available compression filters
type FilterMode C.int

const (
	// FilterNo - no use filter
	FilterNo FilterMode = C.SZ_FILTER_NO
	// FilterYes - use x86 filter
	FilterYes FilterMode = C.SZ_FILTER_YES
	// FilterAuto - it tries both alternatives to select best
	FilterAuto FilterMode = C.SZ_FILTER_AUTO
)

// CompressionLevel - compression level: 0 <= level <= 9, the default value for "level" is 5.
type CompressionLevel int

// Compression level 0..9
const (
	Level0 CompressionLevel = iota
	Level1
	Level2
	Level3
	Level4
	Level5
	Level6
	Level7
	Level8
	Level9
)

// EncodeParameters describe parameters of compression algorithm lzma
type EncodeParameters struct {
	level      C.int
	dictSize   C.UInt32
	filterMode C.int
}

// NewEncodeParameters create new EncodeParameters with default values
func NewEncodeParameters() *EncodeParameters {
	return &EncodeParameters{
		level:      C.int(Level5),
		dictSize:   C.UInt32(1 << 23),
		filterMode: C.int(FilterAuto)}
}

// CompressionLevel set compression level
func (p *EncodeParameters) CompressionLevel(level CompressionLevel) {
	p.level = C.int(level)
}

// DictionarySize set dictionary size
func (p *EncodeParameters) DictionarySize(dictSize uint32) {
	p.dictSize = C.UInt32(dictSize)
}

// FilterMode set filter mode
func (p *EncodeParameters) FilterMode(filterMode FilterMode) {
	p.filterMode = C.int(filterMode)
}
