package main

import (
	"github.com/agnivade/levenshtein"
	"math"
	"strings"
)

func isSimilar(bsIdx, cktIdx int) bool {
	bs := bsData[bsIdx]
	ckt := cktData[cktIdx]

	buyerNameSplit := strings.Split(ckt.buyerName, " ")
	descSplit := strings.Split(bs.desc, " ")

	maxScore := len(buyerNameSplit)
	score := 0

	for _, descPart := range descSplit {
		for _, buyerNamePart := range buyerNameSplit {
			if levenshtein.ComputeDistance(descPart, buyerNamePart) <= 3 {
				score++
			}
		}
	}

	if (score == maxScore || score >= 2) && math.Abs(bs.amount-ckt.amount) < 1e-6 {
		return true
	}

	return false
}
