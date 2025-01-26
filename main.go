package main

import (
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	defer goncurses.End()

	stdscr.Print("bgotop: btop but in go!")
	stdscr.MovePrint(3, 0, "Press any key to continue")

	stdscr.Refresh()

	stdscr.GetChar()
}
