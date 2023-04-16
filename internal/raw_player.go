package internal

/*
#include <stdlib.h>
#include <string.h>
#include "miniaudio.h"
extern ma_device_data_proc rawPlayerDataCallback;
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

type RawPlayerState int

const (
	RawPlayerStateIdle = RawPlayerState(iota)
	RawPlayerStatePlaying
	RawPlayerStateStopped
	RawPlayerStatePaused
)

type RawPlayer struct {
	playerState      RawPlayerState
	bytesPerFrame    uint32
	device           *MaDevice
	deviceConfig     *MaDeviceConfig
	reader           *utils.RawReader
	playerStateMutex sync.RWMutex
}

func NewRawPlayer(data []byte, format MaFormat, numberOfChannels MaUint32) (*RawPlayer, error) {
	rawPlayer, err := newDefaultMaRawPlayerDevice(data, format, numberOfChannels)
	if err != nil && !errors.Is(err, MaResultSuccess) {
		return nil, err
	}
	return rawPlayer, nil
}

func (r *RawPlayer) Play() (err error) {
	playerState := r.State()
	if playerState == RawPlayerStatePlaying {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		r.setState(RawPlayerStatePlaying)
	case MaDeviceStateStopped, MaDeviceStateStopping:
		err = MaResult(C.ma_device_start(r.device))
		if !errors.Is(err, MaResultSuccess) && !errors.Is(err, MaResultInvalidOperation) {
			return err
		}
		r.setState(RawPlayerStatePlaying)
	case MaDeviceStateUninitialized:
		// Todo
	}
	return nil
}

func (r *RawPlayer) Pause() (err error) {
	state := r.State()
	if state == RawPlayerStatePaused {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		r.setState(RawPlayerStatePaused)
		return nil
	}
	return MaResultInvalidOperation
}

func (r *RawPlayer) getDeviceState() MaDeviceState {
	return (MaDeviceState)(C.ma_device_get_state(r.device))
}

func (r *RawPlayer) setState(state RawPlayerState) {
	r.playerStateMutex.Lock()
	r.playerState = state
	r.playerStateMutex.Unlock()
}

func (r *RawPlayer) Stop() (err error) {
	defer func() {
		_, _ = r.reader.Seek(0, io.SeekStart)
	}()
	state := r.State()
	if state == RawPlayerStateStopped {
		return nil
	}
	deviceState := r.getDeviceState()
	switch deviceState {
	case MaDeviceStateStarted, MaDeviceStateStarting:
		err = MaResult(C.ma_device_stop(r.device))
		if !errors.Is(err, MaResultSuccess) && !errors.Is(err, MaResultInvalidOperation) {
			return err
		}
		r.setState(RawPlayerStateStopped)
		return nil
	case MaDeviceStateStopped, MaDeviceStateStopping:
		r.setState(RawPlayerStateStopped)
		return nil
	case MaDeviceStateUninitialized:
		// Todo:
	}
	return nil
}

func (r *RawPlayer) State() RawPlayerState {
	r.playerStateMutex.RLock()
	state := r.playerState
	r.playerStateMutex.RUnlock()
	return state
}

var rawPlayers = utils.NewMap[*MaDevice, *RawPlayer]()

// RawPlayerDataCallback Do not call these APIs directly from inside -> Ref: https://miniaud.io/docs/manual/index.html#Introduction
//
//	ma_device_init()
//	ma_device_init_ex()
//	ma_device_uninit()
//	ma_device_start()
//	ma_device_stop()
//
// The alternate option would be using go routine
//
//export RawPlayerDataCallback
func RawPlayerDataCallback(device *MaDevice, output unsafe.Pointer, pointer unsafe.Pointer, frameCount uint32) {
	if player, ok := rawPlayers.Get(device); ok {
		// if player is not paused
		if player.State() != RawPlayerStatePaused {
			bytesSize := frameCount * player.bytesPerFrame
			bytes := ([]byte)(C.GoBytes(output, C.int(bytesSize)))
			count, err := player.reader.Read(bytes)
			var shouldStop bool
			if err != nil {
				if errors.Is(err, io.EOF) {
					shouldStop = true
					err = nil
				}
				if err != nil {
					alog.Logger().Errorln(err)
				}
			}
			if count > 0 {
				cBytes := C.CBytes(bytes)
				C.memcpy(output, cBytes, (C.size_t)(bytesSize))
				C.free(cBytes)
			}
			shouldStop = shouldStop || count == 0 || uint32(count) < bytesSize
			if shouldStop {
				go func() {
					_ = player.Stop()
				}()
			}
		}
	}
}

func GetBytesPerSample(format MaFormat) uint32 {
	switch format {
	case MaFormatS16:
		return 2
	case MaFormatS24:
		return 3
	case MaFormatF32, MaFormatS32:
		return 4
	case MaFormatU8, MaFormatUnknown:
		fallthrough
	default:
		return 1
	}
}

func GetBytesPerFrame(format MaFormat, channels uint32) uint32 {
	return GetBytesPerSample(format) * channels
}

func newDefaultMaRawPlayerDevice(data []byte, format MaFormat, channels MaUint32) (*RawPlayer, error) {
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
	*deviceConfig = C.ma_device_config_init(MaDeviceTypePlayback)
	deviceConfig.playback.format = format
	deviceConfig.playback.channels = channels
	deviceConfig.sampleRate = 0
	deviceConfig.dataCallback = C.rawPlayerDataCallback
	err = MaResult(C.ma_device_init(nil, deviceConfig, device))
	if !errors.Is(err, MaResultSuccess) {
		return nil, err
	}
	p := &RawPlayer{
		bytesPerFrame: bytesPerFrame,
		device:        device,
		deviceConfig:  deviceConfig,
		reader:        utils.NewRawReader(data),
	}
	rawPlayers.Set(device, p)
	return p, nil
}
