package humanity

import (
	"SoftwareGoDay1/data"
	"log"
)

type Human struct {
	Name    string
	Age     string
	Country string
}

func NewHumanFromCSV(csv []string) *Human {
	human := Human{csv[0], csv[1], csv[2]}
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
