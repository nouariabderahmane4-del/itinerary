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
		os.Exit(0)
	}else if os.Args[1] != "./input.txt" || os.Args[2] != "./output.txt" || os.Args[3] != "./airport_lookup.csv"{
		fmt.Println("one or more arguments are wrong. Use -h for help")
		os.Exit(0)
	}
	fmt.Println("Proceeding with the program")


}
