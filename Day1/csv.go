package main

import (
	"SoftwareGoDay1/humanity"
	"fmt"
)

func main() {
	// Exo 1, part 1: Lire un fichier
	/*arr_str, err := data.ReadFile((os.Args[1]))
	if err != nil {
		log.Fatalf("failed to open")
	}*/
	//data.Print_arr(arr_str)
	// Exo 1, part 2 : Traiter les données (coupe la string en fonction des virgules)
	/* str, err := data.LineToCSV(arr_str[0])
	if err != nil
		log.Fatalf("pas de virgule")
	data.Print_arr(data.LineToCSV(arr_str[0]))
	*/
	// Exo 1, part 3 : Créer un tableau de structure qui contient chaque users séparement
	/*human := humanity.NewHumanFromCSV(data.LineToCSV(arr_str[0]))
	fmt.Println(human)*/
	/*human_arr, err := humanity.NewHumansFromCsvFile(os.Args[1])
	if err != nil {
		log.Fatalf("problem")
	}
	for _, each_line := range human_arr {
		fmt.Println(each_line)
	}*/
	// Exo 2 : Pareil que le exo 1 mais avec un fichier JSON
	/*human_arr, err := humanity.NewHumansFromJsonFile(os.Args[1])
	if err != nil {
		log.Fatalf("error")
	}
	for i := 0; i < len(human_arr); i++ {
		fmt.Print("Name: ", human_arr[i].Name)
		fmt.Print(", Age:", human_arr[i].Age)
		fmt.Println(", Country", human_arr[i].Country)
	}*/
	// Exo 3 : Formatter le code
	/* Je crois que vscode le fait automatiquement*/
	// Exo 4 : Test en go
	// Exo 5 : créer des pilotes
	fmt.Println(&humanity.Pilot{&humanity.Human{"Jason", 10, "Fr", false}})
	//human := humanity.Human{"Jason", 10, "Fr"}
	//fmt.Println(human.String())
}

/*func (h *Human) String() string {

}*/
