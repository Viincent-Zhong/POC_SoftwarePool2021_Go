package humanity

import (
	"SoftwareGoDay1/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
	Ready   bool
}

func (hums *Human) String() string {
	/*var str string
	str = Join(hums.String(), ",")*/
	//return hums.String()
	return fmt.Sprintf("%v (%v years) %v", hums.Name, hums.Age, hums.Country)
}

func (h *Human) Prepare() error {
	h.Ready = true
	if h.Ready == true {
		fmt.Sprintf("is ready !")
	}
	return nil
}

func NewHumanFromCSV(csv []string) *Human {
	age, err := strconv.Atoi(csv[1])
	if err != nil {
		log.Fatalf("probleme convert")
	}
	human := Human{csv[0], age, csv[2], false}
	return &human
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	arr_str, err := data.ReadFile(path)
	human_arr := []*Human{}

	if err != nil {
		log.Fatalf("failed to open")
	}
	for _, each_line := range arr_str {
		human := new(Human)
		str, err := data.LineToCSV(each_line)
		if err != nil {
			return human_arr, errors.New("error")
		}
		human = NewHumanFromCSV(str)
		human_arr = append(human_arr, human)
	}
	return human_arr, nil
}

func NewHumansFromJsonFile(path string) ([]*Human, error) {
	file, err := ioutil.ReadFile(path)
	human_arr := []*Human{}

	err = json.Unmarshal([]byte(file), &human_arr)
	if err != nil {
		return human_arr, err
	}
	return human_arr, nil
}
