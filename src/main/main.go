package main

import "fmt"
import "os"
import "sync"
import "io/ioutil"
import "github.com/lab-neuro-comp/SST/src/sst"

// This is the main application to be used to process the generated data from
// the Stop Signal Test implemented on E-Prime.
func main() {
	var wg sync.WaitGroup
	var source string

	if len(os.Args) > 1 {
		source = os.Args[1]
	} else {
		source = "."
	}
	source += "/"

	wg.Add(2)
	go calculateData(source, &wg)
	go getLimits(source, &wg)
	wg.Wait()
}

// Studies the score obtained by that file set on the source folder. It is made
// to be executed on parallel, as defined by the determined wait group.
func calculateData(source string, wg *sync.WaitGroup) {
	files, _ := ioutil.ReadDir(source)
	outlet, _ := os.Create(source + "sst.csv")
	analysis := sst.BeginAnalysis()
	defer wg.Done()
	defer outlet.Close()

	sst.Write(outlet, sst.BeginCSV())
	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			data := sst.AnalyzeSingle(sst.Read(source + file.Name()))
			analysis = sst.UpdateAnalysis(data, analysis)
			sst.Write(outlet,
				      fmt.Sprintf("%s%s\n",
				      	          file.Name(),
				      	          sst.FormatSingleCSV(data)))
		}
	}
	sst.Write(outlet, sst.FormatMultipleCSV(analysis))
}

// Analyzes the time performance of a data set. It is made to be executed on
// parallel, as defined by the determined wait group.
func getLimits(source string, wg *sync.WaitGroup) {
	files, _ := ioutil.ReadDir(source)
	outClock, _ := os.Create(source + "clock.csv")
	outInts, _ := os.Create(source + "intervals.csv")
	timer := sst.BeginTimer()
	stopwatch := sst.BeginStopwatch()
	ids := make(map[string]string)
	defer wg.Done()
	defer outClock.Close()
	defer outInts.Close()

	for _, file := range files {
		if sst.ValidFile(file.Name()) {
			where := source + file.Name()
			// Clock information
			data := sst.ExtractTimer(where)
			timer = sst.UpdateTimer(timer, file.Name(), data)
			ids[file.Name()] = sst.GetId(data)
			// Intervals information
			intervals := sst.ExtractIntervals(where)
			stopwatch = sst.UpdateStopwatch(stopwatch, file.Name(), intervals)
		}
	}

	results := sst.MergeData(timer, stopwatch)
	sst.Write(outClock, sst.FormatTimer(results))
	sst.Write(outInts, sst.FormatStopwatch(stopwatch, ids))
}
