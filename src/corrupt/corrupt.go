package main

import "fmt"
import "os"
import "github.com/lab-neuro-comp/SST/src/sst"

func main() {
	dataFile := os.Args[1]
	lines := make([]string, 0)

	// TODO Read file
	fp, _ := os.Open(dataFile)
	for line := sst.ReadLine(fp); len(line) > 0; line = sst.ReadLine(fp) {
		lines = append(lines, line)
	}
	fp.Close()

	// TODO Replace dots to commas
	lines = replace(lines, '.', ',')

	// TODO Save to file
	fp, _ = os.Create(dataFile)
	for _, line := range lines {
		fp.WriteString(fmt.Sprintf("%s\n", line))
	}
	fp.Close()
}

func replace(inlet []string, old, new rune) []string {
	outlet := make([]string, 0)

	for _, it := range inlet {
		outlet = append(outlet, sst.ReplaceInString(it, old, new))
	}

	return outlet
}
