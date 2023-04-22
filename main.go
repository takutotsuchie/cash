package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var file = os.Getenv("cash_json")

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}
	ls := ReadFile()
	nowYear := time.Now().Year()
	nowMonth := int(time.Now().Month())
	yearString := strconv.Itoa(nowYear)
	monthString := strconv.Itoa(nowMonth)
	now := yearString + "-" + monthString

	// 今月の出費を改変したい。
	num := len(ls.CashList) - 1
	// 今月の初日の時は、新しい日付を追加。
	if now != ls.CashList[num].Month {
		ls.CashList = append(ls.CashList, MonthCash{now, 0})
		WriteFile(ls)
		num++
	}
	switch os.Args[1] {
	case "add":
		cost := os.Args[2]
		costInt, _ := strconv.Atoi(cost)
		ls.CashList[num].Cash += costInt
		WriteFile(ls)
		return
	case "ls":
		for _, l := range ls.CashList {
			fmt.Println("月:", l.Month, "出費", l.Cash)
		}
		return
	case "sub":
		cost := os.Args[2]
		costInt, _ := strconv.Atoi(cost)
		ls.CashList[num].Cash -= costInt
		WriteFile(ls)
		return
	default:
		help()
		return
	}
}
