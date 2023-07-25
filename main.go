package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Download a url to a local file
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Download images from a catalogue
func downloadCatalogueImages() {
	lots := make(map[int]int)
	var err error
	var lotId int
	var fileUrl, fileName string

	// Open catalog xlsx file
	f, err := excelize.OpenFile("catalog.xlsx")
	if err != nil {
		log.Fatalln("Cannot find catalog.xlsx")
		return
	}

	// Create catalog folder
	if err := os.MkdirAll("catalog", os.ModePerm); err != nil {
		log.Fatalln("Cannot create folder for images")
		return
	}

	// Get all the rows in lots sheet
	rows, err := f.GetRows("lots")
	if err != nil {
		log.Fatalln("Cannot find 'lots' sheet")
		return
	}

	for i, row := range rows {
		// Skip title row
		if i == 0 {
			continue
		}

		// Row parsing
		rawId := strings.Trim(row[0], " ")
		if lotId, err = strconv.Atoi(rawId); err != nil {
			log.Fatalf("Line %d, col 1: invalid value (expected number): %x", i+1, rawId)
			return
		}

		fileUrl = row[2]
		if fileUrl != "" {
			if v, ok := lots[lotId]; ok {
				lots[lotId] = v + 1
			} else {
				lots[lotId] = 1
			}

			fileName = fmt.Sprintf("%d-%d.jpg", lotId, lots[lotId]) //TODO: Check format with bizdev

			// Download
			err := downloadFile("catalog/"+fileName, fileUrl)
			if err != nil {
				log.Panicf("Cannot download %s", fileName)
			}
			fmt.Println("Downloaded: " + fileName)
		}

	}
	fmt.Println(lots)
}

func main() {
	downloadCatalogueImages()
}
