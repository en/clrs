package clrs

type stack struct {
	top  int
	data []interface{}
}

func (s *stack) New() {
	s.top = -1
	s.data = make([]interface{}, 0)
}

func (s *stack) empty() bool {
	return s.top == -1
}

func (s *stack) push(x interface{}) {
	s.top = s.top + 1
	if len(s.data) > s.top {
		s.data[s.top] = x
	} else {
		s.data = append(s.data, x)
	}
}

func (s *stack) pop() (interface{}, error) {
	if s.empty() {
		return 0, errUnderflow
	}
	s.top = s.top - 1
	return s.data[s.top+1], nil
}
