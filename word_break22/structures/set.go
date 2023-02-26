package structures

type Set map[string]struct{}
func (s Set) Add(word string) {
	s[word] = struct{}{}
}
func (s Set) Has(word string) bool {
	_, ok := s[word]
	return ok
}
