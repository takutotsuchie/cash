package command

import (
	"fmt"

	"github.com/takutotsuchie/cash/file"
)

func List() {
	ls := file.ReadFile()
	for _, l := range ls {
		fmt.Println("月:", l.Month, "出費", l.Cash)
	}
}
