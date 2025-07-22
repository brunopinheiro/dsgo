package algorithms

type runeStack struct {
	runes []rune
}

func (s *runeStack) empty() bool {
	return len(s.runes) == 0
}

func (s *runeStack) push(r rune) {
	s.runes = append(s.runes, r)
}

func (s *runeStack) pop() rune {
	if s.empty() {
		panic("stack is empty")
	}
	r := s.runes[len(s.runes)-1]
	s.runes = s.runes[:len(s.runes)-1]
	return r
}

func (s *runeStack) peek() rune {
	if s.empty() {
		panic("stack is empty")
	}
	return s.runes[len(s.runes)-1]
}

func (s *runeStack) String() string {
	return string(s.runes)
}

func RemoveAdjacentDuplicates(input string) string {
	s := &runeStack{}

	for _, r := range input {
		if !s.empty() && r == s.peek() {
			s.pop()
		} else {
			s.push(r)
		}
	}
	return s.String()
}
