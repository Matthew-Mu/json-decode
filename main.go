package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type weather struct {
	Date,
	Day,
	TempLow,
	TempHigh,
	Rain,
	Evap,
	Sun,
	WindDir,
	MaxWindSpd,
	Time,
	AMTemp,
	AMRH,
	AMCloud,
	AMWindDir,
	AMMaxWindSpd,
	AMPressure,
	PMTemp,
	PMRH,
	PMCloud,
	PMWindDir,
	PMMaxWindSpd,
	PMPressure string
}

type weatherTable struct {
	table []weather
}

func main() {

	jsonBytes, err := os.ReadFile("../web-scraper/BOMweather.json")
	if err != nil {
		fmt.Println(err)
	}

	var structArray []weather
	error := json.Unmarshal([]byte(jsonBytes), &structArray)
	if error != nil {
		fmt.Println(error)
	}
	/*wTable := weatherTable{
		table: structArray,
	}*/

	for i := 0; i < len(structArray); i++ {
		fmt.Println(structArray[i].Date, "-", structArray[i].Day, structArray[i].TempLow, structArray[i].TempHigh)
		fmt.Println(structArray[i])
	}
}
