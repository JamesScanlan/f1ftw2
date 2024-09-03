package main

import (
	"f1ftw/f1ftw2/json_converter"
	"fmt"
)

func convertPredictions() {
	seasons := json_converter.ReadPredictionsJSON("data/Predictions.json")
	fmt.Println(len(seasons.Seasons))
	new_seasons := json_converter.ConvertToNewPredictions(seasons)
	json_converter.WritePredictions("data/NewPredicions.json", new_seasons)
	json_converter.WritePredictions("data/2024Prections.json", json_converter.GetCurrentYearPredictions(new_seasons))
}

func convertGrandPrix() {
	originalGrandPrixSeasons := json_converter.ReadGPData("data/OriginalGPData.json")

	// // Display Parsed GP Data
	// for i := 0; i < len(originalGrandPrixSeasons.GrandPrixSeasons); i++ {
	// 	season := originalGrandPrixSeasons.GrandPrixSeasons[i]
	// 	fmt.Println(season.Year)
	// 	for raceIndex := 0; raceIndex < len(season.Races); raceIndex++ {
	// 		race := season.Races[raceIndex]
	// 		fmt.Println("\t" + race.Name)
	// 		for resultIndex := 0; resultIndex < len(race.Results); resultIndex++ {
	// 			result := race.Results[resultIndex]
	// 			fmt.Println("\t\t" + result.Driver + " " + result.Position)
	// 		}
	// 	}
	// }

	json_converter.WriteGPData("data/NewGPData.json", json_converter.ConvertToNewGPData(originalGrandPrixSeasons))
}

func main() {
	convertGrandPrix()
}
