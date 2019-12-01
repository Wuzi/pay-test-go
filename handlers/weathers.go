package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wuzi/pay-test-go/models"
)

// Weathers is our in-memory "database"
var Weathers = []*models.Weather{}

// LoadWeathers reads weather_list.json file and loads into memory
func LoadWeathers() {
	jsonFile, err := os.Open("weather_list.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully opened weather_list.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Cities)
}
