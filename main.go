package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"golang.org/x/term"
)

type asciiArt string

//nolint:gochecknoglobals
var (
	stdout = colorable.NewColorableStdout()
	stderr = colorable.NewColorableStderr()
)

//nolint:gochecknoglobals
var gopher = []asciiArt{
	`                                   `,
	`     CCCCCCCCCCCCCCCCCCCCCCCC      `,
	` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
	`CCCCCCCWWWWWWWWCCCCWWWWWWWCCCCCCCC `,
	`CCCCCCCWBBWWWWWCCCCWBBWWWWCCCCCCCC `,
	` CCCCCCWBBWWWWWCCCCWBBWWWWCCCCCCC  `,
	`   CCCCCCCCCCCCCBBBCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCYYYYYYYCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCWWWCCCCCCCCCCCC    `,
	` YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
	` YYYYYCCCCCCCCCCCCCCCCCCCCCCYYYYY  `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`                                   `,
}

//nolint:gochecknoglobals
var runningGopher = [][]asciiArt{
	{
		`                                   `,
		`     CCCCCCCCCCCCCCCCCCCCCCCC      `,
		` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
		`CCCCCCCWWWWWWWWCCCCWWWWWWWCCCCCCCC `,
		`CCCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCCC `,
		` CCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCC  `,
		`   CCCCCCCCCCCCCBBBCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCYYYYYYYCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
		`   CCCCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
		` YYYYYCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		` YYYYYCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`     CCCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`                                   `,
	},
	{
		`                                   `,
		`     CCCCCCCCCCCCCCCCCCCCCCCC      `,
		` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
		`CCCCCCCWWWWWWWWCCCCWWWWWWWCCCCCCCC `,
		`CCCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCCC `,
		` CCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCC  `,
		`   CCCCCCCCCCCCCBBBCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCYYYYYYYCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
		` YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
		` YYYYYCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`                                   `,
	},
	{
		`                                   `,
		`     CCCCCCCCCCCCCCCCCCCCCCCC      `,
		` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
		`CCCCCCCWWWWWWWWCCCCWWWWWWWCCCCCCCC `,
		`CCCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCCC `,
		` CCCCCCWWWWWBBWCCCCWWWWBBWCCCCCCC  `,
		`   CCCCCCCCCCCCCBBBCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCYYYYYYYCCCCCCCCCC    `,
		` YYYYYCCCCCCCCCCWWWCCCCCCCCCCCC    `,
		` YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCYYYYY  `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
		`   CCCCCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
		`  YYYYYCCCCCCCCCCCCCCCCCCCCCC      `,
		`                                   `,
	},
}

//nolint:gochecknoglobals
var cryingGopher = []asciiArt{
	`                                   `,
	`   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^    `,
	`  <    _   _      _       _    >   `,
	`  <   | | | | ___| |_ __ | |   >   `,
	`  <   | |_| |/ _ \ | '_ \| |   >   `,
	`  <   |  _  |  __/ | |_) |_|   >   `,
	`  <   |_| |_|\___|_| .__/(_)   >   `,
	`  <                |_|         >   `,
	`  <vvvvvvvvvvvvvvvvvvvvv| /vvvv>   `,
	`                        |/         `,
	`                                   `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	` YYYYYCCCbbCCCCCCCCCCCCbbCCCYYYYY  `,
	` YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
	`   CCCCCCbbCCCCCWWWCCCCbbCCCCCC    `,
	`   CCCCCCbbCCCYYYYYYYCCbbCCCCCC    `,
	`   CCCCCCbbCCCCCBBBCCCCbbCCCCCC    `,
	` CCCCCCWWWWWWWWCCCCCWWWWWWWWCCCCC  `,
	`CCCCCCCBBBBBBWWCCCCCBBBBBBWWCCCCCC `,
	`CCCCCCCWWWWWWWWCCCCCWWWWWWWWCCCCCC `,
	` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
	`     CCCCCCCCCCCCCCCCCCCCCCCC      `,
	`                                   `,
}

