package logic

import "Queue/Fifo/logic/utils"

type Logic struct {
	Public       Queue
	Nonepriorety Queue
	Priorety     Queue
}

func NewLogic() *Logic {
	l := Logic{}
	l.Public = *New()
	l.Priorety = *New()
	l.Nonepriorety = *New()

	return &l
}

// разбивает и распледиляет заказы в на приор и не
func (l *Logic) Add(offerName string) *Offer {

	if offerName == "-4" {
		l.Priorety.Enqueue(Offer{Name: offerName,
			Number: utils.UniqeNumbers()})

	} else {
		l.Nonepriorety.Enqueue(Offer{Name: offerName,
			Number: utils.UniqeNumbers()})
	}
	l.choose()
	return &Offer{}
}

// выполняет условия добавление на выход очереди
func (l *Logic) choose() *Offer {
	if l.Priorety.Isempty() {
		l.Public.Enqueue(*l.Nonepriorety.Dequeue())

		//fmt.Println(" done offer (Exchange money),  number", num)
	} else {
		//element := (*s)[len(*s)-1]
		//*s = (*s)[:len(*s)-1]

		l.Public.Enqueue(*l.Priorety.Dequeue())
		//fmt.Println(" done offer (Exchange money),  number", num)
	}
	return &Offer{}
}

// выбор кассира выполнение заказа
func (l *Logic) Getoffer() *Offer {
	return l.Public.Dequeue()
}
