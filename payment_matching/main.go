package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type bsType struct {
	amount float64
	desc   string
}

var bsData = make(map[int]bsType)

type cktType struct {
	amount    float64
	buyerName string
}

var cktData = make(map[int]cktType)

var loaderWg sync.WaitGroup
var reAlpha *regexp.Regexp
var reSpace *regexp.Regexp

type visitedType struct {
	strIdx   string
	bsOrCkt  string
	entryIdx int
}

// name substr, bs/ckt, idx
var blocking = make(map[string]map[string][]int)
var blockingLock sync.Mutex

// blocking visit
var visited = make(map[visitedType]bool)
var visitedLock sync.Mutex

var bsVisited = make(map[int]bool)
var cktVisited = make(map[int]bool)

type resultType struct {
	bsIdx  int
	cktIdx int
}

type resultArr []resultType

func (r resultArr) Len() int {
	return len(r)
}

func (r resultArr) Less(i, j int) bool {
	return r[i].bsIdx < r[j].bsIdx
}

func (r resultArr) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

var result = make(resultArr, 0)

var tempResult = make(map[int]int)

func main() {
	reAlpha, _ = regexp.Compile("[^a-z ]+")
	reSpace, _ = regexp.Compile(" {2,}")

	loaderWg.Add(2)
	go bankStatementLoad()
	go checkoutLoad()
	loaderWg.Wait()

	loaderWg.Add(2)
	go indexBs()
	go indexCkt()
	loaderWg.Wait()

	fmt.Println(len(bsData))
	fmt.Println(len(cktData))
	fmt.Println(len(blocking))

	for cktIdx, content := range cktData {
		buyerNameSplit := strings.Split(content.buyerName, " ")
		for _, buyerNamePart := range buyerNameSplit {
			bsIdxs := blocking[buyerNamePart]["bs"]
			for _, bsIdx := range bsIdxs {
				if bsVisited[bsIdx] {
					continue
				}
				if cktVisited[cktIdx] {
					continue
				}
				if isSimilar(bsIdx, cktIdx) {
					bsVisited[bsIdx] = true
					cktVisited[cktIdx] = true
					tempResult[bsIdx] = cktIdx
				}
			}
		}
	}

	for i := 1; i <= 240000; i++ {
		result = append(result, resultType{i, tempResult[i]})
	}

	sort.Sort(result)

	fw, err := os.Create("result.csv")
	if err != nil {
		panic(err.Error())
	}
	defer fw.Close()

	w := csv.NewWriter(fw)
	defer w.Flush()

	w.Write([]string{"stmt_id", "ckt_id"})

	for _, res := range result {
		bsIdx := strconv.Itoa(res.bsIdx)
		cktIdx := strconv.Itoa(res.cktIdx)

		w.Write([]string{bsIdx, cktIdx})
	}
}
