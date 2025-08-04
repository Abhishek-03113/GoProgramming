package dataStructures

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, bool) {

	if len(q.items) == 0 {
		return -1, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

func (q *Queue) Peek() (interface{}, bool) {

	if len(q.items) == 0 {
		return -1, false
	}

	item := q.items[0]

	return item, true
}
