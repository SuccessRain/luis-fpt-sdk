package luis

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//Luis :
type Luis struct {
	appid  string
	client *Client
}

//NewLuis :
func NewLuis(key string, appid string) *Luis {
	if len(key) == 0 {
		return nil
	}

	e := new(Luis)
	e.appid = appid
	e.client = NewClient(key)
	return e
}

func (l *Luis) IntentList() ([]byte, *ErrorResponse) {
	url := getIntentListURL(l.appid)
	return l.client.Connect("GET", url, nil, true)
}

func (l *Luis) Predict(utterance string) ([]byte, *ErrorResponse) {
	url := getPredictURL(l.appid)
	data := getStringDataByteBuffer(utterance)
	return l.client.Connect("POST", url, data, true)
}

//CongVV
func (l *Luis) Train() ([]byte, *ErrorResponse) {
	url := trainAppNew(l.appid)//getTrainURL(l.appid)
	fmt.Print("URLLLLLLLLL:\t"); fmt.Println(url)
	return l.client.Connect("POST", url, nil, true)
}

//CongVV
func (l *Luis) TrainStatus() ([]byte, *ErrorResponse) {
	url := trainAppNew(l.appid)//getTrainURL(l.appid)
	fmt.Print("URLLLLLLLLL:\t"); fmt.Println(url)
	return l.client.Connect("GET", url, nil, true)
}

//CongVV
func (l *Luis) CreateIntent(userSays string) ([]byte, *ErrorResponse) {
	url := createIntentURL(l.appid)
	data := getNameByteBuffer(userSays)
	return l.client.Connect("POST", url, data, true)
}

//CongVV
func (client *Client) CreateApp(appName string, intents []Intent, utterances []Utterance) ([]byte, *ErrorResponse) {
	url := createAppURL(appName)
	var app AppJson = AppJson {
		Name: appName,
		Desc: appName,
		Culture: "en-us",
		Intents: intents,
		Entities: []string{},
		BingEntities: []string{},
		Actions:  []string{},
		ModelFeatures: []string{},
		RegexFeatures: []string{},
		Utterances: utterances,
	}
	bytExample, err := json.Marshal(app)
	if err != nil {

	}
	return client.Connect("POST", url, bytes.NewBuffer(bytExample), true)
}
