package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func loadRules() {
	f, err := os.Open("rules.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	isFirst := true

	fw, err := os.Create("result.csv")
	defer fw.Close()
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(fw)
	defer w.Flush()

	w.Write([]string{"rule", "confidence"})

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

		rule := record[0]
		confidence := evaluateConfidence(rule)
		confidenceStr := strconv.Itoa(confidence)

		w.Write([]string{rule, confidenceStr})
	}
}

func evaluateConfidence(rules string) int {
	var A, B, C, confidence int
	if _, err := fmt.Sscanf(rules,"%d>%d", &A, &B); err == nil {
		confidence = measureAThenB(A, B)
	} else if _, err := fmt.Sscanf(rules, "%d>%d&%d", &A, &B, &C); err == nil {
		confidence = measureAThenBC(A, B, C)
	} else {
		fmt.Sscanf(rules, "%d&%d>%d", &A, &B, &C)
		confidence = measureABThenC(A, B, C)
	}

	return confidence
}

func measureAThenB(a, b int) int {
	allProduct := make(map[int]int)

	for _, orderID := range orderData[a] {
		allProduct[orderID]++
	}
	for _, orderID := range orderData[b] {
		allProduct[orderID]++
	}

	allTrue := 0
	for _, product := range allProduct {
		if product == 2 {
			allTrue++
		}
	}

	return int(float64(allTrue) / float64(len(orderData[a])) * 1000)
}

func measureAThenBC(a, b, c int) int {
	allProduct := make(map[int]int)

	for _, orderID := range orderData[a] {
		allProduct[orderID]++
	}
	for _, orderID := range orderData[b] {
		allProduct[orderID]++
	}
	for _, orderID := range orderData[c] {
		allProduct[orderID]++
	}

	allTrue := 0
	for _, product := range allProduct {
		if product == 3 {
			allTrue++
		}
	}

	return int(float64(allTrue) / float64(len(orderData[a])) * 1000)
}

func measureABThenC(a, b, c int) int {
	allProduct := make(map[int]int)
	ab := make(map[int]int)

	for _, orderID := range orderData[a] {
		allProduct[orderID]++
		ab[orderID]++
	}
	for _, orderID := range orderData[b] {
		allProduct[orderID]++
		ab[orderID]++
	}
	for _, orderID := range orderData[c] {
		allProduct[orderID]++
	}

	allTrue := 0
	for _, product := range allProduct {
		if product == 3 {
			allTrue++
		}
	}
	abTrue := 0
	for _, product := range ab {
		if product == 2 {
			abTrue++
		}
	}

	return int(float64(allTrue) / float64(abTrue) * 1000)
}
