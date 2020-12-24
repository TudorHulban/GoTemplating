package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

// file contains handling of data file import in CSV format

// exportProducts Helper exports products to file.
func exportProducts(p []*Product, path string) error {
	dataFile, errOpen := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if errOpen != nil {
		return errOpen
	}
	defer dataFile.Close()

	return gocsv.MarshalFile(&p, dataFile)
}

func importProducts(path string) ([]*Product, error) {
	dataFile, errOpen := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if errOpen != nil {
		return nil, errOpen
	}
	defer dataFile.Close()

	result := []*Product{}
	return result, gocsv.UnmarshalFile(dataFile, &result)
}
