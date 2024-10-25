package utils

var number int = 0

func UniqeNumbers() int {
	d := number
	number++
	return d
}
