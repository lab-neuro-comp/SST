package sst

import "fmt"
import "os"

// Get the specified fields that will be used during the data analysis
func GetAnalysisParameters() []string {
	return []string {
		"totalGo",
		"correctGo",
		"totalStop",
		"correctStop",
		"%INHIB",
		"%AUS",
		"GENERAL",
		"SSD",
		"RT",
		"SSRT",
		"SUBJECT",
		"SESSION",
	}
}

// Gets the output parameters that will populate the results' table
func GetResultParameters() []string {
	return []string {
		"SUBJECT",
		"SESSION",
		"RT",
		"SSD",
		"SSRT",
		"%INHIB",
		"%AUS",
		"GENERAL",
	}
}

//  Generates an analysis structure. Consists of a map relating a parameter to
// an array of numbers. The needed parameters can be obtained through the
// "GetAnalysisParameters() string" function.
func BeginAnalysis() map[string][]float64 {
	data := make(map[string][]float64)

	for _, param := range GetAnalysisParameters() {
		data[param] = make([]float64, 0)
	}

	return data
}

// Updates a multiple analysis structure by increment the outlet structure
// depending on the data presented on the inlet structure.
func UpdateAnalysis(inlet map[string]float64,
	                outlet map[string][]float64) map[string][]float64 {
	for _, param := range GetAnalysisParameters() {
		outlet[param] = append(outlet[param], inlet[param])
	}

	return outlet
}

// Calls the calculations functions upon the analysis structure: the arrays
// containing the results of the individual analysis. Returns a map relating
// the requested parameters and strings of results.
func EndAnalysis(analysis map[string][]float64) map[string]string {
	outlet := make(map[string]string)

	for key, value := range analysis {
		mean := Mean(value)
		dev := StdDev(value)
		outlet[key] = fmt.Sprintf("%s: %3f +- %3f", key, mean, dev)
	}

	return outlet
}

// Formats the output of an analysis structure.
func FormatSingle(data map[string]float64) (box string) {
	for _, param := range GetAnalysisParameters() {
		box += fmt.Sprintf("%s: %3f\n", param, data[param])
	}
	return
}

// Formats the output of the complete analysis structure.
func FormatMultiple(inlet map[string]string) (box string) {
	for _, item := range GetResultParameters() {
		box += fmt.Sprintf("%s\n", inlet[item])
	}

	return
}

// Writes data of an analysis structure to a stream. If the outlet pointer is
// nil, it will write to the standard output.
func Write(outlet *os.File, data string) {
	if outlet == nil {
		fmt.Printf("%s", data)
	} else {
		outlet.WriteString(data)
	}
}

/* ------------------------------------------------------------ */

// Generates a header for a CSV output.
func BeginCSV() string {
	box := "File name"

	for _, it := range GetAnalysisParameters() {
		box += fmt.Sprintf(";%s", it)
	}

	return box + "\n"
}

// Formats the output of an analysis structure to the CSV format.
// It assumes there is a file name so this string can be be concatenated
// to its end.
func FormatSingleCSV(data map[string]float64) (box string) {
	for _, param := range GetAnalysisParameters() {
		box += replaceInString(fmt.Sprintf(";%3f", data[param]), '.', ',')
	}
	return
}

// Calls the calculations functions upon the analysis structure;
func FormatMultipleCSV(analysis map[string][]float64) string {
	outlet := "Resultado"

	for _, param := range GetAnalysisParameters() {
		mean := Mean(analysis[param])
		dev := StdDev(analysis[param])
		outlet += replaceInString(fmt.Sprintf(";%3f +- %3f", mean, dev), '.', ',')
	}

	return fmt.Sprintf("%s\n", outlet)
}
