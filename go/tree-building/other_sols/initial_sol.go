package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}


// ValidRecords ensures records have sequential parents, that
// a root node exists, and there are no duplicate records
func ValidRecords(records []Record) error {
	UniqueRecords := map[Record]bool{}
	ContainsDuplicates := false
	RootNodeExists := false
	HasSequentialParents := true
	CurrentParent := 0
	for _, rec := range records {
		if _, ok := UniqueRecords[rec]; ok {
			ContainsDuplicates = true
			break
		} else {
			UniqueRecords[rec] = true
		}

		// Check if the records contain a root node
		if rec.ID == 0 && rec.Parent == 0 {
			RootNodeExists = true
		}
		// Check if records have continuous parents
		if rec.Parent > CurrentParent + 1 {
			HasSequentialParents = false
			break
		} else {
			CurrentParent = rec.Parent
		}
	}

	switch {
	case !RootNodeExists:
		return errors.New("No root node exists")
	case !HasSequentialParents:
		return errors.New("Non-sequential parents")
	case ContainsDuplicates:
		return errors.New("Records contain duplicatess")
	default:
		return nil
	}
}

func Build(records []Record) (*Node, error) {
	// Check there are any records
	if len(records) == 0 {
		return nil, nil
	}
	// Sort records by Parent, then ID
	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent < records[j].Parent {
			return true
		}
		return records[i].ID < records[j].ID
	})

	err := ValidRecords(records)
	if err != nil {
		return nil, err
	}

	n := Node{}
	NodePointer := &n
	for _, item := range records {
		if item.Parent != 0 && item.Parent >= item.ID {
			return nil, errors.New("Parent ID must be lower than item ID except for root node")
		} else if item.ID != 0 && item.Parent == 0 {
			// non-root node, with parent as root node
			NewNode := Node{item.ID, nil}
			NodePointer.Children = append(NodePointer.Children, &NewNode)
		} else if item.ID != 0 && item.Parent != 0{
			// non-root node, with parent as non-root node
			NewNode := Node{item.ID, nil}
			NodePointer = n.Children[item.Parent-1]
			NodePointer.Children = append(NodePointer.Children, &NewNode)
		} else if item.ID == 0 && item.Parent == 0 {
			// root node
			NodePointer.ID = item.ID
		}
	}
	NodePointer = &n
	return NodePointer, nil
}
