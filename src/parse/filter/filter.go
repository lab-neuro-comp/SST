package filter

import "strings"

/* INTERFACE FUNCTIONS */
func Process(inlet string) (string, string) {
	var data []string
	var moment string
	var comment string
	var outlet string

	if (len(inlet) == 0) {
		return "", ""
	}

	data = strings.Split(inlet, " ")
	moment = data[0]
	comment = reduce(data[1:])
	outlet = moment

	if len(moment) > 3 {
		outlet = comment
	} else {
		outlet = " "
	}

	return moment, outlet[1:]
}

/* AUXILIAR FUNCTIONS */
func reduce(inlet []string) string {
	outlet := inlet[0]

	for i := 1; i < len(inlet); i++ {
		outlet = outlet + " " + inlet[i]
	}

	return outlet
}