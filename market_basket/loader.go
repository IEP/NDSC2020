package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func loadOrder() {
	f, err := os.Open("association_order.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

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

		orderID, _ := strconv.Atoi(record[0])
		itemID, _ := strconv.Atoi(record[1])

		orderData[itemID] = append(orderData[itemID], orderID)
	}
}
