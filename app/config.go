package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Service_Account_Email       string
	Service_Account_Private_Key string
	Service_Account_Scopes      []string
	Admin_Email                 string
	Token_Refresh_Count         int
	Frontend_Port               int
	Backend_Port                string
	Workloc_File                string
	Empinfo_File                string
	Orginfo_File                string
	Test_String                 string
	Print_Status_Every          int
	Debug_Mode                  bool
	Update_Single_Account       string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config/config.toml", &c); err != nil {
		Log.Printf("problem with config file %v\n", err)
		fmt.Printf("problem with config file %v\n", err)
		panic(err)
	}
}
