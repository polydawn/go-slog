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
	partial    *bytes.Buffer
	scwr       *scrollbackWriter
}

func newSlog(wr io.Writer) *Slog {
	s := &Slog{
		wr:      wr,
		partial: &bytes.Buffer{},
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
	// scan for the last (or any) cuts so we can decide what to do
	lastBreak := bytes.LastIndexByte(msg, '\n') + 1
	//	fmt.Printf("::> %q ; %d/%d <::\n", string(msg), lastBreak, len(msg))
	if lastBreak <= 0 {
		s.partial.Write(msg)
		return
	}
	// roll back over our previous banner print
	s.retract(s.footprint)
	// if we have any buffered partial lines, flush that first
	if s.partial.Len() > 0 {
		io.Copy(s.wr, s.partial)
		s.partial.Reset()
	}
	// now flush all other complete lines
	io.Copy(s.wr, bytes.NewBuffer(msg[0:lastBreak]))
	// and the banner again
	io.Copy(s.wr, bytes.NewBuffer(s.bannerMemo))
	// retain anything after lastbreak in the partial line buffer
	s.partial.Write(msg[lastBreak:])
}

type scrollbackWriter struct {
	s *Slog
}

func (scwr *scrollbackWriter) Write(msg []byte) (int, error) {
	scwr.s.write(msg)
	return len(msg), nil
}
