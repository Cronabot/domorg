package commands

import (
	"fmt"
)

var Quit Command = Command{
	Name:        "quit",
	Aliases:     []string{"q", "exit"},
	Description: "Quit the program",
	Func:        runQuit,
}

func runQuit(args []string) error {
	return fmt.Errorf("ExitProgramSignal")
}
