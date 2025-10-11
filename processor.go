package processor

import (
	"bufio"
	"fmt"
	"input"
	"os"
	"regexp"
	"strings"
	"time"

	input "main.go"
)

func TrimSpace(text string) string{
	text = strings.NewReplacer("\r", "\n", "\v", "\n", "\f", "\n").Replace(text)
	text = regexp.MustCompile(`[ ]{2,}`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
	text= strings.TrimSpace(text)
	return text
}

func ConvertATAcodes(match string, csvData []input.Airport) string {
	for _, airport := range csvData {
		if airport.Iata_code == match[1:]{
			return airport.Name
		}
	} 
	return match
}
func ConvertICAOcodes(match string, csvData []input.Airport) string {
	for _, airport := range csvData {
		if airport.Icao_code == match[2:] {
			return airport.Name
		}
	}
	return match
}
func replaceMunicipality(match string, csvData  []input.Airport) string {
	for _, airport := range csvData{
		if strings.TrimPrefix(match, "*##") == airport.Icao_code || strings.TrimPrefix(match, "*#") == airport.Iata_code{
			return airport.Municipality
		}
	}
	return match
}

func ConvertTime(match string, inputData string) string{

	timePrefix := strings.Split(match, "(")[0]
	timeWithoutParentheses := strings.TrimSuffix(strings.Split(match, "(")[1], ")")
	
	const ISO_Layout = "2006-01-02T15:04Z07:00"

	parsedTime, err := time.Parse(ISO_Layout, timeWithoutParentheses)
	if err != nil {
		fmt.Println("Error parsing time", err)
		return match
	}

	formattedTime := ""
	switch timePrefix{
	case  "D":
		formattedTime = parsedTime.Format("02 Jan 2006")
	case "T12":
		formattedTime = parsedTime.Format("03:04PM (Z07:00)")
	case "T24":
		formattedTime = parsedTime.Format("15:04 (Z07:00)") 
	}
	return formattedTime
}

func Input_analyzing(inputData string, csvData []input.Airport) (string, string) {

	colorToTerminal := inputData

	reg_city := regexp.MustCompile(`\*#+[A-Z]{3,4}\b`)
	reg_iata := regexp.MustCompile(`\#[A-Z]{3}\b`)
	reg_icao := regexp.MustCompile(`\##[A-Z]{4}\b`)
	reg_time := regexp.MustCompile(`(D|T12|T24)(\([^)]+\))`)

	matchCity := reg_city.FindAllAtring(inputData, -1)
	for _, match := range matchCity{
		cityName := replaceMunicipality(match, csvData)
		input_data = strings.ReplaceAll(inputData, match, cityName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+cityName+"\033[0m")
	}

	matchIata := reg_iata.FindAllAtring(inputData, -1)
	for _, match := range matchIata {
		airportName := ConvertATAcodes(match, csvData)
		inputData = strings.ReplaceAll(inputData, match, airportName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+cityName+"\033[0m")
	}

	matchIcao := reg_icao.FindAllAtring(inputData, -1)
	for _, match := range matchIcao {
		airportName := ConvertICAOcodes(match, csvData)
		inputData = strings.ReplaceAll(inputData, match, airportName)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[36m"+cityName+"\033[0m")
	}

	matchTime := reg_time.FindAllAtring(inputData, -1)
	for _, match := range matchTime {
		EnteredTime := ConvertTime(match, inputData)
		inputData = strings.ReplaceAll(inputData, match, EnteredTime)
		colorToTerminal = strings.ReplaceAll(colorToTerminal, match, "\033[91m"+cityName+"\033[0m")
	}
	return trim_spaces(inputData), trim_spaces(colorToTerminal)
}
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

	width := 50   // distance to fly across the terminal
	fmt.Println() // start on a new line for the plane

	for i := 0; i < width; i++ {
		// Move cursor up to overwrite previous plane
		fmt.Print("\033[8A") // 8 lines tall
		for _, line := range plane {
			fmt.Print("\r" + strings.Repeat(" ", i) + line + "\n")
		}
		time.Sleep(120 * time.Millisecond) // control speed
	}

	fmt.Println() // clean line after animation
}
