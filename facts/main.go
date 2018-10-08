package main

import (
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(requestEnvelope RequestEnvelope) (ResponseEnvelope, error) {
	outputText := factForLocale(requestEnvelope.Request.Locale)

	responseEnvelope := ResponseEnvelope{
		Version: "1.0",
		Response: Response{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: outputText,
			},
		},
	}

	return responseEnvelope, nil
}

var factsDict = map[string][]string{
	"en-US": []string{
		"In the United States, the sport known as soccer is known as football.",
		`Soccer is a British term appearing in the 1880s 
		as an abbreviation for the word association, as in association football,  
		a formal name used to distinguish itself from rugby football.`,
	},
	"en-GB": []string{
		"In the United States, the sport known as football is known as soccer.",
		`Soccer is a British term appearing in the 1880s 
		as an abbreviation for the word association, as in association football,  
		a formal name used to distinguish itself from rugby football.`,
	},
}

func factForLocale(locale string) string {
	if _, exists := factsDict[locale]; exists {
		return randomFact(factsDict[locale])
	} else {
		return randomFact(factsDict["en-US"])
	}
}

func randomFact(facts []string) string {
	rand.Seed(time.Now().UnixNano())
	return facts[rand.Intn(len(facts))]
}

// Boilerplate JSON mappings
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html
type RequestEnvelope struct {
	Version string  `json:"version"`
	Request Request `json:"request"`
}

type Request struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	Locale      string `json:"locale"`
	Intent      Intent `json:"intent,omitempty"`
	Reason      string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

type Intent struct {
	Name string `json:"name"`
}

type ResponseEnvelope struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Ssml string `json:"ssml,omitempty"`
}
