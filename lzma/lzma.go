package lzma

// #cgo CFLAGS: -O3
// #define _7ZIP_ST
// #include "src/Lzma86.h"
// #include "src/Lzma86Dec.c"
// #include "src/Lzma86Enc.c"
// #include "src/Alloc.c"
// #include "src/Bra.c"
// #include "src/Bra86.c"
// #include "src/CpuArch.c"
// #include "src/LzmaDec.c"
// #include "src/LzmaEnc.c"
// #include "src/LzFind.c"
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Get a pointer to the first byte of a slice
func toBytePointer(input *[]byte) *C.Byte {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(input))
	return (*C.Byte)(unsafe.Pointer(hdr.Data))
}

// Lzma86Encode compress lzma data with X86 filter
func Lzma86Encode(input []byte, output []byte, params *EncodeParameters) (int, error) {
	inp := toBytePointer(&input)
	outp := toBytePointer(&output)
	inLen := C.size_t((len(input)))
	outLen := C.size_t((len(output)))

	res := int(C.Lzma86_Encode(outp, &outLen, inp, inLen, params.level, params.dictSize, params.filterMode))
	if res != C.SZ_OK {
		return 0, fmt.Errorf("Lzma86_Encode failed (%d)", res)
	}

	return int(outLen), nil
}

// Lzma86GetUnpackSize return size of uncompressed stream
func Lzma86GetUnpackSize(input []byte) (int, error) {
	inp := toBytePointer(&input)
	inLen := C.SizeT((len(input)))

	var unpackSize C.UInt64

	res := int(C.Lzma86_GetUnpackSize(inp, inLen, &unpackSize))
	if res != C.SZ_OK {
		return 0, fmt.Errorf("Lzma86_GetUnpackSize failed (%d)", res)
	}

	return int(unpackSize), nil
}

// Lzma86Decode uncompress lzma data with X86 filter
func Lzma86Decode(input []byte, output []byte) (int, error) {
	inp := toBytePointer(&input)
	outp := toBytePointer(&output)
	inLen := C.SizeT((len(input)))
	outLen := C.SizeT((len(output)))

	res := int(C.Lzma86_Decode(outp, &outLen, inp, &inLen))
	if res != C.SZ_OK {
		return 0, fmt.Errorf("Lzma86_Decode failed (%d)", res)
	}

	return int(outLen), nil
}
