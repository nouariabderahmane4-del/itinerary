package main

import (
	"itinerary/Packages/input_Package"
	"itinerary/Packages/processor_Package"
	"os"
)

func main() {
	//Check the arguments
	input_file, output_file, csv_file := input.Check_args(os.Args)
	//Read the input file
	input_database := input.Read_txt(input_file)
	//Read the csv file and make a database
	csv_database := input.Read_csv(csv_file)
	//Process the input (airport codes, whitespaces, timezones)
	output_database, coloured_database := processor.Input_analyzing(input_database, csv_database)
	//Write the output file
	processor.Final_Output(output_database, output_file, coloured_database)
}