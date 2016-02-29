package clrs

type deque struct {
	head int
	tail int
	data []int
}

func (d *deque) New(size int) {
	d.head = 0
	d.tail = 0
	d.data = make([]int, size)
}

// push back
func (d *deque) push(x int) error {
	var ntail int
	if d.tail == len(d.data)-1 {
		ntail = 0
	} else {
		ntail = d.tail + 1
	}
	if d.head == ntail {
		return errOverflow
	}
	d.data[d.tail] = x
	if d.tail == len(d.data)-1 {
		d.tail = 0
	} else {
		d.tail = d.tail + 1
	}
	return nil
}

// push front
func (d *deque) unshift(x int) error {
	var ntail int
	if d.tail == len(d.data)-1 {
		ntail = 0
	} else {
		ntail = d.tail + 1
	}
	if d.head == ntail {
		return errOverflow
	}
	if d.head == 0 {
		d.head = len(d.data) - 1
	} else {
		d.head = d.head - 1
	}
	d.data[d.head] = x
	return nil
}

// pop back
func (d *deque) pop() (int, error) {
	if d.head == d.tail {
		return 0, errUnderflow
	}
	if d.tail == 0 {
		d.tail = len(d.data) - 1
	} else {
		d.tail = d.tail - 1
	}
	return d.data[d.tail], nil
}

// pop front
func (d *deque) shift() (int, error) {
	if d.head == d.tail {
		return 0, errUnderflow
	}
	x := d.data[d.head]
	if d.head == len(d.data)-1 {
		d.head = 0
	} else {
		d.head = d.head + 1
	}
	return x, nil
}
