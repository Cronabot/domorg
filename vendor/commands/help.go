package commands

import (
	"fmt"
)

var Help Command = Command{
	Name:        "help",
	Aliases:     []string{"h", "?"},
	Description: "Displays a list of commands or information about a specific command",
	Func:        runHelp,
}

func runHelp(args []string) error {
	if len(args) == 0 {
		fmt.Println("Commands: ")
		for _, cmd := range CommandList {
			fmt.Println(cmd.Name + ": " + cmd.Description)
		}
		return nil
	}

	for _, cmd := range CommandList {
		if cmd.Name == args[0] {
			fmt.Println(cmd.Name)
		}
	}
	return nil
}
