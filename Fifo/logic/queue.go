package logic

type Queue struct {
	Arr []Offer
}

func New() *Queue {

	return &Queue{Arr: []Offer{}}
}

func (q *Queue) Contains() bool {
	if len(q.Arr) == 0 {
		return false
	}
	return true
}
func (q *Queue) Dequeue() *Offer {
	//fmt.Println(q.Arr)
	if len(q.Arr) == 0 {
		return nil
	}
	r := q.Arr[0]
	q.Arr = q.Arr[1:]
	return &r
}
func (q *Queue) Enqueue(item Offer) {
	var exist bool
	for i := 0; i < len(q.Arr); i++ {
		if item.Number == q.Arr[i].Number {
			exist = true
		}
	}
	if !exist {

		q.Arr = append(q.Arr, item)
	}
}
func (q *Queue) Peek() *Offer {
	if len(q.Arr) == 0 {
		return nil
	}
	return &q.Arr[0]
}

func (q *Queue) Len() int {
	return len(q.Arr)
}
func Qetnumber(data *int) int {

	*data = *data + 1
	return *data
}
func (q *Queue) Isempty() bool {
	if len(q.Arr) == 0 {
		return true
	}
	return false
}

func (q *Queue) clear() {
	q.Arr = []Offer{}
}

//kakaha
//func (q *Queue) Contains() bool {
//	if len(q.Arr) == 0 {
//		return false
//	}
//	return true
//}
//func (q *Queue) Dequeue() string {
//	if len(q.Arr) == 0 {
//		return ""
//	}
//	r := q.Arr[0]
//	q.Arr = q.Arr[1:]
//	return r
//}
//func (q *Queue) Enqueue(item string) {
//	q.Arr = append(q.Arr, item)
//}
//func (q *Queue) Peek() string {
//	if len(q.Arr) == 0 {
//		return ""
//	}
//	return q.Arr[0]
//}

//	func New() *Queue {
//		return &Queue{}
//	}
//
//	func (q *Queue) Len() int {
//		return len(q.Arr)
//	}
