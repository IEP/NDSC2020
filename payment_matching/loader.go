package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

func bankStatementLoad() {
	f, err := os.Open("bank_statement.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	isFirst := true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isFirst {
			isFirst = false
			continue
		}

		idx, _ := strconv.Atoi(record[0])
		amount, _ := strconv.ParseFloat(record[1], 64)
		desc := strings.ToLower(strings.TrimSpace(record[2]))
		desc = reAlpha.ReplaceAllString(desc, "")
		desc = reSpace.ReplaceAllString(desc, " ")

		bsData[idx] = bsType{
			amount: amount,
			desc:   desc,
		}
	}
	loaderWg.Done()
}

func checkoutLoad() {
	f, err := os.Open("checkout.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	isFirst := true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		if isFirst {
			isFirst = false
			continue
		}

		idx, _ := strconv.Atoi(record[0])
		amount, _ := strconv.ParseFloat(record[1], 64)
		buyerName := strings.ToLower(strings.TrimSpace(record[2]))
		buyerName = reAlpha.ReplaceAllString(buyerName, "")
		buyerName = reSpace.ReplaceAllString(buyerName, " ")

		cktData[idx] = cktType{
			amount:    amount,
			buyerName: buyerName,
		}
	}
	loaderWg.Done()
}
