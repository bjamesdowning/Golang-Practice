package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Example of how to work with CSV data given @
//appliedgo.net/spreadsheet
func main() {
	rows := readOrders("test.csv")
	rows = calculate(rows)
	writeOrders("ordersReport.csv", rows)
}

func readOrders(s string) [][]string {
	f, err := os.Open(s)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", s, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Cannot CSV Data: %s\n", err.Error())
	}
	return rows
}

func calculate(ss [][]string) [][]string {
	sum := 0
	nb := 0
	for i := range ss {
		if i == 0 {
			ss[0] = append(ss[0], "Total")
			continue
		}
		item := ss[i][2]

		price, err := strconv.Atoi(strings.Replace(ss[i][3], ".", "", -1))
		if err != nil {
			log.Fatalf("Cannot retrieve price of %s: %s\n", item, err)
		}
		qty, err := strconv.Atoi(ss[i][4])
		if err != nil {
			log.Fatalf("Cannot retrieve quantity of %s: %s\n", item, err)
		}

		total := price * qty
		sum += total

		if item == "Ball Pen" {
			nb += qty
		}
	}
	ss = append(ss, []string{"", "", "", "Sum", "", intToFloatString(sum)})
	ss = append(ss, []string{"", "", "", "", "Ball Pens", fmt.Sprint(nb), ""})

	return ss
}

func intToFloatString(n int) string {
	intgr := n / 100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d", intgr, frac)
}

func writeOrders(s string, ss [][]string) {
	f, err := os.Create(s)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", s, err.Error())
	}
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", s, err.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(ss)
}
