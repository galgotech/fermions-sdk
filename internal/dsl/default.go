package dsl

import (
	"github.com/galgotech/fermions-sdk/graph"
)

func ApplyDefault(node *graph.Node) error {
	lookup := node.Lookup("do.*.*.call=http.with.content")
	for _, node := range lookup.List() {
		if !node.HasValue() {
			node.SetString("content")
		}
	}

	lookup = node.Lookup("do.*.*.then")
	for _, node := range lookup.List() {
		if !node.HasValue() {
			node.SetString("continue")
		}
	}

	lookup = node.Lookup("do.*.*.fork")
	for _, node := range lookup.List() {
		if !node.Edge("compete").HasValue() {
			node.SetBool(false)
		}
	}

	lookup = node.Lookup("do.*.*.run.workflow")
	for _, node := range lookup.List() {
		if !node.Edge("version").HasValue() {
			node.Edge("version").SetString("latest")
		}
	}

	lookup = node.Lookup("do.*.*.catch")
	for _, node := range lookup.List() {
		if !node.Edge("catch").HasValue() {
			node.Edge("catch").SetString("error")
		}
	}

	lookup = node.Lookup("evaluate.language")
	if !lookup.Empty() {
		node.Edge("evaluate").Edge("language").SetString("jq")
	}

	lookup = node.Lookup("evaluate.mode")
	if !lookup.Empty() {
		node.Edge("evaluate").Edge("mode").SetString("strict")
	}

	return nil
}
