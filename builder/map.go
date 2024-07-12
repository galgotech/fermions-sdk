package builder

import "github.com/galgotech/fermions-sdk/graph"

type MapBuilder struct {
	root *graph.Node
}

func (b *MapBuilder) Set(name string, value string) {
	b.root.Edge(name).SetString(value)
}

func (b *MapBuilder) Get(name string) string {
	return b.root.Edge(name).GetString()
}

func NewMapBuilder(root *graph.Node) *MapBuilder {
	return &MapBuilder{
		root: root,
	}
}
