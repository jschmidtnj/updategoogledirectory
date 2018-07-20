package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

var (
	jsondata             []uint8
	worklocationdata     []Worklocation
	employeeinfodata     []Employeeinfo
	organizationinfodata []Organizationinfo
)

type Worklocation struct {
	Entry1 string `json:"worklocation"`    //worklocation    //work location (used to connect to empinfo.csv data) (not used by google)
	Entry2 string `json:"extendedAddress"` //extendedAddress //Location Name
	Entry3 string `json:"streetAddress"`   //streetAddress   //Street
	Entry4 string `json:"poBox"`           //poBox           //po-box
	Entry5 string `json:"locality"`        //locality        //city
	Entry6 string `json:"region"`          //region          //state
	Entry7 string `json:"postalCode"`      //postalCode      //zip code
}

type Employeeinfo struct {
	Entry1  string `json:"personnelnumber"`              //personnelnumber              //personnel # - not used ever (not used by google)
	Entry2  string `json:"externalid"`                   //externalid                   //global personnel #
	Entry3  string `json:"fullName"`                     //fullName                     //full name
	Entry4  string `json:"givenName"`                    //givenName                    //nick-name
	Entry5  string `json:"organizationnumber"`           //organizationnumber           //organization number (used to connect to orginfo.csv data) (not used by google)
	Entry6  string `json:"title"`                        //title                        //title - can possibly add to full name
	Entry7  string `json:"sapid"`                        //sapid                        //SAP ID number (not used by google)
	Entry8  string `json:"managerflag"`                  //managerflag                  //Y or N for yes or no to show if you are a manager (not used by google)
	Entry9  string `json:"managerpersonnelnumber"`       //managerpersonnelnumber       //number as manager - not used ever (not used by google)
	Entry10 string `json:"managerglobalpersonnelnumber"` //managerglobalpersonnelnumber //global personnel number (not used by google)
	Entry11 string `json:"deskCode"`                     //deskCode                     //your desk number
	Entry12 string `json:"phonenumber"`                  //phonenumber                  //phone number
	Entry13 string `json:"countryCode"`                  //countryCode                  //country code
	Entry14 string `json:"personnelarea"`                //personnelarea                //personnel area (not used by google)
	Entry15 string `json:"worklocation"`                 //worklocation                 //work location (not used by google) (used to connect to workloc.csv data)
	Entry16 string `json:"status"`                       //status                       //status at company (active A, Terminated T, Intern I) (not used by google)
	Entry17 string `json:"salarydeclerical"`             //salarydeclerical             //salary declarical (yes Y, no N) (not used by google)
}

