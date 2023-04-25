package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/takutotsuchie/cash/command"
)

func main() {
	if len(os.Args) == 1 {
		command.Help()
		return
	}
	changeMonth := Logic()
	if changeMonth {
		fmt.Println("月が変わったので、もう一度やり直してください")
	}
	switch os.Args[1] {
	case "add":
		cost := os.Args[2]
		costInt, _ := strconv.Atoi(cost)
		command.Add(costInt)
		return
	case "ls":
		command.List()
		return
	case "sub":
		cost := os.Args[2]
		costInt, _ := strconv.Atoi(cost)
		command.Sub(costInt)
		return
	default:
		command.Help()
		return
	}
}
