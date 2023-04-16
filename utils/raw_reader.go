package utils

import (
	"encoding/binary"
	"io"
	"sync"
)

type RawReader struct {
	bytes      []byte
	pos        int
	bytesMutex sync.RWMutex
	posMutex   sync.RWMutex
}

func NewRawReader(b []byte) *RawReader {
	return NewVirtualFileReader(b)
}

func NewVirtualFileReader(b []byte) *RawReader {
	vf := &RawReader{bytes: b}
	return vf
}

// Read satisfies io.Reader interface
func (r *RawReader) Read(b []byte) (int, error) {
	pos := r.Pos()
	n, err := r.readAt(b, int64(pos))
	pos += n
	r.setPos(pos)
	return n, err
}

// ReadAt satisfies io.ReaderAt
func (r *RawReader) ReadAt(b []byte, offset int64) (int, error) {
	return r.readAt(b, offset)
}

// Seek behavior is similar to io.Seeker,
// allowed whence value is io.SeekStart, io.SeekCurrent and io.SeekEnd
func (r *RawReader) Seek(offset int64, whence int) (int64, error) {
	var absolute int64
	switch whence {
	case io.SeekStart:
		absolute = offset
	case io.SeekCurrent:
		absolute = int64(r.Pos()) + offset
	case io.SeekEnd:
		absolute = int64(len(r.Bytes())) + offset
	default:
		return 0, ErrInvalidWhence
	}
	if absolute < 0 {
		return 0, ErrInvalidOffset
	}
	r.setPos(int(absolute))
	return absolute, nil
}

func (r *RawReader) Pos() int {
	r.posMutex.RLock()
	defer r.posMutex.RUnlock()
	return r.pos
}

func (r *RawReader) Bytes() []byte {
	r.bytesMutex.RLock()
	defer r.bytesMutex.RUnlock()
	return r.bytes
}

func (r *RawReader) ReadFloat(out []float32, order binary.ByteOrder) (int, error) {
	return len(out), binary.Read(r, order, out)
}

func (r *RawReader) setPos(pos int) {
	r.posMutex.Lock()
	defer r.posMutex.Unlock()
	r.pos = pos
}

func (r *RawReader) readAt(b []byte, offset int64) (int, error) {
	if offset < 0 {
		return 0, ErrInvalidOffset
	}
	if offset > int64(len(r.Bytes())) {
		return 0, io.EOF
	}
	n := copy(b, r.Bytes()[offset:])
	if n < len(b) {
		return n, io.EOF
	}
	return len(b), nil
}

func (r *RawReader) setBytes(b []byte) {
	r.bytesMutex.Lock()
	defer r.bytesMutex.Unlock()
	r.bytes = b
}
