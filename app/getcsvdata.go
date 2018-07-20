package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readFiles(file string) [][]string {
	csvFile, err := os.Open("data/" + file)
	if err != nil {
		fmt.Printf("error with opening file: %v\n", err)
		Log.Printf("error with opening file: %v\n", err)
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';' //Use semicolon-delimited instead of comma <---- here!
	reader.LazyQuotes = true
	allrecords := [][]string{}
	//get the data from buffer
	for {
		record, err := reader.Read()
		//stop at end of file
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error with reading file: %v\n", err)
			Log.Printf("error with reading file: %v\n", err)
			panic(err)
		}
		//record is an array of values
		//fmt.Println(record[0])
		allrecords = append(allrecords, record)
	}
	return allrecords
}

func getDataCsv() {
	values := readFiles(getConfig.Workloc_File)
	//fmt.Println(values)
	values = readFiles(getConfig.Empinfo_File)
	//fmt.Println(values)
	values = readFiles(getConfig.Orginfo_File)
	fmt.Println(values)
}
