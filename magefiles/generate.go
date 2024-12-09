//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

// data to be used in the templates
var templateData = map[string]string{
	"RepoName":     "github.com/insixio/uo-launcher",
	"WailsVersion": "v2.9.2",
	"ProjectName":  "Insixio UO Launcher",
	"BinaryName":   "insixio-launcher",
	"AuthorName":   "Insixio",
	"AuthorEmail":  "ask+launcher@in6.io",
}

// list of template file names
var templateFiles = []string{
	"sometemplate.tmpl.go",
}

// --------------------------------------------------------------------------
// Generate namespace
type Generate mg.Namespace

// This script generates the necessary files for the project based on go templates.
func (Generate) Project() error {
	// walk through the base directory and its subdirectories
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip directories
		if d.IsDir() {
			return nil
		}

		// check if the file name matches any in the targetTemplates list
		fileName := filepath.Base(path)
		if contains(templateFiles, fileName) {
			// parse the template
			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			// create the output file by removing `.tmpl` from the name
			outputFileName := strings.Replace(path, ".tmpl", "", 1) // remove `.tmpl`
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			// execute the template with the data
			err = tmpl.Execute(outputFile, templateData)
			if err != nil {
				return err
			}

			fmt.Println("Generated:", outputFileName)

			// deletes the template file
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error traversing path %q: %w", ".", err)
	}

	return nil
}

// --------------------------------------------------------------------------
// utils

// Helper function to check if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
