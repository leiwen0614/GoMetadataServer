package main

import "fmt"

/*
func (i *Metadata) Less(other memdb.Indexer) bool {
	switch o := other.(type) {
	case *Metadata:
		if i.Title < o.Title {
			return true
		}
		if i.Title > o.Title {
			return false
		}
		if i.Version < o.Version {
			return true
		}
		if i.Version > o.Version {
			return false
		}
		return false
	}
	return memdb.Unsure(i, other)

}

func (i *Metadata) GetField(field string) string {
	switch field {
	case "Title":
		return i.Title
	case "Version":
		return i.Version
	default:
		return "" // Indicates should not be indexed
	}
}
*/

func AddOneEndTry(entry *Metadata) {
	mdb.Put(entry)
}

func IteratorDatabase() {
	fmt.Println("Iterating over all metadata: \n")
	count := 0
	mdb.Ascend(func(indexer interface{}) bool {
		entry := indexer.(*Metadata)
		fmt.Printf("%s %s ($%d rrp)\n", entry.Title, entry.Version, entry.Company)
		count++
		return true
	})
	fmt.Println("Found %d metadata entry\n", count)
}
