// Mostly taken from inspecting https://exercism.io/tracks/go/exercises/tree-building/solutions/30309c4a90504287a419bea876cfd707

package tree

import (
	"errors"
	"sort"
)

// Record contains a single data entry
type Record struct {
	ID     int
	Parent int
}

// Node is an entity in a tree, containing an ID and pointers to other nodes
type Node struct {
	ID       int
	Children []*Node
}

// Build implements tree-building logic for a set of unsorted records
func Build(records []Record) (*Node, error) {
	// Check there are any records
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make([]*Node, len(records))
	for ind, rec := range records {
		if (rec.ID != ind) || (rec.Parent > rec.ID) || (rec.ID > 0 && rec.Parent == rec.ID) {
			return nil, errors.New("invalid set of records")
		}

		nodes[ind] = &Node{ID: rec.ID}
		if ind != 0 {
			nodes[rec.Parent].Children = append(nodes[rec.Parent].Children, nodes[ind])
		}
	}

	return nodes[0], nil
}
