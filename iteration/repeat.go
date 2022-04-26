package iteration

// Repeat repeats the "character" "repeatedCount" times
func Repeat(character string, repeatedCount int) string {
	var repeated string

	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}

	return repeated
}
