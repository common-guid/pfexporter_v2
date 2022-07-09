package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	parseStdout "github.com/common-guid/iftop_parser/parser"
	prom "github.com/common-guid/iftop_parser/prom"
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

// https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/

func main() {
	go func() {
		prom.Init()
	}()

	i := 0
	for i < 3 {
		raw_data := cmdExc()
		parseStdout.Parse_iftop(raw_data)

		//println(parsed_data)

		go func() {
			prom.Display_metrics(parseStdout.Dir, parseStdout.IP, parseStdout.Bytes_total, prom.Ip_addr)
			for line, value := range parseStdout.Dir {
				fmt.Println(line, value)
			}
		}()
	}
}
func cmdExc() []byte {
	cmd := exec.Command("sudo", "iftop", "-nbBt", "-s 5")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	err = cmd.Wait()
	return data
}
