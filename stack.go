package clrs

type stack struct {
	top  int
	data []int
}

func (s *stack) New() {
	s.top = -1
	s.data = make([]int, 0)
}

func (s *stack) empty() bool {
	return s.top == -1
}

func (s *stack) push(x int) {
	s.top = s.top + 1
	if len(s.data) > s.top {
		s.data[s.top] = x
	} else {
		s.data = append(s.data, x)
	}
}

func (s *stack) pop() (int, error) {
	if s.empty() {
		return 0, errUnderflow
	}
	s.top = s.top - 1
	return s.data[s.top+1], nil
}
