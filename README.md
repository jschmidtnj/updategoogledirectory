# updateprofile
This code edits gsuite profiles to contain newest information about employees.

## method
It uses three csv files, name configurable in the `config.toml` file, and parses them into JSON. Then this data is fed into the Google Directory API.

## download go

```
sudo apt-get upgrade && \
sudo apt-get update && \
sudo snap install --classic go && \
mkdir ~/go ~/go/bin ~/go/src ~/go/src/hello && \
export GOPATH=~/go && \
export PATH=$PATH:$GOPATH/bin && \
nano ~/go/src/hello/hello.go;

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}

sudo apt-get update && \
sudo apt-get install build-essential software-properties-common -y && \
sudo add-apt-repository ppa:ubuntu-toolchain-r/test -y && \
sudo apt-get update && \
sudo apt-get install gcc-snapshot -y && \
sudo apt-get update && \
sudo apt-get install gcc-6 g++-6 -y && \
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-6 60 --slave /usr/bin/g++ g++ /usr/bin/g++-6 && \
sudo apt-get install gcc-4.8 g++-4.8 -y && \
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 60 --slave /usr/bin/g++ g++ /usr/bin/g++-4.8;

go get ./…
go run *.go

```

## production
To get into production, change the service account private key located in the `config.toml` file. Then put the three csv files with the correct names (configurable) in the `data/` directory. Then run `go get ./...` and `go run *.go` in the `app/` directory. Make sure go is installed. Add cronjob for scheduling this task on Linux server (Ubuntu 16.4LTS tested).

## schema for the csv files:
### workloc: 
`Country Code, Location Name, Street, PO Box, City, State, Zip Code`  
### empinfo:
`personnel # (not needed), global personel # (use this), Full name, known-as, org number (orginfo.csv), Title, SAP-ID, manager flag, manager personel #, manager global personnel #, work location (desk) not accurate, phone #, country code, personel area, work location, active (not terminated) - A/T/I, etc., salary declarical`  
### orginfo:
`org-num, organization`  
