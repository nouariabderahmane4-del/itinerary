package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {


	if len(os.Args)==2 && os.Args[1] == "-h" {
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
	}else if len(os.Args) != 4 {
		fmt.Println("Wrong number of arguments. Use -h for help")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

	if strings.Contains(line, "#"){
		fmt.Println("City name code was found:", line)
	}
}
if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
