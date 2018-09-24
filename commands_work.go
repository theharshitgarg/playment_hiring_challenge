package main

import (
	"fmt"
	// "os"
	// "strconv"
)

type Command struct {
	Name string
	Desc string
	// ShortDesc   string
	Usage string
}

func GetCommandsMap() map[string]Command {
	commands := make(map[string]Command)
	commands["ls"] = Command{"ls", "Lists the directory contents", "ls <directory path>"}
	commands["pwd"] = Command{"ls", "Prints the path of current working directory", "pwd"}
	commands["rm"] = Command{"ls", "Removes a directory", "rm <directory path>"}
	commands["mkdir"] = Command{"mkdir", "Creates a directory", "mkdir <directory path>"}
	commands["help"] = Command{"help", "Gives information about help", "help [command-name]"}

	return commands
}

func PrintAvailableCommand(cmdName string) {
	commands := GetCommandsMap()
	cmd, err := commands[cmdName]

	if !err {
		fmt.Println("Error in command name. Try again with correct name.")
	}

	fmt.Println("Name : ", cmd.Name)
	fmt.Println("Desc : ", cmd.Desc)
	fmt.Println("Usage : ", cmd.Usage)
	fmt.Println()
}

func PrintAvailableCommands() {
	commands := GetCommandsMap()

	for _, cmd := range commands {
		fmt.Println("Name : ", cmd.Name)
		fmt.Println("Desc : ", cmd.Desc)
		fmt.Println("Usage : ", cmd.Usage)
		fmt.Println()
	}
}

func main() {
	var cmd string
	fmt.Println("Starting your application.....")
	fmt.Println("To get started, type help followed by enter")

	for true {
		fmt.Scanf("%s", &cmd)
		PrintAvailableCommand(cmd)
	}
}
