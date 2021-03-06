package main

import (
	//    "bufio"
	"fmt"
	"net/http"

	//    "io"
	"io/ioutil"
	//    "os"
	"gopkg.in/yaml.v2"
)

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
		GetReuqestPayLoadAsMetadataEntry(w, r)
	case "POST":
		fmt.Printf("Go server handling POST request: \n")
		var entry *Metadata = GetReuqestPayLoadAsMetadataEntry(w, r)
		fmt.Printf("Title = %v\n", entry.Title)
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
