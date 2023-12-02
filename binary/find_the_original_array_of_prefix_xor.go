package binary

func findArray(pref []int) []int {
	res := make([]int, len(pref))
	for i := range pref {
		if i == 0 {
			res[0] = pref[0]
		} else {
			res[i] = pref[i-1] ^ pref[i]
		}
	}
	return res
}
