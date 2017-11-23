package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type L struct {
	Name   string
	GSS    string
	Decile int
}

func main() {
	file, err := os.Open("2015imd.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	m := make(map[string]L)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !strings.HasPrefix(record[1], "Sheffield") {
			continue
		}
		i, _ := strconv.Atoi(record[5])
		m[record[0]] = L{record[1], record[0], i}
	}
	b, _ := json.MarshalIndent(m, "", "  ")
	os.Stdout.Write(b)
}
