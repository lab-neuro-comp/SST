package sst

import "os"
import "strings"

/*****************
* MAIN FUNCTIONS *
*****************/
func BeginGlobalClock() map[string]string {
	return make(map[string]string)
}

func ExtractGlobalClock(input string) string {
    inlet, _ := os.Open(input)
    // outlet := make([]float64, 0)

    // read variables
    defer inlet.Close()
    header := ReadHeader(inlet, getGlobalClockTags())
    records := ReadRecords(inlet, header)

    // further analysis
    key := ""
    for k, _ := range header {
    	key = k
    	break
    }
    xml := records[key][0]
    lower, upper := FindBorders(xml, "DateUtc")

    return xml[lower:upper]
}

func UpdateGlobalClock(analysis map[string]string, 
	                   input, data string) map[string]string {
	analysis[input] = data
	return analysis
}

/****************
* XML FUNCTIONS *
****************/
func FindBorders(text, tag string) (int, int) {
	lower := strings.Index(text, tag)
	for text[lower]	!= '>' {
		lower++
	}

	upper := 1 + lower
	for text[upper] != '<' {
		upper++
	}

	lower++
	return lower, upper
}

/*********************
* AUXILIAR FUNCTIONS *
*********************/
func getGlobalClockTags() []string {
    return []string {
        "Clock.Information",
    }
}