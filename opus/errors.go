package opus

import "fmt"

/*
#include <opus.h>
*/
import "C"

// Error is a typy for libopus errors
type Error int

var _ error = Error(0)

// Libopus errors
const (
	ErrOK             = Error(C.OPUS_OK)
	ErrBadArg         = Error(C.OPUS_BAD_ARG)
	ErrBufferTooSmall = Error(C.OPUS_BUFFER_TOO_SMALL)
	ErrInternalError  = Error(C.OPUS_INTERNAL_ERROR)
	ErrInvalidPacket  = Error(C.OPUS_INVALID_PACKET)
	ErrUnimplemented  = Error(C.OPUS_UNIMPLEMENTED)
	ErrInvalidState   = Error(C.OPUS_INVALID_STATE)
	ErrAllocFail      = Error(C.OPUS_ALLOC_FAIL)
)

// Error return string value for libopus errors.
func (e Error) Error() string {
	return fmt.Sprintf("opus: %s", C.GoString(C.opus_strerror(C.int(e))))
}
