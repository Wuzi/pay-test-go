package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wuzi/pay-test-go/models"
	"github.com/Wuzi/pay-test-go/restapi/operations/weather"
	"github.com/go-openapi/runtime/middleware"
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
	json.Unmarshal(byteValue, &Weathers)
}

// GetCitiesIDWeathersHandlerFunc returns the wethers of a city by id
func GetCitiesIDWeathersHandlerFunc(params weather.GetCitiesIDWeathersParams) middleware.Responder {
	var weathers []*models.WeatherData
	var foundCity = false

	for _, c := range Cities {
		if c.ID == params.ID {
			foundCity = true
		}
	}

	if foundCity == false {
		return weather.NewGetCitiesIDWeathersNotFound()
	}

	for _, w := range Weathers {
		if w.CityID == params.ID {
			weathers = w.Data
		}
	}

	return weather.NewGetCitiesIDWeathersOK().WithPayload(weathers)
}
