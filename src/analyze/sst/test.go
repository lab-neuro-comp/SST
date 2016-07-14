package sst

import "fmt"

func TestTime() {
    y, m, d := getDate("2015-09-18T15:22:21Z")
    fmt.Printf("Testing date: %#v %#v %#v\n", y, m, d)
    hr, mn, sc := getHour("2015-09-18T15:22:21Z")
    fmt.Printf("Testing hour: %#v %#v %#v\n", hr, mn, sc)
}
