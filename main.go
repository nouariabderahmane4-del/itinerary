package main
import (
	"os" //command line arguments handling, file I/O, error output.
	"fmt" 
	"encoding/csv" // Parsing the airport lookup CSV data.
	"regexp" // Finding and replacing all patterns (airport codes, date/time strings, whitespaces).
	"time" // Parsing and formatting ISO 8601 dates and times.
	"errors"
	"strings"
)

const usage = "itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv"

type Airport struct {
	Name string
	City string
}

var LookupTable = make(map[string]Airport) //to store the lookup map.


var requiredHeaders = [string]{ //Required headers for malformed data check.
	"name",
	"municipality",
	"icao_code",
	"iata_code",
}

// Regex constants 

var airportCodeRegex = regexp.MustCompile(`(\*?)(#|##)([A-Z]{3,4})`)// To match airport codes.
var dateTimeRegex = regexp.MustCompile(`(D|T12|T24)\((.*?)\)`)// To match date/time patterns.
var verticalWhitespaceRegex = regexp.MustCompile(`[\v\f\r]`) // For vertical whitespaces characters
var excessiveNewLineRegex = regexp.MustCompile(`\n{3,}`) // For excessive newlines

// Error definition (I use os.Exit(0) )
var ErrMalformed = error.new("Airport lookup malformed")

//Helper functions

func checkSlice(s []string, e string) bool{
	for _, a := range s {
		if a = e{
			return true
		}
	}
	return false
}


//Applying the specific time formatting and offset display rules
func formatTimeWithOffset(t time.Time, format string) string{
	//Getting the offset from UTC in seconds (+00:00)
	_, offsetSeconds := t.Zone()

	if offsetSeconds == o{
		return t.Format(format) + " (+00:00)"
	}

	offsetString := t.Format("-07:00")
	//string manipulation
	if strings.HasPrefix(offsetString, "+"){
		offsetString = string.Replace(offsetString, "+", "(+", 1)
	}else if strings.HasPrefix(offsetString, "-"){
		offsetString = string.Replace(offsetString, "-", "(-", 1)
	}
	// Time + (Offset)
	return t.Format(format) + " " + offsetString + ")"


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