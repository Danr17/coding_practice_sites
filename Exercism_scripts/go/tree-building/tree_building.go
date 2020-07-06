package tree

import (
	"errors"
	"sort"
)

//Record is a struct containing the ID and Parent for the tree
type Record struct {
	ID     int
	Parent int
}

//Node holds the tree
type Node struct {
	ID       int
	Children []*Node
}

//Build is the function for constructing the tree
func Build(records []Record) (*Node, error) {

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	var tree = map[int]*Node{}

	for idx, r := range records {
		if idx != r.ID || r.ID < r.Parent || r.ID > 0 && r.ID == r.Parent {
			return nil, errors.New("not in sequence or has invaid parent")
		}
		subtree := &Node{ID: r.ID}
		tree[r.ID] = subtree

		if r.ID != 0 {
			p := tree[r.Parent]
			p.Children = append(p.Children, tree[r.ID])
		}

	}
	return tree[0], nil
}
