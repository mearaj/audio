package internal

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestQueryCacheDevices(t *testing.T) {
	ch := make(chan struct{})
	now := time.Now()
	go func() {
		for i := 0; i < 1; i++ {
			devices := QueryDevicesInfoCache()
			for i, device := range devices.PlaybackDevices() {
				fmt.Printf("%d. %s, %s\n", i, device.Name, string(device.ID))
			}
			for i, device := range devices.CaptureDevices() {
				fmt.Printf("%d. %s, %s\n", i, device.Name, string(device.ID))
			}
		}
		ch <- struct{}{}
	}()
	<-ch
	then := time.Now()
	fmt.Println(then.Sub(now).String())
}

func TestNewDefaultPlaybackDevice(t *testing.T) {
	_, err := newDefaultMaRawPlayerDevice(AudioRawData, MaFormatUnknown, 0)
	time.Sleep(time.Second * 12)
	if err != nil && !errors.Is(err, MaResultSuccess) {
		t.Error(err)
	}
}

func TestNewDefaultCaptureDevice(t *testing.T) {
	d, err := newDefaultMaRawRecorderDevice(MaFormatUnknown, 0)
	if err != nil && !errors.Is(err, MaResultSuccess) {
		t.Error(err)
	}
	err = d.Record()
	time.Sleep(time.Second * 4)
	err = d.Stop()
	if err != nil && !errors.Is(err, MaResultSuccess) {
		t.Error(err)
	}
	e, err := newDefaultMaRawPlayerDevice(d.Bytes(), MaFormatUnknown, 0)
	err = e.Play()
	if err != nil && !errors.Is(err, MaResultSuccess) {
		t.Error(err)
	}
	time.Sleep(time.Second * 4)
	err = e.Stop()
	if err != nil && !errors.Is(err, MaResultSuccess) {
		t.Error(err)
	}
}
