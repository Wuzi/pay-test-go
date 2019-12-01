package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wuzi/pay-test-go/models"
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
