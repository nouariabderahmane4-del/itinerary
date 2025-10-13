package processor

import (
	"bufio"
	"fmt"
	"itinerary/Packages/input_Package"
	"os"
	"regexp"
	"strings"
	"time"
)

// TrimSpace cleans up a string by:
// - Replacing carriage return, vertical tab, and form feed with newlines
// - Reducing multiple spaces to a single space
// - Replacing three or more consecutive newlines with two
// - Trimming leading and trailing whitespace
func TrimSpace(text string) string {
	text = strings.NewReplacer("\r", "\n", "\v", "\n", "\f", "\n").Replace(text)
	text = regexp.MustCompile(`[ ]{2,}`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
	text = strings.TrimSpace(text)
	return text
}

// ConvertATAcodes replaces an IATA airport code (#XXX) with the full airport name.
// Returns the original code if not found.
func ConvertATAcodes(match string, csvData []input.Airport) string {
	for _, airport := range csvData {
		if airport.Iata_code == match[1:] {
			return airport.Name
		}
	}
	return match
}

// ConvertICAOcodes replaces an ICAO airport code (##XXXX) with the full airport name.
// Returns the original code if not found.
func ConvertICAOcodes(match string, csvData []input.Airport) string {
	for _, airport := range csvData {
		if airport.Icao_code == match[2:] {
			return airport.Name
		}
	}
	return match
}

// replaceMunicipality replaces city codes (*#IATA or *##ICAO) with the airport's municipality name.
// Returns the original code if not found.
func replaceMunicipality(match string, csvData []input.Airport) string {
	for _, airport := range csvData {
		if strings.TrimPrefix(match, "*##") == airport.Icao_code || strings.TrimPrefix(match, "*#") == airport.Iata_code {
			return airport.Municipality
		}
	}
	return match
}

// ConvertTime parses ISO 8601 date/time strings and converts them to human-readable formats.
// Supports D(), T12(), and T24() prefixes.
func ConvertTime(match string, inputData string) string {
	timePrefix := strings.Split(match, "(")[0]
	timeWithoutParentheses := strings.TrimSuffix(strings.Split(match, "(")[1], ")")

	const ISO_Layout = "2006-01-02T15:04Z07:00"
	parsedTime, err := time.Parse(ISO_Layout, timeWithoutParentheses)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return match
	}

	formattedTime := ""
	switch timePrefix {
	case "D":
		formattedTime = parsedTime.Format("02 Jan 2006")
	case "T12":
		formattedTime = parsedTime.Format("03:04PM (-07:00)")
	case "T24":
		formattedTime = parsedTime.Format("15:04 (-07:00)")
	}
	return formattedTime
}

// Input_analyzing processes the input text:
// - Replaces airport codes with full names
// - Replaces city codes with municipality names
// - Converts ISO 8601 dates/times to readable formats
// - Normalizes whitespace
// Returns two versions: one for file output, one for colored terminal output.
func Input_analyzing(inputData string, csvData []input.Airport) (string, string) {
	colorToTerminal := inputData

	reg_city := regexp.MustCompile(`\*#+[A-Z]{3,4}\b`)
	reg_iata := regexp.MustCompile(`\#[A-Z]{3}\b`)
	reg_icao := regexp.MustCompile(`\##[A-Z]{4}\b`)
	reg_time := regexp.MustCompile(`(D|T12|T24)(\([^)]+\))`)

	// Replace city codes with municipality names
	matchCity := reg_city.FindAllString(inputData, -1)
	for _, match := range matchCity {
		cityName := replaceMunicipality(match, csvData)
		inputData = strings.ReplaceAll(inputData, match, cityName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+cityName+"\033[0m")
	}

	// Replace IATA codes with airport names
	matchIata := reg_iata.FindAllString(inputData, -1)
	for _, match := range matchIata {
		airportName := ConvertATAcodes(match, csvData)
		inputData = strings.ReplaceAll(inputData, match, airportName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+airportName+"\033[0m")
	}

	// Replace ICAO codes with airport names
	matchIcao := reg_icao.FindAllString(inputData, -1)
	for _, match := range matchIcao {
		airportName := ConvertICAOcodes(match, csvData)
		inputData = strings.ReplaceAll(inputData, match, airportName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+airportName+"\033[0m")
	}

	// Convert ISO 8601 time strings
	matchTime := reg_time.FindAllString(inputData, -1)
	for _, match := range matchTime {
		EnteredTime := ConvertTime(match, inputData)
		inputData = strings.ReplaceAll(inputData, match, EnteredTime)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[91m"+EnteredTime+"\033[0m")
	}

	return TrimSpace(inputData), TrimSpace(colorToTerminal)
}

// Final_Output writes the formatted output to a file and optionally displays it in the terminal.
// Also includes:
// - Typing animation
// - Colored output
// - Farewell message
// - ASCII airplane animation
func Final_Output(output_database string, output_file string, coloured_output string) {
	good_bye := "\nThank you for using Anywhere Holidays Prettifier Tool ✈️\n\nSee you soon! ✈️\n\n"
	output_done := fmt.Sprintf("\n-= Output successfully written to -> %s ✈️ =-\n\nDo you want to print the result in the command line? (Y/N) ✈️\n", output_file)

	// Write output to file
	err := os.WriteFile(output_file, []byte(output_database), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(0)
	}

	// Typing animation for output_done
	for _, letter := range output_done {
		time.Sleep(25 * time.Millisecond)
		fmt.Print(string(letter))
	}

	// Ask user if they want to print output
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		answer := scanner.Text()
		if answer == "n" || answer == "N" {
			break
		}
		if answer == "y" || answer == "Y" {
			fmt.Println()
			// Typing animation for colored output
			for _, letter := range coloured_output {
				time.Sleep(20 * time.Millisecond)
				fmt.Print(string(letter))
			}
			fmt.Println()
			time.Sleep(500 * time.Millisecond)
			break
		}
		fmt.Println("Invalid answer. Please enter Y or N.")
	}

	// Typing animation for goodbye message
	for _, letter := range good_bye {
		time.Sleep(20 * time.Millisecond)
		fmt.Print(string(letter))
	}

	// --- ASCII airplane animation ---
	plane := []string{
		"            ______            ",
		"            _\\ _~ -\\___      ",
		"    =  = ==(____AA____D      ",
		"                \\_____\\___________________,-~~~~~~~`-.._",
		"                /     o O o o o o O O o o o o o o O o  |\\_",
		"                `~-.__        ___..----..                  )",
		"                      `---~~\\___________/------------`````",
		"                      =  ===(_________D",
	}

	moves := 10      // number of steps the plane moves before stopping
	height := len(plane)

	fmt.Println() // start on a new line

	// Animate the airplane moving across the terminal
	for i := 0; i <= moves; i++ {
		if i > 0 {
			// Move cursor up to overwrite previous plane
			fmt.Print(fmt.Sprintf("\033[%dA", height))
		}
		for _, line := range plane {
			fmt.Print("\r" + strings.Repeat(" ", i*2) + line + "\n")
		}
		time.Sleep(120 * time.Millisecond)
	}
}
