package main

import "fmt"

func help() {
	fmt.Println("Usage:")
	fmt.Println("   	`cash add 'your expense'` to register cost")
	fmt.Println("   	`cash ls` to list your monthly cost")
	fmt.Println("   	`cash sub 'your expense'` to subtract your cost ")
}
