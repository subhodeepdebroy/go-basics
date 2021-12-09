package main

// func main() {
// 	var toSearch int = 1
// 	input := []int{7, 8, 7, 9, 1, 7, 1, 4, 1, 3}
// 	fmt.Print(BubbleSort((input)))
// 	fmt.Print(BinarySearch(toSearch, input))
// 	//fmt.Print(input)
// }

func BubbleSort(in []int) []int {
	var temp int
	for i := 1; i < len(in); i++ {
		for j := 0; j < len(in)-i; j++ {
			if in[j] > in[j+1] {
				temp = in[j]
				in[j] = in[j+1]
				in[j+1] = temp
			}
		}
	}
	return in
}

func BinarySearch(toSearch int, inputArray []int) bool {
	if len(inputArray) != 0 {
		var mid int = len(inputArray) / 2
		if inputArray[mid] == toSearch {
			return true
		} else if inputArray[mid] < toSearch {
			return BinarySearch(toSearch, inputArray[mid+1:])
		} else {
			return BinarySearch(toSearch, inputArray[:mid])
		}
	} else {
		return false
	}
}
