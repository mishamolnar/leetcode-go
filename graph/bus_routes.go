package graph

import "math"

func EntryPoint() {
	numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6)
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	//create map: bus stop -> index
	//then create graph where index is a vertex and edges based on same stops
	//then run bfs from starting index
	stopToRoute := make(map[int][]int)
	for index, route := range routes {
		for _, stop := range route {
			stopToRoute[stop] = append(stopToRoute[stop], index) //[[1,2,7],[3,6,7]] produces 7 -> [0, 1]
		}
	}
	routesGraph := make(map[int][]int)
	for _, routesList := range stopToRoute {
		for _, route := range routesList {
			routesGraph[route] = append(routesGraph[route], routesList...) //appending all adj routes
		}
	}
	startRoutes, endRoutes := stopToRoute[source], stopToRoute[target]
	steps := make([]int, len(routes))
	for i := range steps {
		steps[i] = -1
	}
	for i := range startRoutes {
		steps[i] = 0
	}
	for len(startRoutes) > 0 {
		nextStart := make([]int, 0)
		for _, curRoute := range startRoutes {
			for _, nextRoute := range routesGraph[curRoute] {
				if steps[nextRoute] == -1 {
					steps[nextRoute] = steps[curRoute] + 1
					nextStart = append(nextStart, nextRoute)
				}
			}
		}
		startRoutes = nextStart
	}
	res := math.MaxInt
	for _, route := range endRoutes {
		res = min(res, steps[route])
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
