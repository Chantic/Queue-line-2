package main

import (
	"bufio"
	"fmt"
	"os"
)

type Offer struct {
	number int
	name   string
}

type Queue struct {
	arr []Offer
}

func main() {

	var Nonepriorety = Queue{}
	var Priorety = Queue{}
	var Public = Queue{}
	var data int
	var data2 int
	var data3 int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose your role\n 1.Customer\n 2.Cashier")

	for scanner.Scan() {

		fmt.Println("Choose your role\n 1.Customer\n 2.Cashier")
		input := scanner.Text()

		switch input {

		case "-1":
			fmt.Println("Choose service\n 1.Order offer\n 2.Show number \n, 3.Exchange money\n, 4.Get a Credit \n, 5. Get an Allowance ")

			var input1 string
			//fmt.Scan(&input1)

			switch input1 {
			case "-1":
				Qetnumber(&data)
				Nonepriorety.Enqueue(Offer{number: data,
					name: "Order offer"})
				fmt.Println(Nonepriorety)
			case "-2":
				Nonepriorety.Peek()
				fmt.Println(Nonepriorety.Peek())

			case "-3":
				Qetnumber(&data)
				Nonepriorety.Enqueue(Offer{number: data,
					name: "Exchange money"})
				fmt.Println(" offer Exchange money, your number", data)

				if Priorety.Isempty() {

					Public.Enqueue(*Nonepriorety.Peek())

					//fmt.Println(" done offer (Exchange money),  number", num)
				} else {
					Public.Enqueue(*Priorety.Peek())
					//fmt.Println(" done offer (Exchange money),  number", num)
				}

			case "-4":
				Qetnumber(&data2)
				Priorety.Enqueue(Offer{number: data2,
					name: "Get a Credit "})
				fmt.Println(" offer Get a Credit , your number", data2)

				if Priorety.Isempty() {

					Public.Enqueue(*Nonepriorety.Peek())

					//fmt.Println(" done offer (Exchange money),  number", num)
				} else {
					Public.Enqueue(*Priorety.Peek())
					//fmt.Println(" done offer (Exchange money),  number", num)
				}

			case "-5":
				Qetnumber(&data)
				Nonepriorety.Enqueue(Offer{number: data,
					name: " Get an Allowance "})
				fmt.Println(" offer an Get an Allowance , your number", data)
				Qetnumber(&data3)
				if Priorety.Isempty() {

					Public.Enqueue(*Nonepriorety.Peek())

					//fmt.Println(" done offer (Exchange money),  number", num)
				} else {
					Public.Enqueue(*Priorety.Peek())
					//fmt.Println(" done offer (Exchange money),  number", num)
				}

			default:
				fmt.Printf("Error: wrong command for Customer ")

			}

		case "-2":

			fmt.Println("Choose Service ,\n 1.done offer,\n 2.show offer  \n ")

			var input2 string
			//fmt.Scan(&input2)

			println("in", input2)
			switch input2 {

			case "-1":
				num := Public.Dequeue()
				fmt.Println("offer is done, number ", num)

				//if Priorety.Isempty() {
				//	num := Nonepriorety.Dequeue()
				//	fmt.Println(" done offer (Exchange money),  number", num)
				//} else {
				//	num := Priorety.Dequeue()
				//	fmt.Println(" done offer (Exchange money),  number", num)
				//}

			case "-2":
				num2 := Nonepriorety.Peek()
				fmt.Println(" show offer   ,  number", num2)

			default:
				fmt.Printf("Error: wrong command for Cashier\n")

			}

		}
	}
}

func (q *Queue) Contains() bool {
	if len(q.arr) == 0 {
		return false
	}
	return true
}
func (q *Queue) Dequeue() *Offer {
	//fmt.Println(q.arr)
	if len(q.arr) == 0 {
		return nil
	}
	r := q.arr[0]
	q.arr = q.arr[1:]
	return &r
}
func (q *Queue) Enqueue(item Offer) {
	var exist bool
	for i := 0; i < len(q.arr); i++ {
		if item.number == q.arr[i].number {
			exist = true
		}
	}
	if !exist {

		q.arr = append(q.arr, item)
	}
}
func (q *Queue) Peek() *Offer {
	if len(q.arr) == 0 {
		return nil
	}
	return &q.arr[0]
}

//	func New() *Queue {
//		return &Queue{}
//	}
func (q *Queue) Len() int {
	return len(q.arr)
}
func Qetnumber(data *int) int {

	*data = *data + 1
	return *data
}
func (q *Queue) Isempty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

//func Checking(q *Queue) {
//
//	if Priorety.Isempty() {
//
//		Public.Enqueue(*Nonepriorety.Peek())
//
//		//fmt.Println(" done offer (Exchange money),  number", num)
//	} else {
//		Public.Enqueue(*Priorety.Peek())
//		//fmt.Println(" done offer (Exchange money),  number", num)
//	}
//}
//
