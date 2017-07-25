package main

import (
	"io/ioutil"
	"encoding/json"
)

type ClientConfiguration struct {
	PhoneNumber string `json:"phoneNumber"`
	Twilio Twilio `json:"twilio"`
}

type Twilio struct {
	Sid string `json:"sid"`
	AuthToken string `json:"authToken"`
	PhoneNumber string `json:"phoneNumber"`
}

func LoadConfiguration(filename string) (*ClientConfiguration, error) {
	var config ClientConfiguration
	fileContents, err := ioutil.ReadFile(filename)
	
	if err == nil {
		err = json.Unmarshal(fileContents, &config)
	}

	return &config, err
}