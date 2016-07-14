package main

import "fmt"
import "os"
import "./sst"
import "io/ioutil"

func main() {
	source := "."
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	source += "/"
	calculateData(source)
	getIntervals(source)
	getLimits(source)
}

func calculateData(source string) {
	files, _ := ioutil.ReadDir(source)
	outlet, _ := os.Create(source + "sst.csv")
	analysis := sst.BeginAnalysis()
	defer outlet.Close()

	sst.Write(outlet, sst.BeginCSV())
	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			data := sst.AnalyzeSingle(sst.Read(source + file.Name()))
			analysis = sst.UpdateAnalysis(data, analysis)
			sst.Write(outlet, fmt.Sprintf("%s%s\n", file.Name(),
				                                    sst.FormatSingleCSV(data)))
		}
	}
	sst.Write(outlet, sst.FormatMultipleCSV(analysis))
}

func getIntervals(source string) {
	files, _ := ioutil.ReadDir(source)
	outlet, _ := os.Create(source + "intervals.csv")
	analysis := sst.BeginClock()
	defer outlet.Close()

	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			analysis = sst.UpdateClock(sst.ExtractIntervals(source + file.Name()),
		                               file.Name(), analysis)
		}
	}
	sst.Write(outlet, sst.FormatClock(analysis))
}

func getLimits(source string) {
	files, _ := ioutil.ReadDir(source)
	outlet, _ := os.Create(source + "clock.csv")
	analysis := sst.BeginGlobalClock()
	defer outlet.Close()

	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			xml := sst.ExtractGlobalClock(source + file.Name())
			analysis = sst.UpdateGlobalClock(analysis, file.Name(), xml)
		}
	}

	// TODO complete analysis by adding the test's final moment
	fmt.Printf("%#v\n", analysis)
}