package sst

import "os"

// Reads the choosen file and translates it into a collection of arrays. It
// needs the name of the input file, and creates a map relating the needed
// information, as described in the `GetNeededVariables()` function; and the
// chosen fields.
func Read(input string) map[string][]string {
	inlet, _ := os.Open(input)
    defer inlet.Close()
	header := ReadHeader(inlet, GetNeededVariables())
	outlet := ReadRecords(inlet, header)
	return outlet
}

// Gets the indexes of the needed fields from the header.
func ReadHeader(inlet *os.File, needed []string) map[string]int {
	ReadLine(inlet)
	header := make(map[string]int)
	stuff := Split(ReadLine(inlet), '\t')

	for i, it := range stuff {
		if Contains(needed, it) {
			header[it] = i
		}
	}

	return header
}

// Read the records into arrays of contigous information. Requires the pointer
// to the file that is being read; and the map created on the `ReadHeader()`
// function, relating the indexes and the needed fields.
func ReadRecords(inlet *os.File, header map[string]int) map[string][]string {
	outlet := make(map[string][]string)

	for line := ReadLine(inlet); len(line) > 2; line = ReadLine(inlet) {
		stuff := Split(line, '\t')
		for key, value := range header {
			outlet[key] = append(outlet[key], stuff[value])
		}
	}

	return outlet
}
