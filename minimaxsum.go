package main

import "fmt"

func totalSum(data []int64, except int64) int64 {
	count := int64(0)
	for _, dataElement := range data {
		if dataElement != except {
			count += dataElement
		}
	}

	return count
}

func main() {
	data := []int64{254961783, 604179258, 462517083, 967304281, 860273491}
	countedSum := []int64{}
	for _, dataEle := range data {
		countedSum = append(countedSum, totalSum(data, dataEle))
	}

	max := int64(0)
	min := int64(0)

	// search for minimum
	for _, searchMaxEle := range countedSum {
		if max < searchMaxEle {
			max = searchMaxEle
		}
	}

	// search max
	for _, searchMinEle := range countedSum {
		if max > searchMinEle {
			min = searchMinEle
		}
	}

	fmt.Println(countedSum)
	fmt.Println(fmt.Sprintf("%d %d", min, max))
}
