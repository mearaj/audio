package internal

/*
#include <stdlib.h>
#include <string.h>
#include "miniaudio.h"
extern ma_device_data_proc rawRecorderDataCallback;
#cgo LDFLAGS: -lm
*/
import "C"

import (
	"errors"
	"github.com/mearaj/audio/utils"
	"github.com/mearaj/protonet/alog"
	"io"
	"sync"
	"unsafe"
)

type RawRecorderState int

const (
	RawRecorderStateIdle = RawRecorderState(iota)
	RawRecorderStateRecording
	RawRecorderStatePaused
	RawRecorderStateStopped
)

type RawRecorder struct {
	recorderState      RawRecorderState
	bytesPerFrame      uint32
	device             *MaDevice
	deviceConfig       *MaDeviceConfig
	writer             *utils.RawWriter
	recorderStateMutex sync.RWMutex
}

func NewRawRecorder(format MaFormat, numberOfChannels MaUint32) (*RawRecorder, error) {
	rawRecorder, err := newDefaultMaRawRecorderDevice(format, numberOfChannels)
	if err != nil && !errors.Is(err, MaResultSuccess) {
		return nil, err
	}
	return rawRecorder, nil
}

func (r *RawRecorder) Record() (err error) {
	recorderState := r.State()
	if recorderState == RawRecorderStateRecording {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		r.setState(RawRecorderStateRecording)
	case MaDeviceStateStopped, MaDeviceStateStopping:
		err = MaResult(C.ma_device_start(r.device))
		if !errors.Is(err, MaResultSuccess) && !errors.Is(err, MaResultInvalidOperation) {
			return err
		}
		r.setState(RawRecorderStateRecording)
	case MaDeviceStateUninitialized:
		// Todo
	}
	return nil
}

func (r *RawRecorder) Pause() (err error) {
	state := r.State()
	if state == RawRecorderStatePaused {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		r.setState(RawRecorderStatePaused)
		return nil
	}
	return MaResultInvalidOperation
}

func (r *RawRecorder) getDeviceState() MaDeviceState {
	return (MaDeviceState)(C.ma_device_get_state(r.device))
}

func (r *RawRecorder) setState(state RawRecorderState) {
	r.recorderStateMutex.Lock()
	r.recorderState = state
	r.recorderStateMutex.Unlock()
}

func (r *RawRecorder) Stop() (err error) {
	defer func() {
		_, _ = r.writer.Seek(0, io.SeekStart)
	}()
	state := r.State()
	if state == RawRecorderStateStopped {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		err = MaResult(C.ma_device_stop(r.device))
		if !errors.Is(err, MaResultSuccess) && !errors.Is(err, MaResultInvalidOperation) {
			return err
		}
		r.setState(RawRecorderStateStopped)
		return nil
	case MaDeviceStateStopped, MaDeviceStateStopping:
		r.setState(RawRecorderStateStopped)
		return nil
	case MaDeviceStateUninitialized:
		// Todo:
	}
	return nil
}

func (r *RawRecorder) State() RawRecorderState {
	r.recorderStateMutex.RLock()
	state := r.recorderState
	r.recorderStateMutex.RUnlock()
	return state
}
func (r *RawRecorder) Bytes() []byte {
	return r.writer.Bytes()
}

var rawRecorders = utils.NewMap[*MaDevice, *RawRecorder]()

//export RawRecorderDataCallback
func RawRecorderDataCallback(device *MaDevice, output unsafe.Pointer, input unsafe.Pointer, frameCount uint32) {
	if recorder, ok := rawRecorders.Get(device); ok {
		// if recorder is not paused
		//if recorder.State() != RawRecorderStatePaused {
		bytesSize := frameCount * recorder.bytesPerFrame
		bytes := ([]byte)(C.GoBytes(input, C.int(bytesSize)))
		_, err := recorder.writer.Write(bytes)
		if err != nil {
			alog.Logger().Errorln(err)
		}
		//}
	}
}

func newDefaultMaRawRecorderDevice(format MaFormat, channels MaUint32) (*RawRecorder, error) {
	var err error
	if channels == 0 {
		channels = 1
	}
	if format == MaFormatUnknown {
		format = MaFormatS16
	}
	var bytesPerFrame = GetBytesPerFrame(format, uint32(channels))
	var deviceConfig *MaDeviceConfig = (*MaDeviceConfig)(C.malloc(C.sizeof_ma_device_config))
	var device *MaDevice = (*MaDevice)(C.malloc(C.sizeof_ma_device))
	*deviceConfig = C.ma_device_config_init(MaDeviceTypeCapture)
	deviceConfig.capture.format = format
	deviceConfig.capture.channels = channels
	deviceConfig.sampleRate = 0
	deviceConfig.dataCallback = C.rawRecorderDataCallback
	err = MaResult(C.ma_device_init(nil, deviceConfig, (*C.ma_device)(device)))
	if !errors.Is(err, MaResultSuccess) {
		return nil, err
	}
	p := &RawRecorder{
		bytesPerFrame: bytesPerFrame,
		device:        device,
		deviceConfig:  deviceConfig,
		writer:        utils.NewRawWriter(),
	}
	rawRecorders.Set(device, p)
	return p, nil
}
