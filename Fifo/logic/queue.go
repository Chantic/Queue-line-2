package logic

type Queue struct {
	arr []string
}

func (q *Queue) clear() {
	q.arr = []string{}
}
func (q *Queue) Contains() bool {
	if len(q.arr) == 0 {
		return false
	}
	return true
}
func (q *Queue) Dequeue() string {
	if len(q.arr) == 0 {
		return ""
	}
	r := q.arr[0]
	q.arr = q.arr[1:]
	return r
}
func (q *Queue) Enqueue(item string) {
	q.arr = append(q.arr, item)
}
func (q *Queue) Peek() string {
	if len(q.arr) == 0 {
		return ""
	}
	return q.arr[0]
}

//	func New() *Queue {
//		return &Queue{}
//	}
func (q *Queue) Len() int {
	return len(q.arr)
}
