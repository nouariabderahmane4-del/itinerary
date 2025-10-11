package processor

import (
	"bufio"
	"fmt"
	"input"
	"os"
	"regexp"
	"strings"
	"time"
)

func TrimSpace(text string) string{
	text = strings.NewReplacer("\r", "\n", "\v", "\n", "\f", "\n").Replace(text)
	text = regexp.MustCompile(`[ ]{2,}`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
	text= strings.TrimSpace(text)
	return text
}

func Input_analyzing(inputData string, csvData []input.Airport) (string, string) {

	colorToTerminal := inputData

	reg_city := regexp.MustCompile(`\*#+[A-Z]{3,4}\b`)
	reg_iata := regexp.MustCompile(`\#[A-Z]{3}\b`)
	reg_icao := regexp.MustCompile(`\##[A-Z]{4}\b`)
	reg_time := regexp.MustCompile(`(D|T12|T24)(\([^)]+\))`)

	matchCity := reg_city.FindAllAtring(inputData, -1)
	for _, match := range matchCity{
		cityName := replace_Municipality(match, csvData)
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