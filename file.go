package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile() CashLs {
	f, err := os.Open(file)
	fmt.Println(file)
	if err != nil {
		log.Print(err)
	}
	data, _ := io.ReadAll(f)
	var ls CashLs
	_ = json.Unmarshal(data, &ls)
	return ls
}

func WriteFile(ls CashLs) {
	f, _ := os.Create(file)
	data, _ := json.Marshal(ls)
	_, _ = f.Write(data)
}
