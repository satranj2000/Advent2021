package utils

import "log"

type Stack struct {
	items []string
}

func (s *Stack) Push(d string) {
	s.items = append(s.items, d)
}

func (s *Stack) Pop() (val string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Pop failed with error -  ", r)
			err = r.(error)
		}
	}()
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, err
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) <= 0
}
