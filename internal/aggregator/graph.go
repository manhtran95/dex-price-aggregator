package aggregator

import (
	"github.com/ethereum/go-ethereum/common"
)

type Graph struct {
	nodes map[common.Address]*Node
	edges []*Edge
}

type Node struct {
	token common.Address
	edges []*Edge
}

type Edge struct {
	poolAddress common.Address
	token0      common.Address
	token1      common.Address
	dex         string
	fee         uint32
}

type Path struct {
	Tokens []common.Address // [WETH, DAI, USDC] - addresses only, token metadata resolved at quote time
	Edges  []*Edge          // Specific pool for each hop
}

type PoolInfo struct {
	Address common.Address
	Token0  common.Address
	Token1  common.Address
	DEX     string
	Fee     uint32
}

func NewGraph() *Graph {
	g := &Graph{
		nodes: make(map[common.Address]*Node),
		edges: []*Edge{},
	}

	for _, pool := range KnownPools {
		g.AddPool(pool)
	}

	return g
}

func (g *Graph) AddPool(pool PoolInfo) {
	edge := &Edge{
		poolAddress: pool.Address,
		token0:      pool.Token0,
		token1:      pool.Token1,
		dex:         pool.DEX,
		fee:         pool.Fee,
	}

	g.edges = append(g.edges, edge)

	// Create/get nodes
	if _, exists := g.nodes[pool.Token0]; !exists {
		g.nodes[pool.Token0] = &Node{token: pool.Token0}
	}
	if _, exists := g.nodes[pool.Token1]; !exists {
		g.nodes[pool.Token1] = &Node{token: pool.Token1}
	}

	// Link edges to nodes
	g.nodes[pool.Token0].edges = append(g.nodes[pool.Token0].edges, edge)
	g.nodes[pool.Token1].edges = append(g.nodes[pool.Token1].edges, edge)
}

func (g *Graph) FindAllPaths(tokenIn, tokenOut common.Address, maxHops int) []Path {
	paths := []Path{}
	visited := make(map[common.Address]bool)
	currentPath := Path{
		Tokens: []common.Address{tokenIn},
		Edges:  []*Edge{},
	}

	g.dfs(tokenIn, tokenOut, currentPath, visited, &paths, maxHops)

	return paths
}

func (g *Graph) dfs(
	current, target common.Address,
	currentPath Path,
	visited map[common.Address]bool,
	allPaths *[]Path,
	maxHops int,
) {
	// Too many hops
	if len(currentPath.Tokens) > maxHops+1 {
		return
	}

	// Found target
	if current == target {
		// Make a copy of the path
		pathCopy := Path{
			Tokens: make([]common.Address, len(currentPath.Tokens)),
			Edges:  make([]*Edge, len(currentPath.Edges)),
		}
		copy(pathCopy.Tokens, currentPath.Tokens)
		copy(pathCopy.Edges, currentPath.Edges)
		*allPaths = append(*allPaths, pathCopy)
		return
	}

	visited[current] = true
	defer func() { visited[current] = false }()

	node := g.nodes[current]
	if node == nil {
		return
	}

	// Try all edges from current node
	for _, edge := range node.edges {
		// Determine next token
		nextToken := edge.token1
		if current == edge.token1 {
			nextToken = edge.token0
		}

		// Skip if already visited
		if visited[nextToken] {
			continue
		}

		// Add edge and token to path
		newPath := Path{
			Tokens: append(currentPath.Tokens, nextToken),
			Edges:  append(currentPath.Edges, edge),
		}

		// Recurse
		g.dfs(nextToken, target, newPath, visited, allPaths, maxHops)
	}
}
