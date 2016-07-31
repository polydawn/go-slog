package slog

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

const (
	_CSI = "\x1B["
)

var (
	_CR = []byte{'\r'}
	_LF = []byte{'\n'}
)

type Slog struct {
	wr         io.Writer
	footprint  int
	bannerMemo []byte
	mu         sync.Mutex
	scwr       *scrollbackWriter
}

func newSlog(wr io.Writer) *Slog {
	s := &Slog{
		wr: wr,
	}
	s.scwr = &scrollbackWriter{s}
	return s
}

func (s *Slog) SetBanner(b Banner) {
	// Ask the new thing to fmt itself all out.
	var buf bytes.Buffer
	b.WriteTo(&buf)
	// Lock for the remainder.
	s.mu.Lock()
	defer s.mu.Unlock()
	// Save the new state.
	footprint := s.footprint
	s.bannerMemo = buf.Bytes()
	s.footprint = bytes.Count(s.bannerMemo, _LF)
	// Render.
	s.retract(footprint) // the previous one
	io.Copy(s.wr, bytes.NewBuffer(s.bannerMemo))
}

func (s *Slog) retract(n int) {
	if n > 0 {
		//all these ops should work even on windows ANSI.SYS.
		s.wr.Write(_CR)                // set cursor to beginning of line.
		fmt.Fprint(s.wr, _CSI, n, "A") // jump cursor up.  should be supported even on windows.
		fmt.Fprint(s.wr, _CSI, "J")    // clear from cursor to end of screen.
	}
}

// must be full lines or we will have a sad time
func (s *Slog) write(msg []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.retract(s.footprint)
	io.Copy(s.wr, bytes.NewBuffer(msg))
	io.Copy(s.wr, bytes.NewBuffer(s.bannerMemo))
}

type scrollbackWriter struct {
	s *Slog
}

func (scwr *scrollbackWriter) Write(msg []byte) (int, error) {
	last := bytes.LastIndexByte(msg, '\n') + 1
	scwr.s.write(msg[0:last])
	return last, nil
}
