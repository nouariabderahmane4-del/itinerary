package input

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)
type Airport struct {
	Name			string
	Iso_country		string
	Municipality	string
	Icao_code		string
	Iata_code		string
	Coordinates		string
}

func TestArgs(args []string)(string, string, string){

	if len(args) != 4 || len(args) == 2 && args[1] == "-h"{
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
		os.Exit(0)
	}
	return args[1], args[2], args[3]
}

func GetText(input_file string)string{
	file_content, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Input file not found ", err)
		os.Exit(0)
	}
	return string(file_content)
}
func GetCsv(csv_file string) []Airport {
	lookup, err := os.Open(csv_file)
	if err != nil {
		fmt.Print("Airport lookup not found", err)
		os.Exit(1)
	}
	defer lookup.Close()

	reader := csv.NewReader(lookup)
	header, err := reader.Read()
	if err != nil {
		return nil
	}
	airport_data:= []Airport {}
	lineNumber := 1

	columnPositions := make(map[string]int)
	if len(header) < 6 {
		fmt.Println("Airport lookup malformed (Missing column)")
		os.Exit(0)
	}

	for i, column := range header {
		if column == "" {
			fmt.Println("Airport lookup malformed (Empty column)")
			os.Exit(0)
		}
		columnPositions[column] = i
	}

	for {
		line, err := reader.Read() {
			if err == io.EO {
				break
			}
		}
		if err != nil {
			fmt.Println("Airport lookup malformed", err)
			os.Exit(0)
		}
	for _, field := range line {
		if strings.trimspace(field) == "" {
			fmt.Printf("Airport lookup malformed (Empty field on line %d)\n", lineNumber)
			os.Exit(0)
		}
	}
	lineNumber++
	airport := Airport{
		Name:			line[columnPositions["name"]],
		Iso_country:	line[columnPositions["iso_country"]],
		Municipality:	line[columnPositions["municipality"]],
		Icao_code:		line[columnPositions["icao_code"]],
		Iata_code:		line[columnPositions["iata_code"]],
		Coordinates:	line[columnPositions["coordinates"]],
	}
	airport_data = append(airport_data, airport)
}
return airport_data
}
