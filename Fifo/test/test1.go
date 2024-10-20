package main

import "fmt"

func Qetnumber(data *int) int {

	*data = *data + 1
	return *data
}

func main() {
	var data int
	fmt.Println(Qetnumber(&data))
	fmt.Println(Qetnumber(&data))

}
