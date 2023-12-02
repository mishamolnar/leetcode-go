package math

import "math"

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy && t > 0 {
		return false
	}
	if sx > fx {
		sx, fx = fx, sx
	}
	if sy > fy {
		sy, fy = fy, sy
	}
	routeX, routeY := fx-sx, fy-sy
	return (int(math.Min(float64(routeX), float64(routeY))) + int(math.Abs(float64(routeX)-float64(routeY)))) <= t
}
