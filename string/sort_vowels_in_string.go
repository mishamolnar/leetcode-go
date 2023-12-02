package string

func sortVowels(s string) string {
	arr := []byte(s)
	vowelIndex := map[byte]int{
		'A': 0,
		'E': 1,
		'I': 2,
		'O': 3,
		'U': 4,
		'a': 5,
		'e': 6,
		'i': 7,
		'o': 8,
		'u': 9,
	}
	counts := [10]int{}
	for ind, element := range arr {
		if indx, ok := vowelIndex[element]; ok {
			counts[indx]++
			arr[ind] = '#'
		}
	}
	vowels := [10]byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	pointer := 0
	for ind, element := range arr {
		if element == '#' {
			for counts[vowelIndex[vowels[pointer]]] == 0 {
				pointer++
			}
			arr[ind] = vowels[pointer]
			counts[vowelIndex[vowels[pointer]]]--
		}
	}
	return string(arr)
}
