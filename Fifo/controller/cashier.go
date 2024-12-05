package controller

import (
	"Queue/Fifo/logic"
	"fmt"
)

type Cashier struct {
	lg *logic.Logic
}

func NewCashier(logic2 *logic.Logic) *Cashier {
	c := Cashier{
		lg: logic2,
	}
	return &c
}
func (c Cashier) GetOffer() {
	num := c.lg.Getoffer()
	fmt.Println("offer is done, number", num)

}

func (c Cashier) ShowOffer() {
	num2 := c.lg.Public.Peek()
	fmt.Println(" show offer   ,  number", num2)

}
