package slog

import "io"

/*
	Produce a new `Slog`.  All of the output will be written to the provided
	`io.Writer` (often `os.Stderr` is a reasonable choice for this).

	To set a banner of content that should be kept at the bottom of the screen, call `SetBanner`.
	To print regular messages to scrollback, just write them to the returned `io.Writer`.
	Both methods are safe to call concurrently with each other.
*/

func New(term io.Writer) (*Slog, io.Writer) {
	s := newSlog(term)
	return s, s.scwr
}
