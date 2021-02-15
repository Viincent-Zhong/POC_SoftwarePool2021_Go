package data

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Print_arr(arr_str []string) {
	for _, each_line := range arr_str {
		fmt.Println(each_line)
	}
}

func LineToCSV(line string) []string {
	str := strings.Split(line, ",")
	return str
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var arr_str []string
	for scanner.Scan() {
		arr_str = append(arr_str, scanner.Text())
	}
	file.Close()
	return arr_str, err
}
