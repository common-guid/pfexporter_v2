package parseStdout

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var (
	Num         = make(map[int]string)
	IP          = make(map[int]string)
	Dir         = make(map[int]string)
	Two         = make(map[int]string)
	Ten         = make(map[int]string)
	Fourty      = make(map[int]string)
	Total       = make(map[int]string)
	Bytes_40sec = make(map[int]float64)
	Bytes_10sec = make(map[int]float64)
	Bytes_total = make(map[int]float64)
)

func Parse_iftop(data []byte) {
	fileTextLines := make(map[int]string)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines) // Set up the split function.
	i := 0
	for scanner.Scan() {
		fileTextLines[i] = scanner.Text()
		i++
	}

	start, end, line := 2, (len(fileTextLines) - 10), 1

	for line <= end {
		if line > start {
			words := strings.Fields(fileTextLines[line])
			if line%2 != 0 {
				Num[line], IP[line], Dir[line], Two[line], Ten[line], Fourty[line], Total[line] = words[0], words[1], words[2], words[3], words[4], words[5], words[6]
			} else {
				IP[line], Dir[line], Two[line], Ten[line], Fourty[line], Total[line] = words[0], words[1], words[2], words[3], words[4], words[5]
			}
		}
		//println(Num[line], IP[line], Dir[line], Two[line], Ten[line], Fourty[line], Total[line])
		line++
	}
	for line, item := range Dir {
		if item == "<=" {
			Dir[line] = "dest"
		} else {
			Dir[line] = "src"
		}
		mre := regexp.MustCompile("M")
		kre := regexp.MustCompile("K")
		//compile the 40s list of values
		for line, value := range Fourty {
			if mre.Match(([]byte(value))) {
				newval := regexp.MustCompile("MB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				intval, _ := strconv.ParseFloat(newval2, 64)
				//	println(newval2, reflect.TypeOf(newval2))
				finval := intval * 1000000
				//	println(finval, reflect.TypeOf(finval))
				//add to new list
				Bytes_40sec[line] = finval
			} else if kre.Match(([]byte(value))) {
				newval := regexp.MustCompile("KB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				intval, _ := strconv.ParseFloat(newval2, 64)
				finval := intval * 1000
				//add to new list
				Bytes_40sec[line] = finval
			} else {
				newval := regexp.MustCompile("B").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				finval, _ := strconv.ParseFloat(newval2, 64)
				//add to new list
				Bytes_40sec[line] = finval
			}
		}
		// compile the 10s list of values
		for line, value := range Ten {
			if mre.Match(([]byte(value))) {
				newval := regexp.MustCompile("MB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				intval, _ := strconv.ParseFloat(newval2, 64)
				//	println(newval2, reflect.TypeOf(newval2))
				finval := intval * 1000000
				//	println(finval, reflect.TypeOf(finval))
				//add to new list
				Bytes_10sec[line] = finval
			} else if kre.Match(([]byte(value))) {
				newval := regexp.MustCompile("KB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				intval, _ := strconv.ParseFloat(newval2, 64)
				finval := intval * 1000
				//add to new list
				Bytes_10sec[line] = finval
			} else {
				newval := regexp.MustCompile("B").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				finval, _ := strconv.ParseFloat(newval2, 64)
				//add to new list
				Bytes_10sec[line] = finval
			}
		}
		// compile the cumulative list of values
		for line, value := range Total {
			if mre.Match(([]byte(value))) {
				newval := regexp.MustCompile("MB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				intval, _ := strconv.ParseFloat(newval2, 64)
				//	println(line, newval2, reflect.TypeOf(newval2))
				finval := intval * 1000000
				//	println(line, finval, reflect.TypeOf(finval))
				//add to new list
				Bytes_total[line] = finval
			} else if kre.Match(([]byte(value))) {
				newval := regexp.MustCompile("KB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				//println(line, newval2, reflect.TypeOf(newval2))
				intval, _ := strconv.ParseFloat(newval2, 64)
				finval := intval * 1000
				//println(line, finval, reflect.TypeOf(finval))
				//add to new list
				Bytes_total[line] = finval
			} else {
				newval := regexp.MustCompile("B").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				finval, _ := strconv.ParseFloat(newval2, 64)
				//add to new list
				Bytes_total[line] = finval
			}
		}
	}
}