//nolint:gochecknoglobals
var readingGopher = []asciiArt{
	`                                   `,
	`   ┌─────────────────────────────┐ `,
	`   │   I'm reading documents...  │ `,
	`   └───────────────────| /───────┘ `,
	`                       |/          `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`  YYYYYCCCCCCCCCCCCCCCCCCCCYYYYY   `,
	`   CCCCCCCCCCCCCCCCCCCCCCCCCCCC    `,
	`   CCGGGGGGGGGGGBBGGGGGGGGGGGCC    `,
	`   CCGGGGGGGGGGGBBGGGGGGGGGGGCC    `,
	`   CCGGGGGGGGGGGBBGGGGGGGGGGGCC    `,
	`   CCGGGGGGGGGGGBBGGGGGGGGGGGCC    `,
	`   CCGGGGGGGGGGGBBGGGBBBBBBGGCC    `,
	` YYYYGGGGGGGGGGGBBGGGGGGGGGGGYYYY  `,
	` YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY  `,
	`   CCCCCCCCCCCCCWWWCCCCCCCCCCCC    `,
	`   CCCCCCCCCCCYYYYYYYCCCCCCCCCC    `,
	`   CCCCCCCCCCCCCBBBCCCCCCCCCCCC    `,
	` CCCCCCWBBWWWWWCCCCWBBWWWWCCCCCCC  `,
	`CCCCCCCWBBWWWWWCCCCWBBWWWWCCCCCCCC `,
	`CCCCCCCWWWWWWWWCCCCWWWWWWWCCCCCCCC `,
	` CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC  `,
	`     CCCCCCCCCCCCCCCCCCCCCCCC     `,
	`                                   `,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(stderr, "must require arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		runHelpCommand()
		os.Exit(0)
	case "run":
		runRunCommand()
		os.Exit(0)
	case "doc":
		runDocCommand()
		os.Exit(0)
	default:
		break
	}
}

func runHelpCommand() {
	aa := colorizeASCIIArt(cryingGopher)
	for _, line := range aa {
		fmt.Fprintln(stdout, line)
	}
}

func runRunCommand() {
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintln(stderr, err)

		return
	}

	gopherWidth := len(runningGopher[0][0])
	gopherHeight := len(runningGopher[0])

	for pad := 0; pad < termWidth-gopherWidth; pad++ {
		aa := colorizeASCIIArt(runningGopher[pad%3])
		for _, line := range aa {
			fmt.Fprintf(stdout, "%s%s\n", strings.Repeat(" ", pad), line)
		}

		fmt.Fprint(stdout, backCursorAnsi(gopherHeight))
		fmt.Fprint(stdout, eraseDisplayAnsi())

		//nolint:gomnd
		time.Sleep(20 * time.Millisecond)
	}
}

/*
func eraseLineAnsi() string {
	return fmt.Sprintf("%s[%dK", "\x1b", 1)
}
*/

func eraseDisplayAnsi() string {
	return fmt.Sprintf("%s[%dJ", "\x1b", 1)
}

func backCursorAnsi(n int) string {
	return fmt.Sprintf("%s[%dF", "\x1b", n)
}

func runDocCommand() {
	aa := colorizeASCIIArt(readingGopher)
	for _, line := range aa {
		fmt.Fprintln(stdout, line)
	}
}

//nolint:varnamelen,prealloc,wsl
func colorizeASCIIArt(aa []asciiArt) []asciiArt {
	var colorized []asciiArt
	for _, line := range aa {
		a := line.Colorize('C', color.New(color.Bold, color.BgCyan)).
			Colorize('Y', color.New(color.Bold, color.BgYellow)).
			Colorize('W', color.New(color.Bold, color.BgWhite)).
			Colorize('B', color.New(color.Bold, color.BgBlack)).
			Colorize('b', color.New(color.Bold, color.BgBlue)).
			Colorize('G', color.New(color.Bold, color.BgGreen)).
			Colorize('R', color.New(color.Bold, color.BgRed))

		colorized = append(colorized, a)
	}

	return colorized
}

func (aa asciiArt) Colorize(target rune, color *color.Color) asciiArt {
	str := strings.ReplaceAll(
		string(aa),
		string(target),
		color.Sprint(" "),
	)

	return asciiArt(str)
}
