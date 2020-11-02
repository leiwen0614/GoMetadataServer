package main

import (
	//    "bufio"
	"fmt"
	"net/http"

	//    "io"
	"io/ioutil"
	//    "os"
	"github.com/hashicorp/go-memdb"
	"gopkg.in/yaml.v2"
)

// Init the in-memory database as global variable
var schema *memdb.DBSchema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"metadata": &memdb.TableSchema{
			Name: "metadata",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Title"},
				},
				"version": &memdb.IndexSchema{
					Name:    "version",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Version"},
				},
			},
		},
	},
}

// Create a new database instance
var db, err = memdb.NewMemDB(schema)

func main() {
	//http.HandleFunc("/", HelloServer)
	//http.ListenAndServe(":8080", nil)
	//readOneMetadataEntry("test_data/valid1.yml")

	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)

	//readOneMetadataEntry("test_data/valid1.yml")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	fmt.Printf("Entering requestHandler function...\n")

	switch r.Method {
	case "GET":
		fmt.Printf("Go server handling GET request: \n")
		//AddOneEntrytoDatabase()
		ListDatabase()
	case "POST":
		fmt.Printf("Go server handling POST request: \n")
		var entry *Metadata = GetReuqestPayLoadAsMetadataEntry(w, r)
		//fmt.Printf("Title = %v\n", entry.Title)
		AddOneEntrytoDatabase(entry)
	}

}

func readOneMetadataEntry(filePath string) *Metadata {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	//    fmt.Print(string(dat))

	dat_slice := []byte(dat)

	var entry Metadata

	err = yaml.UnmarshalStrict(dat_slice, &entry)
	check(err)

	fmt.Printf("%v\n", entry.Title)
	fmt.Printf("%v, %v\n", entry.Maintainers[0].Name, entry.Maintainers[0].Email)

	return &entry
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
