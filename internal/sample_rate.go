package internal

/*
#include "miniaudio.h"
#cgo LDFLAGS: -lm
*/
import "C"

type MaStandardSampleRate = C.ma_standard_sample_rate

const (
	MaStandardSampleRate48000  = MaStandardSampleRate(C.ma_standard_sample_rate_48000)
	MaStandardSampleRate44100  = MaStandardSampleRate(C.ma_standard_sample_rate_44100)
	MaStandardSampleRate32000  = MaStandardSampleRate(C.ma_standard_sample_rate_32000)
	MaStandardSampleRate24000  = MaStandardSampleRate(C.ma_standard_sample_rate_24000)
	MaStandardSampleRate22050  = MaStandardSampleRate(C.ma_standard_sample_rate_22050)
	MaStandardSampleRate88200  = MaStandardSampleRate(C.ma_standard_sample_rate_88200)
	MaStandardSampleRate96000  = MaStandardSampleRate(C.ma_standard_sample_rate_96000)
	MaStandardSampleRate176400 = MaStandardSampleRate(C.ma_standard_sample_rate_176400)
	MaStandardSampleRate192000 = MaStandardSampleRate(C.ma_standard_sample_rate_192000)
	MaStandardSampleRate16000  = MaStandardSampleRate(C.ma_standard_sample_rate_16000)
	MaStandardSampleRate11025  = MaStandardSampleRate(C.ma_standard_sample_rate_11025)
	MaStandardSampleRate8000   = MaStandardSampleRate(C.ma_standard_sample_rate_8000)
	MaStandardSampleRate352800 = MaStandardSampleRate(C.ma_standard_sample_rate_352800)
	MaStandardSampleRate384000 = MaStandardSampleRate(C.ma_standard_sample_rate_384000)
	MaStandardSampleRateMin    = MaStandardSampleRate(C.ma_standard_sample_rate_min)
	MaStandardSampleRateMax    = MaStandardSampleRate(C.ma_standard_sample_rate_max)
	MaStandardSampleRateCount  = MaStandardSampleRate(C.ma_standard_sample_rate_count)
)
