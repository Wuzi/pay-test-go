package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wuzi/pay-test-go/models"
	"github.com/Wuzi/pay-test-go/restapi/operations/city"
	"github.com/go-openapi/runtime/middleware"
)

// Cities is our in-memory "database"
var Cities = []*models.City{}

// LoadCities reads city_list.json file and loads into memory
func LoadCities() {
	jsonFile, err := os.Open("city_list.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully opened city_list.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Cities)
}

// GetCitiesHandlerFunc returns an array of cities
func GetCitiesHandlerFunc(params city.GetCitiesParams) middleware.Responder {
	return city.NewGetCitiesOK().WithPayload(Cities)
}

// GetCitiesIDHandlerFunc returns a single city by id
func GetCitiesIDHandlerFunc(params city.GetCitiesIDParams) middleware.Responder {
	for _, c := range Cities {
		if c.ID == params.ID {
			return city.NewGetCitiesIDOK().WithPayload(c)
		}
	}
	return city.NewGetCitiesIDNotFound()
}
