package luis

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func getFileByteBuffer(path string) (*bytes.Buffer, error) {
	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("File open err:", err)
		return nil, err
	}
	return bytes.NewBuffer(fileData), nil
}

func getUrlByteBuffer(url string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, url))
	return bytes.NewBuffer(byteData)

}

func getUserDataByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"userData":"%s"}`, userData))
	return bytes.NewBuffer(byteData)

}

func getStringDataByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`["%s"]`, userData))
	return bytes.NewBuffer(byteData)

}

//CongVV
func getNameByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"Name":"%s"}`, userData))
	return bytes.NewBuffer(byteData)
}

var s string = `{
  "name": "%s",
  "desc": "%s",
  "culture": "en-us",
  "intents": [
    {
      "Name": "%s",
    }
  ],
  "entities": [],
  "bing_entities": [],
  "actions": [],
  "model_features": [],
  "regex_features": [],
  "utterances": [
    {
      "text": "%s",
      "intent": "%s",
      "entities": []
    },
    {
      "text": "%s",
      "intent": "%s",
      "entities": []
    },
    {
      "text": "%s",
      "intent": "%s",
      "entities": []
    }
  ]
}`

//CongVV
func getImportAppDataBuffer(appName string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(s, appName, appName,"intentAhihi", "AAAA", "intentAhihi",
		"BBBB", "intentAhihi","CCCC", "intentAhihi"))
	return bytes.NewBuffer(byteData)
}

//ExampleJson :
type ExampleJson struct {
	ExampleText        string `json:"ExampleText"`
	SelectedIntentName string `json:"SelectedIntentName"`
	EntityLabels       []struct {
		EntityType string `json:"EntityType"`
		StartToken int    `json:"StartToken"`
		EndToken   int    `json:"StartToken"`
	} `json:"EntityLabels,omitempty"`
}

//CongVV
type Intent struct {
	Name string `json:"name"`
}

//CongVV
type Utterance struct {
	Text string `json:"text"`
	Intent string `json:"intent"`
	Entities []string `json:"entities"`
}

//CongVV
//AppJson :
type AppJson struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Culture string `json:"culture"`
	Intents       []Intent `json:"intents"`
	Entities []string `json:"entities"`
	BingEntities       []string `json:"bing_entities"`
	Actions       []string `json:"actions"`
	ModelFeatures       []string `json:"model_features"`
	RegexFeatures       []string `json:"regex_features"`
	Utterances       []Utterance `json:"utterances"`
}