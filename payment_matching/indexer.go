package main

import (
	"strings"
)

func indexBs() {
	for idx, bs := range bsData {
		for _, descPart := range strings.Split(bs.desc, " ") {
			func(idx int, descPart string) {
				blockingLock.Lock()
				visitedLock.Lock()
				defer blockingLock.Unlock()
				defer visitedLock.Unlock()

				if blocking[descPart] == nil {
					blocking[descPart] = make(map[string][]int)
				}

				if visited[visitedType{descPart, "bs", idx}] {
					return
				} else {
					visited[visitedType{descPart, "bs", idx}] = true
				}

				blocking[descPart]["bs"] = append(blocking[descPart]["bs"], idx)
			}(idx, descPart)
		}
	}

	loaderWg.Done()
}

func indexCkt() {
	for idx, ckt := range cktData {
		for _, buyerNamePart := range strings.Split(ckt.buyerName, " ") {
			func(idx int, buyerNamePart string) {
				blockingLock.Lock()
				visitedLock.Lock()
				defer blockingLock.Unlock()
				defer visitedLock.Unlock()

				if blocking[buyerNamePart] == nil {
					blocking[buyerNamePart] = make(map[string][]int)
				}

				if visited[visitedType{buyerNamePart, "ckt", idx}] {
					return
				} else {
					visited[visitedType{buyerNamePart, "ckt", idx}] = true
				}

				blocking[buyerNamePart]["ckt"] = append(blocking[buyerNamePart]["ckt"], idx)
			}(idx, buyerNamePart)
		}
	}

	loaderWg.Done()
}
