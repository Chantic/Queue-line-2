package controller

import (
	"Queue/Fifo/logic"
	"fmt"
)

type Customer struct {
	lg *logic.Logic
}

func NewCostumer(logic2 *logic.Logic) *Customer {
	c := Customer{
		lg: logic2,
	}
	return &c
}

func (c Customer) ExchangeMoney() {
	c.lg.Add("Exchange money")

	fmt.Println("offer Exchange money")
}
func (c Customer) GetCredit() {
	c.lg.Add(logic.Priora)
	fmt.Println("Get a Credit")

}
func (c Customer) GetAllowance() {
	c.lg.Add("Get a Allowance ")
	fmt.Println("Get a Allowance ")
}
