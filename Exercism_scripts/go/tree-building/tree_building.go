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

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	tree := Node{}

	for idx, record := range records {
		if idx != record.ID {
			return nil, errors.New("Invalid tree")
		}
		if record.ID < record.Parent {
			return nil, errors.New("Invalid tree, higher id parent of lower id")
		}
		if record.ID > 0 && record.ID == record.Parent {
			return nil, errors.New("Invalid tree, inner cycle")
		}
		if record.ID == 0 && record.Parent == 0 {
			tree.ID = 0
			continue
		}
		if err := findAndAdd(&tree, record); err != nil {
			return nil, err
		}
	}

	return &tree, nil

}

func findAndAdd(tree *Node, record Record) error {

	if tree.ID == record.Parent {
		for _, node := range tree.Children {
			if node.ID == record.ID {
				return errors.New("duplicate ID")
			}
		}
		tree.Children = append(tree.Children, &Node{ID: record.ID})
		return nil
	}
	if len(tree.Children) == 0 {
		return nil
	}
	for i := range tree.Children {
		findAndAdd(tree.Children[i], record)
	}
	return nil
}
