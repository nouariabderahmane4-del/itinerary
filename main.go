package main
import (
	os / fmt //command line arguments handling, file I/O, error output.
	encoding / csv // Parsing the airport lookup CSV data.
	regexp // Finding and replacing all patterns (airport codes, date/time strings, whitespaces).
	time // Parsing and formatting ISO 8601 dates and times.
)

const usage = "itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv"

type Airport struct {
	Name string
	City string
}





























func main(){
	if len(os.Args) == 2 && os.Args[1] == "-h"{
	fmt.Println(usage)
	os.Exit(0)
	}

	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(1) //Failure
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	lookupPath := os.Args[3]

	_, err := os.ReadFile(inputPath)
	if err != nil {
		if os.IsNotExist(err){
			fmt.Println("Input not found")
			os.Exit(1)
		}
		fmt.Println("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	_, err := os.ReadFile(lookupPath)
	if err != nil {
		if os.IsNotExist(err){
			fmt.Println("Airport lookup  not found")
			os.Exit(1)
		}
		fmt.Println("Error reading Airport lookup file: %v\n", err)
		os.Exit(1)
	}
}