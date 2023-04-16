package audio

import (
	"errors"
	"github.com/mearaj/audio/internal"
	"runtime"
)

type RawRecorderState = internal.RawRecorderState

const (
	RawRecorderStateUninitialized = internal.RawRecorderStateIdle
	RawRecorderStateRecording     = internal.RawRecorderStateRecording
	RawRecorderStatePaused        = internal.RawRecorderStatePaused
	RawRecorderStateStopped       = internal.RawRecorderStateStopped
)

var microphonePermissionGranted bool

var androidView uintptr
var sdkVersion int32

// SetView view is a JNI global reference to the android.view.View
// required if android recording permission is needed by the caller of this app
func SetView(view uintptr) {
	androidView = view
}

func init() {
	go func() {
		_ = checkAndSetRecorderPermission()
	}()
}

//type RawRecorder interface {
//	Record() error
//	Pause() error
//	Stop() error
//	State() RawRecorderState
//	Bytes() []byte
//}

type RawRecorder = internal.RawRecorder

func NewRawRecorder(format internal.MaFormat, numberOfChannels internal.MaUint32) (*RawRecorder, error) {
	if !microphonePermissionGranted && runtime.GOOS == "android" {
		_ = requestPermission(androidView)
	}
	if !microphonePermissionGranted {
		return nil, errors.New("microphone permission required")
	}
	return internal.NewRawRecorder(format, numberOfChannels)
}
