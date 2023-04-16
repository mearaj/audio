package utils

import (
	"encoding/binary"
	"errors"
	"io"
	"sync"
)

var ErrInvalidOffset = errors.New("invalid offset")
var ErrInvalidWhence = errors.New("invalid whence")

type RawWriter struct {
	bytesMutex sync.RWMutex
	posMutex   sync.RWMutex
	pos        int
	bytes      []byte
}

func NewRawWriter() *RawWriter {
	return NewVirtualFileWriter()
}

func NewVirtualFileWriter() *RawWriter {
	vf := &RawWriter{bytes: []byte{}}
	return vf
}

func (vf *RawWriter) Write(b []byte) (int, error) {
	pos := vf.Pos()
	n, err := vf.writeAt(b, int64(pos))
	pos += n
	vf.setPos(pos)
	return n, err
}

func (vf *RawWriter) WriteAt(b []byte, offset int64) (int, error) {
	return vf.writeAt(b, offset)
}

func (vf *RawWriter) Bytes() []byte {
	vf.bytesMutex.RLock()
	defer vf.bytesMutex.RUnlock()
	return vf.bytes
}

func (vf *RawWriter) WriteFloat(p []float32) (int, error) {
	return len(p), binary.Write(vf, binary.LittleEndian, p)
}

// Seek behavior is similar to io.Seeker,
// allowed whence value is io.SeekStart, io.SeekCurrent and io.SeekEnd
func (vf *RawWriter) Seek(offset int64, whence int) (int64, error) {
	var absolute int64
	switch whence {
	case io.SeekStart:
		absolute = offset
	case io.SeekCurrent:
		absolute = int64(vf.Pos()) + offset
	case io.SeekEnd:
		absolute = int64(len(vf.Bytes())) + offset
	default:
		return 0, ErrInvalidWhence
	}
	if absolute < 0 {
		return 0, ErrInvalidOffset
	}
	vf.setPos(int(absolute))
	return absolute, nil
}

func (vf *RawWriter) Pos() int {
	vf.posMutex.RLock()
	defer vf.posMutex.RUnlock()
	return vf.pos
}

func (vf *RawWriter) writeAt(b []byte, offset int64) (int, error) {
	if offset < 0 {
		return 0, ErrInvalidOffset
	}
	if offset > int64(len(vf.Bytes())) {
		err := vf.resize(offset)
		if err != nil {
			return 0, err
		}
	}
	n := copy(vf.Bytes()[offset:], b)
	bytes := append(vf.Bytes(), b[n:]...)
	vf.setBytes(bytes)
	return len(b), nil
}

func (vf *RawWriter) setPos(pos int) {
	vf.posMutex.Lock()
	defer vf.posMutex.Unlock()
	vf.pos = pos
}

func (vf *RawWriter) setBytes(b []byte) {
	vf.bytesMutex.Lock()
	defer vf.bytesMutex.Unlock()
	vf.bytes = b
}

// resize n should be the absolute value with respect to length of bytes
func (vf *RawWriter) resize(n int64) error {
	switch {
	case n < 0:
		return ErrInvalidOffset
	case n <= int64(len(vf.Bytes())):
		vf.setBytes(vf.Bytes()[:n])
		return nil
	default:
		bytes := append(vf.Bytes(), make([]byte, int(n)-len(vf.Bytes()))...)
		vf.setBytes(bytes)
		return nil
	}
}
