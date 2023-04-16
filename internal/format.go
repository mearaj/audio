package internal

/*
#include "miniaudio.h"
#cgo LDFLAGS: -lm
*/
import "C"

type MaFormat = C.ma_format

const (
	MaFormatUnknown = MaFormat(C.ma_format_unknown)
	MaFormatU8      = MaFormat(C.ma_format_u8)
	MaFormatS16     = MaFormat(C.ma_format_s16)
	MaFormatS24     = MaFormat(C.ma_format_s24)
	MaFormatS32     = MaFormat(C.ma_format_s32)
	MaFormatF32     = MaFormat(C.ma_format_f32)
	MaFormatCount   = MaFormat(C.ma_format_count)
)
