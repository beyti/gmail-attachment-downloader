package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	labelPtr := flag.String("label", "", "label to process")
	outFolderPtr := flag.String("out", fmt.Sprintf("./%s", "attachments"), "output folder")

	flag.Parse()

	log.Printf("option label: %v ", *labelPtr)
	log.Printf("option out: %v ", *outFolderPtr)

	downloadFolderPath := fmt.Sprintf("%s/%s", *outFolderPtr, *labelPtr)
	err := os.MkdirAll(downloadFolderPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	downloadByLabel(*labelPtr, downloadFolderPath)

}
