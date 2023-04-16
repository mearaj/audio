package internal

/*
#include "miniaudio.h"
#cgo LDFLAGS: -lm
//#cgo linux LDFLAGS: -ldl -lpthread -lm
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type MaDeviceConfig = C.ma_device_config
type MaDevice = C.ma_device
type MaDecoder = C.ma_decoder
type MaDeviceState = C.ma_device_state
type MaDeviceInfo = C.ma_device_info
type MaUint32 = C.ma_uint32
type MaUint64 = C.ma_uint64
type MaDeviceType = C.ma_device_type
type CULongLong = C.ulonglong
type CULong = C.ulong

type DeviceID []byte

const (
	MaDeviceTypePlayback = MaDeviceType(C.ma_device_type_playback)
	MaDeviceTypeCapture  = MaDeviceType(C.ma_device_type_capture)
	MaDeviceTypeDuplex   = MaDeviceType(C.ma_device_type_duplex)
	MaDeviceTypeLoopback = MaDeviceType(C.ma_device_type_loopback)
)

const (
	MaDeviceStateUninitialized = MaDeviceState(C.ma_device_state_uninitialized)
	MaDeviceStateStopped       = MaDeviceState(C.ma_device_state_stopped)
	MaDeviceStateStarted       = MaDeviceState(C.ma_device_state_started)
	MaDeviceStateStarting      = MaDeviceState(C.ma_device_state_starting)
	MaDeviceStateStopping      = MaDeviceState(C.ma_device_state_stopping)
)

var maPlaybackDeviceInfos *MaDeviceInfo
var maPlaybackDeviceCount MaUint32
var maCaptureDeviceInfos *MaDeviceInfo
var maCaptureDeviceCount MaUint32
var cachedDevicesInfos DevicesInfo

type DevicesInfo struct {
	playbackDevices []DeviceInfo
	captureDevices  []DeviceInfo
}

var _ = initializeGlobalContext()
var _, _ = QueryDevicesInfo()

func (d *DevicesInfo) GetAllDevicesInfo() (devices []DeviceInfo) {
	return append(append(devices, d.playbackDevices...), d.captureDevices...)
}

func (d *DevicesInfo) PlaybackDevices() []DeviceInfo {
	return d.playbackDevices
}

func (d *DevicesInfo) CaptureDevices() []DeviceInfo {
	return d.captureDevices
}

func QueryDevicesInfoCache() DevicesInfo {
	return cachedDevicesInfos
}

type DeviceInfo struct {
	Name       string
	ID         DeviceID
	deviceType MaDeviceType
}

func initializeGlobalContext() error {
	if err := MaResult(C.ma_context_init(nil, 0, nil, globalAudioContext)); !errors.Is(err, MaResultSuccess) {
		return err
	}
	return nil
}

// QueryDevicesInfo only call this to re-fetch devices info from OS, else
// cached Data is available from QueryDevicesInfoCache which is much faster
func QueryDevicesInfo() (devicesInfo DevicesInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
		cachedDevicesInfos = devicesInfo
	}()
	err = MaResult(C.ma_context_get_devices(globalAudioContext, &maPlaybackDeviceInfos, &maPlaybackDeviceCount, &maCaptureDeviceInfos, &maCaptureDeviceCount))
	if !errors.Is(err, MaResultSuccess) {
		return devicesInfo, err
	}
	playbackDevices := make([]DeviceInfo, maPlaybackDeviceCount)
	playbackDevicesInfos := unsafe.Slice(maPlaybackDeviceInfos, int(maPlaybackDeviceCount))
	for i, d := range playbackDevicesInfos {
		playbackDevices[i] = DeviceInfo{
			Name:       string(C.GoString(&d.name[0])),
			ID:         d.id[:],
			deviceType: MaDeviceTypePlayback,
		}
	}
	devicesInfo.playbackDevices = playbackDevices
	captureDevices := make([]DeviceInfo, maCaptureDeviceCount)
	captureDevicesInfos := unsafe.Slice(maCaptureDeviceInfos, int(maCaptureDeviceCount))
	for i, d := range captureDevicesInfos {
		captureDevices[i] = DeviceInfo{
			Name:       string(C.GoString(&d.name[0])),
			ID:         d.id[:],
			deviceType: MaDeviceTypeCapture,
		}
	}
	devicesInfo.captureDevices = captureDevices
	return devicesInfo, nil
}
