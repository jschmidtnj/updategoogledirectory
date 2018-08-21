package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type country struct {
	Name       string `json:"Name"`
	Alpha2code string `json:"Alpha2code"`
}

var (
	countrydata []country
)

func getCountryData() {
	resp, err := http.Get(getConfig.Country_Data_Url)
	if err != nil || resp.StatusCode != http.StatusOK {
		Log.Printf("Unable to retrieve country data rest: %v", err)
		fmt.Printf("Unable to retrieve country data rest: %v", err)
		countryfile, err := ioutil.ReadFile("data/" + getConfig.Country_Data_File + ".json")
		if err != nil {
			Log.Printf("Unable to retrieve country data file: %v", err)
			fmt.Printf("Unable to retrieve country data file: %v", err)
			panic(err)
		}
		err = json.Unmarshal(countryfile, &countrydata)
		if err != nil {
			Log.Printf("Unable to parse country data file: %v", err)
			fmt.Printf("Unable to parse country data file: %v", err)
			panic(err)
		}
	} else {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error with get request: %v\n", err)
			Log.Printf("error with get request: %v\n", err)
			panic(err)
		}
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)
		jsonFile, err := os.Create("data/" + getConfig.Country_Data_File + ".json")
		if err != nil {
			fmt.Printf("error with type converting to json: %v\n", err)
			Log.Printf("error with type converting to json: %v\n", err)
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &countrydata)
		if err != nil {
			Log.Printf("Unable to parse country data resp: %v", err)
			fmt.Printf("Unable to parse country data resp: %v", err)
			panic(err)
		}
		defer jsonFile.Close()
		jsonData, err := json.Marshal(countrydata)
		if err != nil {
			Log.Printf("Unable to marshal country data: %v", err)
			fmt.Printf("Unable to marshal country data: %v", err)
			panic(err)
		}
		jsonFile.Write(jsonData)
		jsonFile.Close()
	}
}
