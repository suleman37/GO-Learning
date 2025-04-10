package analyzer

func Combination(fileContent string) (int, int, int, int, int, int, int, int) {
	punchar := 0
	special := 0
	spaces := 0
	vowels := 0
	words := 0
	line := 0
	cons := 0
	digits := 0
	for i := 0; i < len(fileContent); i++ {
		char := fileContent[i]
		if char == ' ' {
			words++
			spaces++
		}
		if char == '\n' {
			line++
		}
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		default:
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				cons++
			}
		}
		switch char {
		case '.', ',', '\'', '(', ')', '!', '?', ';', ':', '-', '"':
			punchar++
		}
		if char >= '0' && char <= '9' {
			digits++
		}
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') &&
			char != ' ' && char != '\n' {
			special++
		}
	}
	return words, line, spaces, vowels, punchar, digits, cons, special
}