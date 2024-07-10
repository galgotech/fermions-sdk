package builder

import "github.com/galgotech/fermions-sdk/internal/graph"

type DurationBuilder struct {
	root *graph.Node
}

func (b *DurationBuilder) SetSeconds(seconds int) *DurationBuilder {
	b.root.Edge("seconds").SetInt(seconds)
	return b
}

func (b *DurationBuilder) GetSeconds() int {
	return b.root.Edge("seconds").GetInt()
}

func NewDurationBuilder(root *graph.Node) *DurationBuilder {
	return &DurationBuilder{
		root: root,
	}
}
