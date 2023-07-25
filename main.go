package main

import (
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
	var err error
	var fileName string
	var nbDownloads int = 0

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
		if i == 0 || len(row) < 3 {
			continue
		}

		// Row parsing
		id := strings.Trim(row[0], " ")
		ext := strings.Trim(row[1], " ")
		fileUrl := row[2]
		if _, err = strconv.Atoi(id); err != nil {
			log.Fatalln("Line " + strconv.Itoa(i+1) + ", col 1: invalid value (expected number): " + id)
			return
		}

		if ext == "" {
			fileName = id + "_1.jpg"
		} else {
			fileName = id + "-" + ext + "_1.jpg"
		}

		// Download
		err := downloadFile("catalog/"+fileName, fileUrl)
		if err != nil {
			log.Panicln("Cannot download " + fileName)
		}
		log.Println("Downloaded: " + fileName)
		nbDownloads = nbDownloads + 1
	}

	log.Println("Downloaded images: " + strconv.Itoa(nbDownloads))
}

func main() {
	downloadCatalogueImages()
}
