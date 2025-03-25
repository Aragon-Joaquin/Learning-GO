package main

import (
	"strconv"
)

func findRelativeRanks(score []int) []string {
	hashMap := map[int]string{1: "Gold Medal", 2: "Silver Medal", 3: "Bronze Medal"}
	arrayOfResults := make([]string, len(score))

	for idx, val1 := range score {
		valueRank := len(score)

		for _, val2 := range score {
			if val2 < val1 {
				valueRank = valueRank - 1
			}
		}
		result, isOk := hashMap[valueRank]
		if isOk == true {
			arrayOfResults[idx] = result
		} else {
			arrayOfResults[idx] = strconv.Itoa(valueRank)
		}
	}

	return arrayOfResults
}