type Organizationinfo struct {
	Entry1 string `json:"worklocation"`     //worklocation     //organization number (used to connect to empinfo.csv data) (not used by google)
	Entry2 string `json:"organizationname"` //organizationname //organization name (add to organizations - name object)
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func csvtojson(file string) []uint8 {
	//read from csv file, convert to json, send to google directory api, do for all three csv files
	csvFile, err := os.Open("data/" + file + ".csv")
	if err != nil {
		fmt.Printf("error with opening worklog file: %v\n", err)
		Log.Printf("error with opening worklog file: %v\n", err)
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';' //Use semicolon-delimited instead of comma <---- here!
	reader.LazyQuotes = true

	csvData, err := reader.ReadAll()
	if err, ok := err.(*csv.ParseError); ok && err.Err == csv.ErrFieldCount {
		fmt.Printf("error with reading workloc file: %v\n", err)
		Log.Printf("error with reading workloc file: %v\n", err)
		return nil
	}

	if file == getConfig.Workloc_File {
		// type assert on it
		var oneRecord Worklocation
		var allRecords []Worklocation
		//fmt.Println("parsed Workloc records")
		for _, each := range csvData {
			oneRecord.Entry1 = each[0]
			oneRecord.Entry2 = each[1]
			oneRecord.Entry3 = each[2]
			oneRecord.Entry4 = each[3]
			oneRecord.Entry5 = each[4]
			oneRecord.Entry6 = each[5]
			oneRecord.Entry7 = each[6]
			allRecords = append(allRecords, oneRecord)
		}
		jsondata, err = json.Marshal(allRecords) // convert to JSON
	} else if file == getConfig.Empinfo_File {
		// type assert on it
		var oneRecord Employeeinfo
		var allRecords []Employeeinfo
		//fmt.Println("parsed Workloc records")
		for _, each := range csvData {
			oneRecord.Entry1 = each[0]
			oneRecord.Entry2 = each[1]
			oneRecord.Entry3 = each[2]
			oneRecord.Entry4 = each[3]
			oneRecord.Entry5 = each[4]
			oneRecord.Entry6 = each[5]
			oneRecord.Entry7 = each[6]
			oneRecord.Entry8 = each[7]
			oneRecord.Entry9 = each[8]
			oneRecord.Entry10 = each[9]
			oneRecord.Entry11 = each[10]
			oneRecord.Entry12 = each[11]
			oneRecord.Entry13 = each[12]
			oneRecord.Entry14 = each[13]
			oneRecord.Entry15 = each[14]
			oneRecord.Entry16 = each[15]
			oneRecord.Entry17 = each[16]
			allRecords = append(allRecords, oneRecord)
		}
		jsondata, err = json.Marshal(allRecords) // convert to JSON
	} else if file == getConfig.Orginfo_File {
		// type assert on it
		var oneRecord Organizationinfo
		var allRecords []Organizationinfo
		//fmt.Println("parsed Workloc records")
		for _, each := range csvData {
			oneRecord.Entry1 = each[0]
			oneRecord.Entry2 = each[1]
			allRecords = append(allRecords, oneRecord)
		}
		//fmt.Printf(allRecords[0].Entry1)
		jsondata, err = json.Marshal(allRecords) // convert to JSON
	}
	//jsondata, err := JSONMarshal(allRecords)
	//enc := json.NewEncoder(os.Stdout)
	//enc.SetEscapeHTML(false)
	//err = enc.Encode(allRecords)
	if err != nil {
		fmt.Printf("error with type converting to json: %v\n", err)
		Log.Printf("error with type converting to json: %v\n", err)
		panic(err)
	}

	//fmt.Println(reflect.TypeOf(jsondata))
	//create json file
	jsonFile, err := os.Create("data/" + file + ".json")
	if err != nil {
		fmt.Printf("error with type converting to json: %v\n", err)
		Log.Printf("error with type converting to json: %v\n", err)
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsondata)
	//jsonFile.Write([]byte("asdf"))
	jsonFile.Close()
	return jsondata
}

func getJsonData() {
	//saves as json unmarshalled object
	err := json.Unmarshal([]byte(csvtojson(getConfig.Workloc_File)), &worklocationdata)
	if err != nil {
		fmt.Printf("error with type converting to employeeinfodata json: %v\n", err)
		Log.Printf("error with type converting to employeeinfodata json: %v\n", err)
		panic(err)
	}
	//fmt.Println(worklocationdata[0].Entry1)
	err = json.Unmarshal(csvtojson(getConfig.Empinfo_File), &employeeinfodata)
	if err != nil {
		fmt.Printf("error with type converting to employeeinfodata json: %v\n", err)
		Log.Printf("error with type converting to employeeinfodata json: %v\n", err)
		panic(err)
	}
	//fmt.Println(employeeinfodata[0].Entry1)
	err = json.Unmarshal(csvtojson(getConfig.Orginfo_File), &organizationinfodata)
	if err != nil {
		fmt.Printf("error with type converting to organizationinfodata json: %v\n", err)
		Log.Printf("error with type converting to organizationinfodata json: %v\n", err)
		panic(err)
	}
	//fmt.Println(organizationinfodata[0].Entry1)
}
