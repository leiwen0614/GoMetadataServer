package main

import (
	"fmt"
)

/*
// Create the DB schema
schema := &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"metadata": &memdb.TablesSchema{
			Name: "metadata",
			Indexes: map[string]*memdb.IndexSchema{
				"title": &memdb.IndexSchema{
					Name:    "title",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Title"},
				}
			}
		}
	}
}


func initDatabase() {
	// Create the DB schema

	// Create a write transaction
	txn := db.Txn(true)

	var entry *Metadata = readOneMetadataEntry("test_data/valid1.yml")

	txn.Insert("metadata", entry)

	ListAllMetadata(db)
	//allData, err = txn.Get("metadata", "title")

}
*/

func AddOneEntrytoDatabase(entry *Metadata) {
	// Create a write transaction
	txn := db.Txn(true)
	//var entry *Metadata = readOneMetadataEntry("test_data/valid1.yml")

	txn.Insert("metadata", entry)

	// Commit the transaction
	txn.Commit()
}

func ListDatabase() {
	txn := db.Txn(false)
	defer txn.Abort()
	allData, err := txn.Get("metadata", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the entry title:")
	for obj := allData.Next(); obj != nil; obj = allData.Next() {
		p := obj.(*Metadata)
		fmt.Printf("  %s\n", p.Title)
		fmt.Printf("  %s\n", p.Version)
	}
}
