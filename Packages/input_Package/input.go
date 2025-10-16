package input

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Airport holds all the data for a single airport.
// Each field corresponds to a column in the CSV file.
type Airport struct {
	Name			string
	Iso_country		string
	Municipality	string
	Icao_code		string
	Iata_code		string
	Coordinates		string
}

//Check_args is where we make sure that the user gave us the right number of arguments.
//If they didn't or if they ask for help with "-h", we print how to use the programme and exit.
func Check_args(args []string)(string, string, string){

	//Check if the user gave too few or too many arguments, or asked for help.
	if len(args) != 4 || len(args) == 2 && args[1] == "-h"{
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
		os.Exit(0)
	}

	//if everything is fine, we just return the paths for input, output, and CSV lookup.
	return args[1], args[2], args[3]
}

//Read_txt reads the entire input text file and gives it back as a string.
//If the file can't be found or opened, we just stop the program and show an error.
func Read_txt(input_file string)string{
	//try to read the whole file.
	file_content, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("\nInput file not found\n")
		os.Exit(0)
	}
	//Return the content as a string.
	return string(file_content)
}

//Read_csv reads our airport lookup csv and turns each row into an Airport struct.
//We also do some checks to make sure the CSV isn't broken or missing fields (MALFORMED).
func Read_csv(csv_file string) []Airport {
	//Try to open CSV file
	lookup, err := os.Open(csv_file)
	if err != nil {
		fmt.Print("\nAirport lookup not found\n")
		os.Exit(0)
	}
	defer lookup.Close()
	
	//Create a Csv reader.
	reader := csv.NewReader(lookup)

	//Read the header row first so we know which column is which.
	header, err := reader.Read()
	if err != nil {
		return nil
	}
	airport_data:= []Airport {} 			//This will store all our airport structs.
	lineNumber := 1							//Keep track of which line we are on for error messages.
	columnPositions := make(map[string]int) //Map column names to their positions
	
	//Does CSv has 6 columns?
	if len(header) < 6 {
		fmt.Println("\nAirport lookup malformed\n")
		os.Exit(0)
	}

	//Going through the header and mapping each column to its index
	for i, column := range header {
		if column == "" {
			fmt.Println("\nAirport lookup malformed\n")
			os.Exit(0)
		}
		columnPositions[column] = i
	}

	//We read exh row of the CSV
	for {
		line, err := reader.Read()
			if err == io.EOF {
				break // End of the file
			}
		if err != nil {
			fmt.Println("\nAirport lookup malformed\n")
			os.Exit(0)
		}
	//Checking if any field in this line is empty
	for _, field := range line {
		if strings.TrimSpace(field) == "" {
			fmt.Printf("\nAirport lookup malformed\n")
			os.Exit(0)
		}
	}
	lineNumber++

	//Create a new airport struct from this row
	airport := Airport{
		Name:			line[columnPositions["name"]],
		Iso_country:	line[columnPositions["iso_country"]],
		Municipality:	line[columnPositions["municipality"]],
		Icao_code:		line[columnPositions["icao_code"]],
		Iata_code:		line[columnPositions["iata_code"]],
		Coordinates:	line[columnPositions["coordinates"]],
	}

	//Add this airport to our list
	airport_data = append(airport_data, airport)
	}

	//Return the full list of airports
	return airport_data
}
