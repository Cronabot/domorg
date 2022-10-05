package commands

import (
	"fmt"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Func        func(args []string) error
}

var CommandList []Command

func LoadCommands() {
	CommandList = append(CommandList,
		// Add commands here
		Quit,
		Help,
	)
}

func GetCommand(name string) (error, *Command) {
	for _, cmd := range CommandList {
		if cmd.Name == name {
			return nil, &cmd
		}
	}
	return fmt.Errorf("Command not found"), nil
}

func HandleCommand(command string, args []string) error {
	commandFound := false
	for _, cmd := range CommandList {
		if cmd.Name == command {
			commandFound = true
			err := cmd.Func(args)
			if err != nil {
				return fmt.Errorf("Error running command: %s\n%v", cmd.Name, err.Error())
			}
		}
		for _, alias := range cmd.Aliases {
			if alias == command {
				commandFound = true
				err := cmd.Func(args)
				if err != nil {
					return fmt.Errorf("Error running command \"%s\": %v", cmd.Name, err.Error())
				}
			}
		}
	}
	if !commandFound {
		return fmt.Errorf("Command not found")
	}

	return nil
}
