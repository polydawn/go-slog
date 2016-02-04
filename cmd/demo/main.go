package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"polydawn.net/go-slog"
)

const (
	CSI   = "\x1B["
	CYAN  = "36m"
	RESET = "m"
)

type multilineStatus struct {
	lines []string
}

func (mls *multilineStatus) Render(wr io.Writer) {
	for _, line := range mls.lines {
		fmt.Fprint(wr, line, "\n")
	}
}

func chill() {
	time.Sleep(300 * time.Millisecond)
}

func main() {
	mls := &multilineStatus{}
	slog := slog.New(os.Stderr, mls.Render)
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 1\n")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 2\n")
	mls.lines = append(mls.lines, "]]] uno")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 3\n")
	mls.lines = append(mls.lines, "]]] dos")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 4\n")
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 5\n")
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 6\n")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 7\n")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 8\n")
	chill()
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 9\n")
	chill()
	fmt.Fprint(slog, CSI, CYAN, "asdf", CSI, RESET, "qwer 0\n")
	chill()
}
