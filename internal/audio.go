package internal

/*
#include "miniaudio.h"
#cgo LDFLAGS: -lm
*/
import "C"

import (
	_ "embed"
)

//go:embed audio.raw
var AudioRawData []byte

type CAudioBuffer = C.ma_audio_buffer
type CAudioBufferConfig = C.ma_audio_buffer_config

// var globalAudioContext = (*C.ma_context)(C.malloc(C.sizeof_ma_context))
var globalAudioContext = (*C.ma_context)(C.malloc(C.sizeof_ma_context))
