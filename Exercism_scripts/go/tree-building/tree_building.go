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

	for idx, record := range records {
		if idx != record.ID || record.ID < record.Parent || record.ID > 0 && record.ID == record.Parent {
			return nil, errors.New("not in sequence or has invaid parent")
		}
		subtree := &Node{ID: record.ID}
		tree[record.ID] = subtree

		if record.ID != 0 {
			if parent, ok := tree[record.Parent]; ok {
				parent.Children = append(parent.Children, subtree)
			} else {
				return nil, errors.New("parent node does not exist")
			}
		}

	}
	return tree[0], nil
}
