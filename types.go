package main

type MonthCash struct {
	Month string `json:"month"`
	Cash  int    `json:"cash"`
}
type CashLs struct {
	CashList []MonthCash `json:"cashlist"`
}
