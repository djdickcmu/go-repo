package string_set

//testupdate

import "fmt"

type StringSet map[string]bool

func NewStringSet() StringSet {
	return make(StringSet)
}
func (s StringSet) Add(val string) {
	s[val] = true
}
func (s StringSet) AddAll(src StringSet) {
	for k, _ := range src {
		s[k] = true
	}
}
func (s StringSet) String() string {
	return fmt.Sprint(s.AsSlice()) // could be made more efficient if needed
}

func (s StringSet) AsSlice() []string {
	ret := make([]string, 0, len(s))
	for k, _ := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s StringSet) Contains(i ...string) bool {
	for _, val := range i {
		if _, ok := s[val]; !ok {
			return false
		}
	}
	return true
}

func (s StringSet) Size() int {
	return len(s)
}

func (s StringSet) Difference(other StringSet) StringSet {
	//_ = other.(*threadUnsafeSet)

	difference := NewStringSet()
	for elem := range s {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return difference
}
