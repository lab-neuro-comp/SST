package sst

import "fmt"
import "time"
import "strings"
import "strconv"

/*****************
* MAIN FUNCTIONS *
*****************/
func ConvertToUnixTime(timestamp string) int {
    // Example of timestamp: 2015-09-18T15:22:21Z
    year, month, day := getDate(timestamp)
    hour, min, sec := getHour(timestamp)
    moment := time.Date(year, month, day, hour, min, sec, 0, time.Local)
    return int(moment.Unix())
}

func ConvertToTimeStamp(unixtime int) string {
    // TODO Convert unix t
    return fmt.Sprintf("%#v", unixtime)
}

/*********************
* AUXILIAR FUNCTIONS *
*********************/
func getDate(timestamp string) (int, time.Month, int)  {
    raw := strings.Split(strings.Split(timestamp, "T")[0], "-")
    year, _ := strconv.ParseInt(raw[0], 10, 0)
    month, _ := strconv.ParseInt(raw[1], 10, 0)
    day, _ := strconv.ParseInt(raw[2], 10, 0)
    return int(year), getMonth(int(month)), int(day)
}

func getMonth(inlet int) time.Month {
    outlet := time.January

    if inlet == 2 {
        outlet = time.February
    } else if inlet == 3 {
        outlet = time.March
    } else if inlet == 4 {
        outlet = time.April
    } else if inlet == 5 {
        outlet = time.May
    } else if inlet == 6 {
        outlet = time.June
    } else if inlet == 7 {
        outlet = time.July
    } else if inlet == 8 {
        outlet = time.August
    } else if inlet == 9 {
        outlet = time.September
    } else if inlet == 10 {
        outlet = time.October
    } else if inlet == 11 {
        outlet = time.November
    } else if inlet == 12 {
        outlet = time.December
    }

    return outlet
}

func getHour(timestamp string) (int, int, int) {
    return 0, 0, 0
}

/* CODE IS POETRY */
