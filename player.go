package audio

import (
	"github.com/mearaj/audio/internal"
)

type RawPlayerState = internal.RawPlayerState

const (
	RawPlayerStateIdle    = internal.RawPlayerStateIdle
	RawPlayerStateStopped = internal.RawPlayerStateStopped
	RawPlayerStatePlaying = internal.RawPlayerStatePlaying
	RawPlayerStatePaused  = internal.RawPlayerStatePaused
)

//type RawPlayer interface {
//	Play() error
//	Stop() error
//	Pause() error
//	State() RawPlayerState
//}

type RawPlayer = internal.RawPlayer

func NewRawPlayer(data []byte, format internal.MaFormat, numberOfChannels internal.MaUint32) (*RawPlayer, error) {
	return internal.NewRawPlayer(data, format, numberOfChannels)
}
