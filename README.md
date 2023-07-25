# Imagine

Download images from URLs in an Excel catalog into a "catalog" folder

## Pre-requisites

 - The Excel file must be named "catalog.xlsx"
 - The Excel tab must be named "lots"
 - Inside this tab, 3 columns are expected
	 - Lot number (must be a number)
	 - Lot extension `ex: bis, ter`
	 - Image URL


## For devs
- Download dependencies: `go mod download`
- Build win exe: `GOOS=windows GOARCH=amd64 go build imagine`

## Run
- Put imagine.exe and the Excel file in the same folder
- Run imagine.exe (The images should be downloaded into a subfolder called "catalog")