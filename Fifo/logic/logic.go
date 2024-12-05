package logic

import (
	"Queue/Fifo/logic/utils"
)

//aplication

type Logic struct {
	Public       Queue
	Nonepriorety Queue
	Priorety     Queue
	CashPublic   int
}

const Priora string = "Get a Credit"

func NewLogic() *Logic {
	l := Logic{}
	l.Public = *New()
	l.Priorety = *New()
	l.Nonepriorety = *New()
	l.CashPublic = 1

	return &l
}

// бивает и распледиляет заказы в на приор и не
func (l *Logic) Add(offerName string) *Offer {

	if offerName == Priora {
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
	if l.Public.Len() < l.CashPublic {

		if l.Priorety.Isempty() {
			v := l.Nonepriorety.Dequeue()
			if v != nil {
				l.Public.Enqueue(*v)
			}
		} else {
			v := l.Priorety.Dequeue()
			if v != nil {
				l.Public.Enqueue(*v)
			}
		}
	}

	return &Offer{}
}

// выбор кассира выполнение заказа
func (l *Logic) Getoffer() *Offer {
	//len(l.Public)= l.CashPublic
	v := l.Public.Dequeue()
	l.choose()
	return v
}
