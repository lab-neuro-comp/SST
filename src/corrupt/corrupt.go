package main

import "fmt"
import "os"
import "github.com/lab-neuro-comp/SST/src/sst"

func main() {
	// TODO Read file
	dataFile := os.Args[1]
	fp, _ := os.Open(dataFile)
	defer fp.Close()
	for bit := sst.ReadLine(fp); bit != 0; bit = sst.ReadLine(fp) {
		fmt.Println(line)
		fmt.Println("---")
	}
}
