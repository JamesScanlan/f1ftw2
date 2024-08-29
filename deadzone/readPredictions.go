package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type PredictionYears struct {
	Years []Predictions `json:"Grands_Prix"`
}

type Predictions struct {
	Year  int    `json:"Year"`
	Races []Race `json:"Races"`
}

type Race struct {
	GrandPrix   string       `json:"Grand_Prix"`
	Predictions []Prediction `json:"Predictions"`
}

type Prediction struct {
	Predictor   string `json:"Predictor"`
	Qualifying  string `json:"Qualifying"`
	Race        string `json:"Race"`
	Progression string `json:"Progression"`
}

func loadPredictionsData(filename string) PredictionYears {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var predictionYears PredictionYears

	json.Unmarshal(byteValue, &predictionYears)

	return predictionYears
}

func main() {
	predictionYears := loadPredictionsData("Predictions.json")

	//fmt.Println(len(predictionYears.Years))
	// fmt.Println(predictionYears.Years[1].Year)
	for i := range predictionYears.Years {

		if predictionYears.Years[i].Year == 2024 {
			fmt.Println(predictionYears.Years[i].Year)

			// fmt.Println(len(predictionYears.Years[i].Races))
			for j := range predictionYears.Years[i].Races {
				fmt.Println(predictionYears.Years[i].Races[j].GrandPrix)

				for k := range predictionYears.Years[i].Races[j].Predictions {
					fmt.Println(predictionYears.Years[i].Races[j].Predictions[k].Predictor)
					fmt.Println("\t" + predictionYears.Years[i].Races[j].Predictions[k].Qualifying)
					fmt.Println("\t" + predictionYears.Years[i].Races[j].Predictions[k].Race)
					fmt.Println("\t" + predictionYears.Years[i].Races[j].Predictions[k].Progression)

				}
			}
		}

	}
}
