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