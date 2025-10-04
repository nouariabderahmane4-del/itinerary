package main
import (
	os / fmt //command line arguments handling, file I/O, error output.
	encoding / csv // Parsing the airport lookup CSV data.
	regexp // Finding and replacing all patterns (airport codes, date/time strings, whitespaces).
	time // Parsing and formatting ISO 8601 dates and times.
)

const usage = "itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv"

func main(){
	if len(os.Args) == 2 && os.Args[1] == "-h"{
	fmt.Println(usage)
	os.Exit(0)
	}

	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(1)
	}
	
}