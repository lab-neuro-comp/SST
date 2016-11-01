package sst

import "os"
import "fmt"
import "math"
import "strconv"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func replaceInString(inlet string, old, now rune) string  {
	outlet := ""

	for _, letter := range inlet {
		if letter == old {
			outlet += string(now)
		} else {
			outlet += string(letter)
		}
	}

	return outlet
}

// Reads a whole line from the file in pointer.
func ReadLine(inlet *os.File) string {
	stuff := ""

	for data := ReadChar(inlet); data != '\n'; data = ReadChar(inlet) {
		stuff += string(data)
	}

	return stuff
}

// Reads a single char from the file pointer.
func ReadChar(inlet *os.File) byte {
	data := make([]byte, 2)
	_, shit := inlet.Read(data)

	if shit != nil {
		return '\n'
	} else {
		return data[0]
	}
}

// Splits a string, dividing them based upon a char.
func Split(inlet string, separator byte) []string {
	var outlet []string
	division := ""

	for i := 0; i < len(inlet); i++ {
		if inlet[i] != separator {
			division += string(inlet[i])
		} else {
			outlet = append(outlet, division)
			division = ""
		}
	}

	if len(division) > 0 {
		outlet = append(outlet, division)
	}

	return outlet
}

// Checks if there is a string inside an string array.
func Contains(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}


// Turns a string into a float64
func ParseFloat64(inlet string) float64 {
	outlet, shit := strconv.ParseFloat(inlet, 64)
	if shit == nil {
		return outlet
	} else {
		return 0
	}
}

// Turns a string into an integer.
func ParseInt(inlet string) int {
	outlet, shit := strconv.Atoi(inlet)
	if shit == nil {
		return outlet
	} else {
		return 0
	}
}

// Provides the most simplistic interface to the standard output.
func Debug(x interface{}) {
	fmt.Println(x)
}

// Checks if a file has a valid format, according the test's specification.
func ValidFile(inlet string) bool {
	isTxt := true
	isCsv := false
	txt := ".txt"
	csv := ".csv"

	for i := 1; i <= 4 && (isTxt || isCsv); i++ {
		if inlet[len(inlet)-i] != txt[len(txt)-i] {
			isTxt = false
		}
		if inlet[len(inlet)-i] != csv[len(csv)-i] {
			isCsv = false
		}
	}

	return isTxt || isCsv
}

/////////////////
// MATHEMATICS //
/////////////////

// Provides a simple sigma operator, analogous to the mathematical sigma
// operator.
func Sigma(inlet []float64, op func(float64)float64) float64 {
	var outlet float64 = 0.0
	for _, it := range inlet {
		outlet += op(it)
	}
	return outlet
}

// Calculates the mean of am array of floats.
func Mean(inlet []float64) float64 {
	return Sigma(inlet, func(it float64) float64 {
		return it
	}) / float64(len(inlet))
}

// Calculates the variance of am array of floats.
func Variance(inlet []float64) float64 {
	mean := Mean(inlet)
	return Sigma(inlet, func(it float64) float64 {
		return math.Pow(it - mean, 2)
	}) / float64(len(inlet))
}

// Calculates the standard deviation of am array of floats.
func StdDev(inlet []float64) float64 {
	return math.Sqrt(Variance(inlet))
}
