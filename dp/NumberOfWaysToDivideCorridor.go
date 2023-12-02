package dp

func numberOfWays(corridor string) int {
	const MOD = 1_000_000_007
	valid, invalid, seats := 0, 0, 0
	for _, letter := range corridor {
		if letter == 'S' {
			seats++
		}
		if seats == 2 && letter == 'S' {
			valid = invalid
		} else if seats == 2 && letter == 'P' {
			invalid += valid
			invalid %= MOD
		} else if seats == 3 {
			seats = 1
		}
	}
	if seats == 1 {
		return 0
	} else {
		return valid
	}
}
