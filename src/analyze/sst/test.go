package sst

import "fmt"

func TestTime() {
    timestamp := "2015-09-18T15:22:21Z"
    y, m, d := getDate(timestamp)
    fmt.Printf("Testing date: %#v %#v %#v\n", y, m, d)
    hr, mn, sc := getHour(timestamp)
    fmt.Printf("Testing hour: %#v %#v %#v\n", hr, mn, sc)
    fmt.Printf("Final test: %#v\n", ConvertToUnixTime(timestamp))
}
