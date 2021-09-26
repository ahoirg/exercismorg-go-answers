package accumulate

func Accumulate(input []string, operation func(string) string) (result []string) {
	for _, e := range input {
		result = append(result, operation(e))
	}
	return result
}
