package main

/*
to be honest, i just copied this answer since i was trying to do everything the most simplistic & performant way without using any libraries but oh god
it turned out to be more tedious and required a lot of boilerplate. Since i didn't want to start typing generic functions it consumed my mind (and misread the problem too)
but hey! at least I learned something new :)
*/

import (
	"sort"
)

func sortEvenOdd(nums []int) []int {
	var oddArr, evenArr []int

	for idx, val := range nums {
		if idx%2 == 0 {
			evenArr = append(evenArr, val)
		} else {
			oddArr = append(oddArr, val)
		}
	}

	sort.Slice(evenArr, func(i, j int) bool {
		return evenArr[i] < evenArr[j]
	})

	sort.Slice(oddArr, func(i, j int) bool {
		return oddArr[i] > oddArr[j]
	})

	resultArray := make([]int, len(nums))
	a, b := 0, 0

	for i := 0; i < len(resultArray); i++ {
		if i%2 == 0 {
			resultArray[i] = evenArr[a]
			a++
		} else {
			resultArray[i] = oddArr[b]
			b++
		}
	}

	return resultArray
}
