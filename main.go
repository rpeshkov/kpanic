package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("kpanic", "Command-line utility for diagnosing kernel panics.")

	ls = app.Command("ls", "List available panics.")

	show     = app.Command("show", "Display panic contents.")
	showName = show.Arg("name", "Kernel panic filename").
			Required().
			String()

	last = app.Command("last", "Display last panic.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case ls.FullCommand():
		panics, err := listPanics(diagnosticsFolder)
		if err != nil {
			panic(err)
		}

		for _, p := range panics {
			fmt.Println(p)
		}

	case show.FullCommand():
		err := displayPanic(*showName)

		if err != nil {
			panic(err)
		}

	case last.FullCommand():
		panics, err := listPanics(diagnosticsFolder)
		if err != nil {
			panic(err)
		}

		if len(panics) == 0 {
			fmt.Println("No panics found")
			return
		}

		lastPanic := panics[len(panics)-1]
		err = displayPanic(lastPanic)
		if err != nil {
			panic(err)
		}
	}
}
