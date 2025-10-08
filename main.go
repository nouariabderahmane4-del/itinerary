package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)==2 && os.Args[1] == "-h" {
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
	}else if len(os.Args) != 4 {
		fmt.Println("Wrong number of arguments. Use -h for help")
	}

}
