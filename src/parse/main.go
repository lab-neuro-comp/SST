package main

import "os"
import "fmt"
import "bufio"
import "./filter"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		moment, comment := filter.Process(scanner.Text())
		if len(comment) > 0 {
			fmt.Println(comment, ":", moment)
		}
	}
}