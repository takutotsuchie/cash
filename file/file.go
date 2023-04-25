package file

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/takutotsuchie/cash/types"
)

var file = os.Getenv("cash_json")

func ReadFile() []types.MonthCash {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var ls []types.MonthCash
	err = json.Unmarshal(data, &ls)
	if err != nil {
		log.Fatal(err)
	}
	return ls
}

func WriteFile(ls []types.MonthCash) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	data := jsonMarshaler(ls)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func jsonMarshaler(ls []types.MonthCash) []byte {
	var ans []byte
	ans = append(ans, []byte("[\n")...)
	for index, cash := range ls {
		data, err := json.Marshal(cash)
		if err != nil {
			log.Print(err)
		}
		ans = append(ans, data...)
		if index != len(ls)-1 {
			ans = append(ans, []byte("\n")...)
		}
	}
	ans = append(ans, []byte("\n]")...)
	return ans
}
