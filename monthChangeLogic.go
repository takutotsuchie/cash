package main

import (
	"strconv"
	"time"

	"github.com/takutotsuchie/cash/file"
	"github.com/takutotsuchie/cash/types"
)

func Logic() bool {
	ls := file.ReadFile()
	nowYear := time.Now().Year()
	nowMonth := int(time.Now().Month())
	yearString := strconv.Itoa(nowYear)
	monthString := strconv.Itoa(nowMonth)
	now := yearString + "-" + monthString

	// 今月の出費を改変したい。
	num := len(ls) - 1
	// 今月の初日の時は、新しい日付を追加。
	if now != ls[num].Month {
		ls = append(ls, types.MonthCash{Month: now, Cash: 0})
		file.WriteFile(ls)
		return true
	}
	return false
}
