package iterations

// Repears entered value specified number of times
func Repeat(times int, value string) string {
	s := ""

	for i := 0; i < times; i++ {
		s += value
	}

	return s
}