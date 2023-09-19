package iteration

func Repeat(char string, reps int) string {
	var repeated string
	for i := 0; i < reps; i++ {
		repeated += char
	}
	return repeated
}
