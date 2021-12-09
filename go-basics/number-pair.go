package pairs

import "errors"

// func main() {
// 	var sumOfTwo int = 12
// 	inputArray := []int{7, 8, 7, 9, 1, 7, 1, 4, 1, 3}
// 	FindPair(inputArray, sumOfTwo)

// }

func FindPair(inputArray []int, sumOfTwo int) ([2]int, error) {
	m := make(map[int]int)
	var temp, i int
	for i = 0; i < len(inputArray); i++ {
		temp = sumOfTwo - inputArray[i]
		if m[temp] == 1 {
			//fmt.Printf("Valid Pairs are %d, %d", temp, inputArray[i])
			break
		} else {
			if i == len(inputArray)-1 {
				//fmt.Print("No valid pair")
				var nullArray [2]int
				return nullArray, errors.New("No valid pair")
				//break
			}
			m[inputArray[i]] = 1
		}
	}
	resp := [2]int{temp, inputArray[i]}
	return resp, nil

}
