package builder

import "github.com/galgotech/fermions-sdk/internal/graph"

type UseBuilder struct {
	root *graph.Node
}

func NewUseBuilder(root *graph.Node) *UseBuilder {
	return &UseBuilder{
		root: root,
	}
}
