package main

import (
	"bufio"
	"commands"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Starting domorg...")

	commands.LoadCommands()

	defer fmt.Println("Exiting domorg...")

	commandListString := ""

	for i, cmd := range commands.CommandList {
		if i == len(commands.CommandList)-1 {
			commandListString += cmd.Name
		} else {
			commandListString += cmd.Name + ", "
		}
	}

	fmt.Printf("Loaded %v Commands: (%v)\n", len(commands.CommandList), commandListString)

	fmt.Print("> ")

	for scanner.Scan() {
		inp := scanner.Text()

		cmd := strings.Split(inp, " ")[0]
		args := strings.Split(inp, " ")[1:]

		err := commands.HandleCommand(cmd, args)
		if err != nil {
			if strings.Contains(err.Error(), "ExitProgramSignal") {
				break
			} else {
				fmt.Println(err)
			}
		}

		fmt.Print("> ")
	}
}
