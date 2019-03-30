package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"regexp"
)

type scoop struct {
	Homepage    string `json:"homepage"`
	Description string `json:"description"`
	License     string `json:"license"`
	Version     string `json:"version"`
	URL         string `json:"url"`
	Bin         string `json:"bin"`
	Hash        string `json:"hash"`
}

func main() {
	urlPattern := flag.String("urlpattern", ".+", "TODO")
	versionPattern := flag.String("versionpattern", ".+", "TODO")

	flag.Parse()
	file := flag.Arg(0)

	urlRegex := regexp.MustCompile(*urlPattern)
	versionRegex := regexp.MustCompile(*versionPattern)

	if file == "" {
		log.Fatalf("Usage: %s <file>\n", os.Args[0])
	}

	openFile, openError := os.Open(file)
	if openError != nil {
		log.Fatalf("Error opening manifest. %s\n", openError)
	}
	decoder := json.NewDecoder(openFile)

	target := &scoop{}
	decoder.Decode(target)

	if !urlRegex.MatchString(target.URL) {
		log.Fatalf("ERROR!\n\tURL '%s' doesn't match pattern '%s'\n", target.URL, *urlPattern)
	}

	if !versionRegex.MatchString(target.Version) {
		log.Fatalf("ERROR!\n\tVersion '%s' doesn't match pattern '%s'\n", target.Version, *versionPattern)
	}

	if len(target.Hash) != 64 {
		log.Fatalf("ERROR!\n\tHash '%s' seems to be incorrect\n", target.Hash)
	}

	if len(target.Bin) == 0 {
		log.Fatalf("ERROR!\n\tBin value is empty\n")
	}

	if len(target.Description) == 0 {
		log.Fatalf("ERROR!\n\tDescription value is empty\n")
	}

	if len(target.Homepage) == 0 {
		log.Fatalf("ERROR!\n\tHomepage value is empty\n")
	}

	if len(target.License) == 0 {
		log.Fatalf("ERROR!\n\tLicense value is empty\n")
	}
}
