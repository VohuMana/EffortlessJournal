package main

import (
	"fmt"
	"flag"
	"time"
	"encoding/json"
	// "log"
)

func main() {
	// Load configuration
	var (
		//configFile = flag.String("config", "", "Name of the config file that contains the client configuration")
	)
	flag.Parse()

	// config, err := LoadConfiguration(*configFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// twilio, err := NewTwilio(config.Twilio.Sid, config.Twilio.AuthToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = twilio.SendSMS(config.PhoneNumber, config.Twilio.PhoneNumber, "hello world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// xml, err := GenerateResponse("hello world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	journal := NewJournal("foo")
	journalEntry := NewJournalEntry(time.Now())
	qa1 := NewQuestionAnswerPair("How are you?", "Very good, thanks!")
	qa2 := NewQuestionAnswerPair("Did you have a good day?", "Yes")
	journalEntry.AddQuestionAnswer(qa1)
	journalEntry.AddQuestionAnswer(qa2)
	journal.AddEntry(journalEntry)
	
	json, err := json.Marshal(journal)
	if err == nil {
		fmt.Println(string(json))
	}
}