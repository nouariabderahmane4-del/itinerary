package main

import (
	"fmt"
	"os"
)

func main() {
	asking_for_help := os.Args[1]
	if asking_for_help == "-h" {
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
	}

}
