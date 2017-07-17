package main

import (
	"io/ioutil"
	"encoding/json"
)

type ClientConfiguration struct {
	Twilio Twilio `json:"twilio"`
}

type Twilio struct {
	Sid string `json:"sid"`
	AuthToken string `json:"authToken"`
}

func LoadConfiguration(filename string) (*ClientConfiguration, error) {
	var config ClientConfiguration
	fileContents, err := ioutil.ReadFile(filename)
	
	if err == nil {
		err = json.Unmarshal(fileContents, &config)
	}

	return &config, err
}