package clrs

type queue struct {
	head int
	tail int
	data []int
}

func (q *queue) New(size int) {
	q.head = 0
	q.tail = 0
	q.data = make([]int, size)
}

func (q *queue) enqueue(x int) error {
	var ntail int
	if q.tail == len(q.data)-1 {
		ntail = 0
	} else {
		ntail = q.tail + 1
	}
	if q.head == ntail {
		return errOverflow
	}
	q.data[q.tail] = x
	if q.tail == len(q.data)-1 {
		q.tail = 0
	} else {
		q.tail = q.tail + 1
	}
	return nil
}

func (q *queue) dequeue() (int, error) {
	if q.head == q.tail {
		return 0, errUnderflow
	}
	x := q.data[q.head]
	if q.head == len(q.data)-1 {
		q.head = 0
	} else {
		q.head = q.head + 1
	}
	return x, nil
}
