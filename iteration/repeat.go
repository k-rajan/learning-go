package iteration

// Repeat thing 5 times
func Repeat(term string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += term
	}
	return repeated
}
