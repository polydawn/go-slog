package slog

import (
	"bytes"
	"fmt"
	"io"
)

const (
	CSI = "\x1B["
)

var (
	CR = []byte{0x0A}
	LF = []byte{0x0D}
)

/*
	Fundamentally we don't really care about anything except backing up far enough
	to overwrite our own previous footprint.  So here's what you remember for that.
*/
type multilineStatus struct {
	footprint int
}

// Slog must be a writer so you can use it in place of `os.Stderr` in other code.
var _ io.Writer = &Slog{}

type Slog struct {
	mls   multilineStatus
	wr    io.Writer
	opine func(io.Writer)
}

func New(proxy io.Writer, statFunc func(io.Writer)) *Slog {
	return &Slog{
		wr:    proxy,
		opine: statFunc,
	}
}

func (slog *Slog) Write(msg []byte) (int, error) {
	// back up.
	if slog.mls.footprint > 0 {
		//all these ops should work even on window ANSI.SYS.
		slog.wr.Write(LF)                                 // set cursor to beginning of line.
		fmt.Fprint(slog.wr, CSI, slog.mls.footprint, "A") // jump cursor up.  should be supported even on windows.
		fmt.Fprint(slog.wr, CSI, "J")                     // clear from cursor to end of screen.
	}

	// write the message
	n, e := slog.wr.Write(msg)

	// replace mls content
	wc := &linecountingWriter{wr: slog.wr}
	slog.opine(wc)
	slog.mls.footprint = wc.n

	// return stats from the parameter's write
	return n, e
}

type linecountingWriter struct {
	n  int
	wr io.Writer
}

func (lcw *linecountingWriter) Write(msg []byte) (int, error) {
	lcw.n += bytes.Count(msg, CR)
	return lcw.wr.Write(msg)
}
