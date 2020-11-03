package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func AddOneEndTry(entry *Metadata) {
	mdb.Put(entry)
}

func IteratorDatabase(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Iterating over all metadata: \n")
	count := 0
	rt := make([]*Metadata, 0)

	mdb.Ascend(func(indexer interface{}) bool {
		entry := indexer.(*Metadata)
		rt = append(rt, entry)

		count++
		return true
	})
	fmt.Println("Found %v metadata entry\n", count)

	json_bytes, _ := json.Marshal(rt)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "%s\n", json_bytes)

}

func QueryDatabase(w http.ResponseWriter, r *http.Request) {
	rt := make([]*Metadata, 0)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var metadataEntry = new(Metadata)

	err = yaml.UnmarshalStrict([]byte(reqBody), &metadataEntry)
	if err != nil {
		http.Error(w, "invalid yaml", http.StatusUnprocessableEntity)
	}
	fmt.Printf("query Title: %v\n", metadataEntry.Title)
	fmt.Printf("query Version: %v\n", metadataEntry.Version)

	var indexers []interface{}

	if len(metadataEntry.Title) != 0 && len(metadataEntry.Version) != 0 {
		indexers = mdb.In("Title", "Version").Lookup(metadataEntry.Title, metadataEntry.Version)
	} else if len(metadataEntry.Title) != 0 && len(metadataEntry.Version) == 0 {
		indexers = mdb.In("Title").Lookup(metadataEntry.Title)
	} else if len(metadataEntry.Title) == 0 && len(metadataEntry.Version) != 0 {
		indexers = mdb.In("Version").Lookup(metadataEntry.Version)
	}

	//indexers = mdb.In("Title", "Version").Lookup(metadataEntry.Title, metadataEntry.Version)

	fmt.Printf("indexers len = %d\n", len(indexers))
	for _, indexer := range indexers {
		entry := indexer.(*Metadata)
		fmt.Printf("found found something: %v\n", metadataEntry.Title)
		rt = append(rt, entry)
	}

	json_bytes, _ := json.Marshal(rt)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "%s\n", json_bytes)

}
