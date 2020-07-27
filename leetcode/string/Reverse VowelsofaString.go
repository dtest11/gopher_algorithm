package leetcode

// ReverseVowels 翻转字符串中的元音字母, 2 2 交换 从left 和right 交换
func ReverseVowels(s string) string {
	left, right := 0, len(s)-1

	for left < right {
		if !isVowels(s[left]) {
			left++
			continue
		}
		if !isVowels(s[right]) {
			right--
			continue
		}

		// left , right 都是元音字符
		s = swap(s, left, right)
		left++
		right--
	}
	return s
}

// swap 交换字符
func swap(str string, indexA, indexB int) string {
	runes := []rune(str)
	runes[indexA], runes[indexB] = runes[indexB], runes[indexA]
	return string(runes)
}

func isVowels(inputByte byte) bool {
	input := string(inputByte)
	bucket := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, val := range bucket {
		if val == input {
			return true
		}
	}
	return false
}
