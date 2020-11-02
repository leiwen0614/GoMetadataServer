package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

		/*
			entryYaml, err := yaml.Marshal(entry)
			check(err)
			entryJson, err := yaml.YAMLToJSON(entryYaml)
			check(err)

				fmt.Printf("======\n")
				fmt.Printf(string(entryJson))
				fmt.Printf("======\n")
		*/

		//fmt.Printf("Title: %v, Version: %v, Maintainer: %v Company: %v\n", entry.Title, entry.Version, entry.Maintainers, entry.Company)
		count++
		return true
	})
	fmt.Println("Found %v metadata entry\n", count)

	json_bytes, _ := json.Marshal(rt)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "%s\n", json_bytes)

}
