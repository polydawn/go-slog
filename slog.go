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

/*
	Produce a new `Slog`.  All of the output will be written to the provided
	`io.Writer` (often `os.Stderr` is a reasonable choice for this).  Every
	time the status area needs to be updated, `statFunc` is called, and is
	provided another writer to use.
*/
func New(proxy io.Writer, statFunc func(io.Writer)) *Slog {
	return &Slog{
		wr:    proxy,
		opine: statFunc,
	}
}

/*
	Write a message to the scrollback -- it will appear just above the
	updating status area.

	This implements the `io.Writer` contract.  Proxy other loggers to this!
	Then you get your choice of logging framework, AND the ability to do
	interactively updating status info with Slog.
*/
func (slog *Slog) Write(msg []byte) (int, error) {
	slog.retract()             // bring the cursor back up
	n, e := slog.wr.Write(msg) // write the message
	slog.place()               // replace mls content
	return n, e                // return stats from the parameter's write
}

/*
	Update the status area.

	Calls the `statFunc` the Slog was initialized with, and the content it
	produces will replace the current status area.

	If you have a bunch of progress bars, they might all feed status events
	into one list of progress info -- call this whenever the list is updated.
	(You're free to do dedup of events or trigger this on a timer, whatever.)
*/
func (slog *Slog) Refresh() {
	slog.retract()
	slog.place()
}

/*
	Set the current status area adrift in the scrollback.  It will no longer
	be overwritten after you call `Drape()` -- as you continue to write
	other logs, the draped status text will just drift up in the scrollback
	like any other lines.  New status writes after `Drape` behave normally
	(they'll still auto-replace themselves).

	You might want to clear out your status text immediatley after calling
	`Drape()`; otherwise you'll get the whole text duplicated in the scrollback,
	and again below it as soon as you next call either `Refresh` or `Write`.
	Of course, you might want that, so the choice is yours.
*/
func (slog *Slog) Drape() {
	slog.mls.footprint = 0
}

func (slog *Slog) retract() {
	if slog.mls.footprint > 0 {
		//all these ops should work even on window ANSI.SYS.
		slog.wr.Write(LF)                                 // set cursor to beginning of line.
		fmt.Fprint(slog.wr, CSI, slog.mls.footprint, "A") // jump cursor up.  should be supported even on windows.
		fmt.Fprint(slog.wr, CSI, "J")                     // clear from cursor to end of screen.
	}
}

func (slog *Slog) place() {
	wc := &linecountingWriter{wr: slog.wr}
	slog.opine(wc)
	slog.mls.footprint = wc.n
}

type linecountingWriter struct {
	n  int
	wr io.Writer
}

func (lcw *linecountingWriter) Write(msg []byte) (int, error) {
	lcw.n += bytes.Count(msg, CR)
	return lcw.wr.Write(msg)
}
