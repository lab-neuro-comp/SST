package main

import "fmt"
import "os"
import "io/ioutil"
import "./sst"

func main() {
	source := "."
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	source += "/"
	calculateData(source)
	getLimits(source)
}

/**
 * Studies the score obtained by that file set
 * @param source the source folder path
 */
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

/**
 * Analyzes the time performance of a data set
 * @param source the source folder path
 */
func getLimits(source string) {
	files, _ := ioutil.ReadDir(source)
	outClock, _ := os.Create(source + "clock.csv")
	outInt, _ := os.Create(source + "intervals.csv")
	clockInfo := sst.BeginGlobalClock()
	intervalsInfo := sst.BeginClock()
	defer outClock.Close()
	defer outInt.Close()

	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			where := source + file.Name()
			xml := sst.ExtractGlobalClock(where)
			intervals := sst.ExtractIntervals(where)
			clockInfo = sst.UpdateGlobalClock(clockInfo, file.Name(), xml)
			intervalsInfo = sst.UpdateClock(intervalsInfo, file.Name(), intervals)
		}
	}

	results := sst.MergeData(clockInfo, intervalsInfo)
	sst.Write(outClock, sst.FormatGlobalClock(results))
	sst.Write(outInt, sst.FormatClock(intervalsInfo))
}
