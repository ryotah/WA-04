package numutils

import "strconv"

func FizzBuzz(num int) string {
	// if num%15 == 0 {
	// 	return "Fizz Buzz"
	// }
	// if num%3 == 0 {
	// 	return "Fizz"
	// }
	// if num%5 == 0 {
	// 	return "Buzz"
	// }
	// return strconv.Itoa(num)
	switch {
	case num%15 == 0:
		return "Fizz Buzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}
}
