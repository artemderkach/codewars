// https://www.codewars.com/kata/first-variation-on-caesar-cipher/train/go
package kata

var alphabetLength = 26

func MovingShift(s string, shift int) []string {
	currentShift := shift
	cipheredString := ""
	for _, char := range s {
		cipheredString += string(shiftedChar(char, currentShift))
		currentShift++
	}
	return stringToArray(cipheredString)
}

func stringToArray(s string) []string {
	result := make([]string, 5)
	arrayLength := 5
	n := 0
	if (len(s) % arrayLength) == 0 {
		n = int(len(s) / arrayLength)
	} else {
		n = int(len(s)/arrayLength) + 1
	}
	for i := 0; i < arrayLength; i++ {
		if (i+1)*n < len(s) {
			result[i] = s[i*n : (i+1)*n]
			continue
		}
		result[i] = s[i*n : len(s)]
		break
	}
	return result
}

func DemovingShift(arr []string, shift int) string {
	currentShift := -shift
	decipheredString := ""
	for _, s := range arr {
		for _, char := range s {
			decipheredString += string(shiftedChar(char, currentShift))
			currentShift--
		}
	}
	return decipheredString
}

func shiftedChar(c rune, shift int) rune {
	if c >= 'a' && c <= 'z' {
		letterNumber := int(c) - int('a')
		relativeShift := (letterNumber + shift) % alphabetLength
		if relativeShift < 0 {
			return rune(int('z') + relativeShift + 1)
		}
		return rune(int('a') + relativeShift)
	}
	if c >= 'A' && c <= 'Z' {
		letterNumber := int(c) - int('A')
		relativeShift := (letterNumber + shift) % alphabetLength
		if relativeShift < 0 {
			return rune(int('Z') + relativeShift + 1)
		}
		return rune(int('A') + relativeShift)
	}
	return c
}
