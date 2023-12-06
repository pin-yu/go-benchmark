package gobenchmark

import "strings"

// str + str
func StringAdd() string {
	s := ""
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func StringBuilder() string {
	b := strings.Builder{}
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}
