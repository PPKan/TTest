package iteration

func repeat(char string) string {
	var repeated string
	for i := 0; i < 3; i++ {
		repeated += char
	}
	return repeated
}
