package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	Maintainers []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"maintainers"`
	Company     string `yaml:"company"`
	Website     string `yaml:"website"`
	Source      string `yaml:"source"`
	License     string `yaml:"license"`
	Description string `yaml:"description"`
}

func GetReuqestPayLoadAsMetadataEntry(w http.ResponseWriter, r *http.Request) *Metadata {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var metadataEntry = new(Metadata)

	err = yaml.UnmarshalStrict([]byte(reqBody), &metadataEntry)
	if err != nil {
		http.Error(w, "invalid yaml", http.StatusUnprocessableEntity)
	}

	fmt.Printf("Title: %v\n", metadataEntry.Title)
	fmt.Printf("Version: %v\n", metadataEntry.Version)

	return metadataEntry
}
