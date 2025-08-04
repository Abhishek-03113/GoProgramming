package dataStructures

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return -1, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	return s.items[len(s.items)-1], true
}
