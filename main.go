package main

import (
	"itinerary/Packages/input_Package"     // Handles input file reading and CSV airport lookup
	"itinerary/Packages/processor_Package" // Handles processing of airport codes, date/time, and output
	"os"
)

func main() {
	// Step 1: Validate command-line arguments
	// Expects three arguments: input file path, output file path, CSV lookup file path
	// Also supports "-h" to display usage
	input_file, output_file, csv_file := input.Check_args(os.Args)

	// Step 2: Read the input text file
	// Reads the entire contents of the input file into a string
	input_database := input.Read_txt(input_file)

	// Step 3: Read the CSV airport lookup file
	// Converts each row of the CSV into an Airport struct for easy lookup
	csv_database := input.Read_csv(csv_file)

	// Step 4: Process the input text
	// - Replaces IATA (#XXX) and ICAO (##XXXX) codes with full airport names
	// - Replaces city codes (*#XXX, *##XXXX) with city/municipality names
	// - Converts ISO 8601 dates/times (D(), T12(), T24()) into human-readable formats
	// - Cleans up whitespace and normalizes blank lines
	// Returns both a plain version for the output file and a colored version for terminal display
	output_database, coloured_database := processor.Input_analyzing(input_database, csv_database)

	// Step 5: Write the processed output to the output file
	// Optionally displays colored output in the terminal with typing animation and farewell message
	processor.Final_Output(output_database, output_file, coloured_database)
}
