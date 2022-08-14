package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"zaunic-lite/models"

	"gopkg.in/yaml.v2"
)

func getWorksheetPath(worksheetName string) string {
	worksheetPath := filepath.Join(
		"worksheets",
		fmt.Sprintf("%s.yml", worksheetName),
	)
	return worksheetPath
}

func getWorkItems(worksheetName string) (models.Worksheet, error) {
	worksheetPath := getWorksheetPath(worksheetName)
	yamlFile, err := ioutil.ReadFile(worksheetPath)
	var wst models.Worksheet
	if err != nil {
		return wst, err
	}
	err = yaml.Unmarshal(yamlFile, &wst)
	if err != nil {
		fmt.Println("Error reading yaml.")
		return wst, err
	}
	return wst, nil
}

func ProcessWorksheet(worksheetName string) {
	var wst models.Worksheet
	wst, _ = getWorkItems(worksheetName)
	for _, i := range wst.Items {
		fmt.Println(i.Name)
	}
}
