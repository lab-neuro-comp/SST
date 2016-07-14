package sst

import "os"

/*****************
* MAIN FUNCTIONS *
*****************/
func BeginGlobalClock() map[string][]float64 {
	return make(map[string][]float64)
}

func ExtractGlobalClock(input string) string {
    inlet, _ := os.Open(input)
    // outlet := make([]float64, 0)

    // read variables
    defer inlet.Close()
    header := ReadHeader(inlet, getGlobalClockTags())
    records := ReadRecords(inlet, header)

    key := ""
    for k, _ := range header {
    	key = k
    	break
    }

    return records[key][0]
}

/*********************
* AUXILIAR FUNCTIONS *
*********************/
func getGlobalClockTags() []string {
    return []string {
        "Clock.Information",
    }
}