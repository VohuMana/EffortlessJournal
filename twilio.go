package main

import (
	"fmt"
	"net/http"
	"net/url"
	"bytes"
	"io/ioutil"
	"encoding/xml"
)

type TwilioApi struct {
	baseUrl *url.URL
	accountSid string
	authToken string
}

type Response struct {
	Message string `xml:"Message"`
}

type IncomingMessage struct {
	MessageSid string
	AccountSid string
	MessagingServiceSid string
	From string
	To string
	Body string
	NumMedia int
}

func NewTwilio(accountSid, authToken string) (*TwilioApi, error) {
	var twilio *TwilioApi
	url, err := url.Parse("https://api.twilio.com")
	if err == nil {
		twilio = &TwilioApi{
			baseUrl: url,
			accountSid: accountSid,
			authToken: authToken,
		}
	}
	
	return twilio, err
}

func (t *TwilioApi) SendSMS(to, from, message string) error {
	// Construct the SMS URL
	smsUrl := t.baseUrl
	smsUrl.Path = fmt.Sprintf("2010-04-01/Accounts/%s/Messages.json", t.accountSid)

	// Construct URL encoded body
	urlBody := &url.Values{}
	urlBody.Add("To", to)
	urlBody.Add("From", from)
	urlBody.Add("Body", message)

	// Create the request
	request, err := http.NewRequest("POST", smsUrl.String(), bytes.NewBufferString(urlBody.Encode()))

	// Set the username and password
	request.SetBasicAuth(t.accountSid, t.authToken)

	// Set the URL encoded data (to, from, message)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request and parse the result
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err == nil {
		if response.Body != nil {
			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(responseBody)
			}
			defer response.Body.Close()
		}
	}
	
	return err
}

func ParseIncomingMessage(postBody string) (*IncomingMessage, error) {
	return nil, nil
}

func GenerateResponse(message string) (string, error) {
	// <?xml version="1.0" encoding="UTF-8"?>
	// <Response>
    // 	<Message>
    //     	Hello, Mobile Monkey
    // 	</Message>
	// </Response>
	
	resp := Response{
		Message:message,
	}

	var xmlEncodedString string
	output, err := xml.Marshal(resp)
	if err == nil {
		xmlEncodedString = xml.Header + string(output)
	}

	return xmlEncodedString, err
}