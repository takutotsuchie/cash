package command

import "github.com/takutotsuchie/cash/file"

func Sub(cost int) {
	ls := file.ReadFile()
	nowMonth := len(ls)
	ls[nowMonth-1].Cash -= cost
	file.WriteFile(ls)
}
