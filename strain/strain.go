package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(predicate func(int) bool) (result Ints) {
	for _, value := range ints {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return
}
func (list Lists) Keep(predicate func([]int) bool) (result Lists) {
	for _, value := range list {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return
}
func (string Strings) Keep(predicate func(string) bool) (result Strings) {
	for _, value := range string {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return
}

func (ints Ints) Discard(predicate func(int) bool) (result Ints) {
	for _, value := range ints {
		if !predicate(value) {
			result = append(result, value)
		}
	}
	return
}
