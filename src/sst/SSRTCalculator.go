package sst

// Creates a slice containing the needed variables for analysis.
func GetNeededVariables() []string {
	return []string {
		"PressStimulus.RT",
		"PressStimulus.ACC",
		"SoundStimulus.RT",
		"SoundStimulus.ACC",
		"VisualStimulus.Duration",
		"Procedure[Trial]",
		"Subject",
		"Session",
	}
}

// Creates a map to be used on single and multiple file analysis.
func CreateAnalysisMap() map[string]float64 {
	result := make(map[string]float64)

	result["RT"] = 0
	result["SSD"] = 0
	result["correctGo"] = 0
	result["totalGo"] = 0
	result["correctStop"] = 0
	result["totalStop"] = 0
	result["SSRT"] = 0
	result["%INHIB"] = 0
	result["%AUS"] = 0
	result["GENERAL"] = 0
	result["SUBJECT"] = 0
	result["SESSION"] = 0

	return result
}

// Performs a study process on one SST data set, by using a map relating the
// needed fields to arrays of the collected data. Returns a map relating the
// desided measures with their respective calculations.
func AnalyzeSingle(data map[string][]string) map[string]float64 {
	result := CreateAnalysisMap()
	limit := len(data["Procedure[Trial]"])

	// Extracting data
	for i := 0; i < limit; i++ {
		switch data["Procedure[Trial]"][i] {
		case "PressProc":
			result = studyPressProc(i, data, result)
		case "NotPressProc":
			result = studyNotPressProc(i, data, result)
		}
	}

	// Subject identification
	result["SUBJECT"] = ParseFloat64(data["Subject"][0])
	result["SESSION"] = ParseFloat64(data["Session"][0])

	// Subject score
	result["RT"] /= result["totalGo"]
	result["SSD"] /= result["correctStop"]
	result["SSRT"] = result["RT"] - result["SSD"]
	result["%INHIB"] = 100*(result["correctStop"] / result["totalStop"])
	result["%AUS"] = 100*(result["totalGo"]-result["correctGo"])/result["totalGo"]
	result["GENERAL"] = ((result["correctGo"] + result["correctStop"]*(1 - result["%AUS"]/100)) /
	                    (result["totalGo"] + result["totalStop"]))

	return result
}

func studyPressProc(where int,
	                data map[string][]string,
					stuff map[string]float64) map[string]float64 {
	reactionTime := ParseFloat64(data["PressStimulus.RT"][where])
	reaction := ParseFloat64(data["PressStimulus.ACC"][where])

	stuff["RT"] += reactionTime
	stuff["correctGo"] += reaction
	stuff["totalGo"]++

	return stuff
}

func studyNotPressProc(where int,
	                   data map[string][]string,
					   stuff map[string]float64) map[string]float64 {
	reactionTime := ParseFloat64(data["VisualStimulus.Duration"][where])
	reaction := ParseFloat64(data["SoundStimulus.ACC"][where])

	if reaction == 0 {
		reactionTime = 0.0
	}

	stuff["SSD"] += reactionTime
	stuff["correctStop"] += reaction
	stuff["totalStop"]++

	return stuff
}
