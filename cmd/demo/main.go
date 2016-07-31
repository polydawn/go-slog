package main

import (
	"fmt"
	"os"
	"time"

	"go.polydawn.net/go-slog"
)

const (
	CSI   = "\x1B["
	CYAN  = "36m"
	RESET = "m"
)

func chill() {
	time.Sleep(500 * time.Millisecond)
}

func main() {
	mls := slog.Banner{}
	slog, std := slog.New(os.Stderr)
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 1 .\n")
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 2  .\n")
	chill()
	mls.Lines = append(mls.Lines, "]]] uno .")
	slog.SetBanner(mls)
	chill()
	slog.SetBanner(mls)
	chill()
	slog.SetBanner(mls)
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 3   .\n")
	chill()
	mls.Lines = append(mls.Lines, "]]] dos  .")
	slog.SetBanner(mls)
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 4    .\n")
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 5     .\n")
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 6    .\n")
	chill()
	mls.Lines = append(mls.Lines, "]]] tres  .")
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 7   .\n")
	//	slog.Drape()
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 8  .\n")
	chill()
	mls.Lines = mls.Lines[0:1]
	slog.SetBanner(mls)
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 9 .\n")
	chill()
	fmt.Fprint(std, CSI, CYAN, "asdf", CSI, RESET, "qwer 0.\n")
	chill()
}
