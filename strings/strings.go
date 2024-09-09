package strings

// Truncate truncates the string to maxRunes characters
func Truncate(str string, maxRunes uint) string {
	runes := []rune(str)
	if len(runes) > int(maxRunes) {
		return string(runes[:maxRunes])
	}

	return str
}
