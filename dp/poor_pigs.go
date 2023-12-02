package dp

import "math"

/*
We can understand it using induction. Let's say we need n pigs and m iterations in order to determine
poisonous bucket out of k buckets. Then by adding 1 new pig we can divide all buckets into smaller parts
(1 / k) smaller. This means that we gain 1/k advantage m times.
*/
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs, iter := 0, minutesToTest/minutesToDie
	for int(math.Pow(float64(iter+1), float64(pigs))) < buckets {
		pigs++
	}
	return pigs
}
