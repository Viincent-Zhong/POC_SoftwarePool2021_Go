package main

import (
	"SoftwareGoDay1/humanity"
	"fmt"
	"os"
)

func main() {
	/*arr_str, err := data.ReadFile((os.Args[1]))
	if err != nil {
		log.Fatalf("failed to open")
	}*/
	//data.Print_arr(arr_str)
	//data.Print_arr(data.LineToCSV(arr_str[0]))
	/*human := humanity.NewHumanFromCSV(data.LineToCSV(arr_str[0]))
	fmt.Println(human)*/
	human_arr := humanity.NewHumansFromCsvFile(os.Args[1])
	for _, each_line := range human_arr {
		fmt.Println(each_line)
	}
}
