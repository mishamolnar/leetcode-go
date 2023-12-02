package graph

type Graph struct {
	dp [][]int
}

func Constructor(n int, edges [][]int) Graph {
	var dp [][]int = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for _, edge := range edges {
		dp[edge[0]][edge[1]] = edge[2]
	}
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] == 0 && i != j {
				dp[i][j] = 1_000_000_000
			}
		}
	}
	for curr := range dp {
		for i := range dp {
			for j := range dp {
				dp[i][j] = min(dp[i][j], dp[i][curr]+dp[curr][j])
			}
		}
	}
	return Graph{dp}
}

func (this *Graph) AddEdge(edge []int) {
	for i := range this.dp {
		for j := range this.dp {
			this.dp[i][j] = min(this.dp[i][j], this.dp[i][edge[0]]+edge[2]+this.dp[edge[1]][j])
		}
	}

}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.dp[node1][node2] < 1_000_000_000 {
		return this.dp[node1][node2]
	} else {
		return -1
	}
}
