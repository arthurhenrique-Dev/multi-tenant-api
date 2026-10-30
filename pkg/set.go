package pkg

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}
func (s Set[T]) Remove(item T) {
	delete(s, item)
}
func (s Set[T]) Contains(item ...T) bool {
	for _, item := range item {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}
func (s Set[T]) HaveAll(items ...T) bool {
	checker := NewSet(items...)
	haveAllUtil(s, checker)
	return true
}
func (s Set[T]) HaveAllSet(items Set[T]) bool {
	return haveAllUtil(s, items)
}
func (s Set[T]) Len() int {
	return len(s)
}

func haveAllUtil[T comparable](src, items Set[T]) bool {
	for item := range src {
		if !items.Contains(item) {
			return false
		}
	}
	return true
}
