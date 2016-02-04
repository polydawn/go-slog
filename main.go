package main

import (
	"fmt"
	"io"
	"os"
	"time"
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
	if len(mls.lines) == 0 {
		return
	}
	fmt.Fprint(wr, CSI, len(mls.lines), "A")
	// shrinking the mls is going to be a tricky case
	// also good fucking luck with a terminal resize that rewraps you
}

func chill() {
	time.Sleep(300 * time.Millisecond)
}

func main() {
	mls := &multilineStatus{}
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 1\n")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 2\n")
	mls.lines = append(mls.lines, "]]] uno")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 3\n")
	mls.lines = append(mls.lines, "]]] dos")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 4\n")
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 5\n")
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 6\n")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 7\n")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 8\n")
	mls.Render(os.Stdout)
	chill()
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 9\n")
	mls.Render(os.Stdout)
	chill()
	fmt.Print(CSI, CYAN, "asdf", CSI, RESET, "qwer 0\n")
	mls.Render(os.Stdout)
	chill()
}
