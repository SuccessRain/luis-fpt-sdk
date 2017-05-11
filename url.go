package luis

const (
	//LuisURL :Basic URL
	LuisURL string = "https://api.projectoxford.ai/luis/v1.0/prog/apps/"
	//LuisAPIIntents :API Intent List
	LuisAPIIntents string = "intents"
	//LuisAPIActionChannels :API Action Channels
	LuisAPIActionChannels string = "actionChannels"
	//LuisAPIPredict :API Predict
	LuisAPIPredict string = "predict"
	//LuisAPITrain :API Train
	LuisAPITrain string = "train"
	//LuisAPIAddExample :API Add Label
	LuisAPIAddExample string = "example"

	//CongVV
	LuisCreateItent string = "https://westus.api.cognitive.microsoft.com/luis/v1.0/prog/apps/"
	LuisCreateApp string = "https://westus.api.cognitive.microsoft.com/luis/v1.0/prog/apps/import?appName="
	LuisTrainNew string = "https://westus.api.cognitive.microsoft.com/luis/v1.0/prog/apps/"
)

func getIntentListURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPIIntents
}

func getActionChannels(apid string) string {
	return LuisURL + apid + "/" + LuisAPIActionChannels
}

func getPredictURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPIPredict
}

func getTrainURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPITrain
}

func getAddExampleURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPIAddExample
}

func createIntentURL(apid string) string {
	return LuisCreateItent + apid + "/" + LuisAPIIntents
}

func createAppURL(apName string) string {
	return LuisCreateApp + apName
}

func trainAppNew(apid string) string {
	return LuisTrainNew + apid + "/" + LuisAPITrain
}