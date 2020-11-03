package main

import (
	//    "bufio"
	"fmt"
	"net/http"

	//    "io"
	"io/ioutil"
	//    "os"
	"github.com/nedscode/memdb"
	"gopkg.in/yaml.v2"
)

var mdb memdb.Storer = memdb.NewStore().PrimaryKey("Title", "Version").CreateIndex("Title").CreateIndex("Version").CreateIndex("Company")

func main() {
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
		if r.Body == http.NoBody {
			fmt.Printf("No Request body, return all metadata entry in json array. \n")
			IteratorDatabase(w, r)
		} else {
			fmt.Printf("Found request body, query database. \n")
			QueryDatabase(w, r)
		}

	case "POST":
		fmt.Printf("Go server handling POST request: \n")
		var entry *Metadata = GetReuqestPayLoadAsMetadataEntry(w, r)
		fmt.Printf("Title = %v\n", entry.Title)

		AddOneEndTry(entry)
	}

}

func readOneMetadataEntry(filePath string) {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	//    fmt.Print(string(dat))

	dat_slice := []byte(dat)

	var entry Metadata

	err = yaml.UnmarshalStrict(dat_slice, &entry)
	check(err)

	fmt.Printf("%v\n", entry.Title)
	fmt.Printf("%v, %v\n", entry.Maintainers[0].Name, entry.Maintainers[0].Email)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
