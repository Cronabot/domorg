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
			aliasesString := ""
			for i, alias := range cmd.Aliases {
				if i == len(cmd.Aliases)-1 {
					aliasesString += alias
				} else {
					aliasesString += alias + ", "
				}
			}
			fmt.Printf("%s (%s): %s\n", cmd.Name, aliasesString, cmd.Description)
		}
		return nil
	}

	return nil
}
