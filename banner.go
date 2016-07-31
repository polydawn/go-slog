package slog

import (
	"fmt"
	"io"
)

type Banner struct {
	Lines []string
}

func (b Banner) WriteTo(wr io.Writer) {
	for _, line := range b.Lines {
		fmt.Fprint(wr, line, "\n")
	}
}
