package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	VERSION = "1.0.0"
)

// imports the config.go file immediately, setting the constants to the getConfig variable, as defined in the config/config.toml file
var getConfig = Config{}

//the first function that runs, and initializes the program
func init() {
	//first get the configuration
	getConfig.Read()
	connectToApi()
	flag.Parse()
	//filelocation := "logs/log_" + time.Now().Format(time.RFC3339)
	var logpath = flag.String("logpath", "logs/main.log", "Log Path")
	NewLog(*logpath)
	//Log.SetOutput(f)
	Log.Println("\n-----------------------------------------------------------------------------------------------------------------------")
	Log.Printf("Server v%s pid=%d started with processes: %d\n", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println("Created Log File")
}

//the second function that runs, the main function, necessary in golang programs
func main() {
	go frontEndApp()
	runDirectoryUpdate()
}

//This function runs the google directory api and updates contacts from csv file info
func runDirectoryUpdate() {
	t := time.Now()
	//Log logs to a file
	fmt.Printf("Starting Directory Update, %s\n", t.Format("2006-01-02 15:04:05"))
	Log.Printf("Starting Directory Update, %s\n", t.Format("2006-01-02 15:04:05"))

	if getConfig.Debug_Mode {
		fmt.Println("DEBUG MODE ON")
		Log.Println("Debug Mode On")
	} else {
		fmt.Println("DEBUG MODE OFF")
		Log.Println("Debug Mode Off")
	}

	//get the data from the csv files
	getJsonData()
	//getDataCsv() //alternative using the csv file directly, without conversion to json

	//get country code data
	getCountryData()

	//then configure the data to be useable
	configuredata()

	//then send the data off with the api request call
	apirequest()
}
