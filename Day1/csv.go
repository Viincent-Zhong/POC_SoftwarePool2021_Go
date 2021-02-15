package main

func main() {
	// Exo 1, part 1
	/*arr_str, err := data.ReadFile((os.Args[1]))
	if err != nil {
		log.Fatalf("failed to open")
	}*/
	//data.Print_arr(arr_str)
	// Exo 1, part 2
	//data.Print_arr(data.LineToCSV(arr_str[0]))
	// Exo 1, part 3
	/*human := humanity.NewHumanFromCSV(data.LineToCSV(arr_str[0]))
	fmt.Println(human)*/
	/*human_arr := humanity.NewHumansFromCsvFile(os.Args[1])
	for _, each_line := range human_arr {
		fmt.Println(each_line)
	}*/
	// Exo 2
	/*human_arr := humanity.NewHumansFromJsonFile(os.Args[1])
	for i := 0; i < len(human_arr); i++ {
		fmt.Print("Name: ", human_arr[i].Name)
		fmt.Print(", Age:", human_arr[i].Age)
		fmt.Println(", Country", human_arr[i].Country)
	}*/
	// Exo 3
}
