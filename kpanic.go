package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

const (
	diagnosticsFolder = "/Library/Logs/DiagnosticReports/"
)

func listPanics(dir string) ([]string, error) {
	result := []string{}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return result, err
	}

	for _, f := range files {
		name := f.Name()
		if strings.HasPrefix(name, "Kernel") && strings.HasSuffix(name, "panic") {
			result = append(result, name)
		}
	}

	return result, nil
}

func readPanic(filename string) (string, error) {
	fullPath := path.Join(diagnosticsFolder, filename)
	c, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(c), nil
}

func displayPanic(filename string) error {
	content, err := readPanic(filename)

	if err != nil {
		return err
	}

	fmt.Println(content)
	return nil
}
