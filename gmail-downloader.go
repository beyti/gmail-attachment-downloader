/*
   Copyright (C) 2021-present Mirko Perillo and contributors

   This file is part of gmail-downloader.

   gmail-downloader is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   gmail-downloader is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with ts-converter.  If not, see <http://www.gnu.org/licenses/>.
*/

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
