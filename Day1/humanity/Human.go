package humanity

import (
	"SoftwareGoDay1/data"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) *Human {
	age, err := strconv.Atoi(csv[1])
	if err != nil {
		log.Fatalf("probleme convert")
	}
	human := Human{csv[0], age, csv[2]}
	return &human
}

func NewHumansFromCsvFile(path string) []*Human {
	arr_str, err := data.ReadFile(path)
	human_arr := []*Human{}

	if err != nil {
		log.Fatalf("failed to open")
	}
	for _, each_line := range arr_str {
		human := new(Human)
		human = NewHumanFromCSV(data.LineToCSV(each_line))
		human_arr = append(human_arr, human)
	}
	return human_arr
}

func NewHumansFromJsonFile(path string) []*Human {
	file, _ := ioutil.ReadFile(path)
	human_arr := []*Human{}

	_ = json.Unmarshal([]byte(file), &human_arr)
	return human_arr
}
