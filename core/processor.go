package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"zaunic-lite/core/models"

	"gopkg.in/yaml.v2"
)

func getWorksheetPath(
	projectName string,
	worksheetName string,
) string {
	worksheetPath := filepath.Join(
		"projects",
		projectName,
		"worksheets",
		fmt.Sprintf("%s.yml", worksheetName),
	)
	return worksheetPath
}

func getWorkItems(
	projectName string,
	worksheetName string,
) (models.Worksheet, error) {
	worksheetPath := getWorksheetPath(
		projectName,
		worksheetName,
	)
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

func ProcessWorksheet(
	projectName string,
	worksheetName string,
) {
	var wst models.Worksheet
	wst, _ = getWorkItems(
		projectName,
		worksheetName,
	)

	//Get secrets from worksheet.
	getCloudSecrets(wst.Secrets)

	//Added processing using waitgroup.
	var wg sync.WaitGroup
	wg.Add(len(wst.Items))
	for _, i := range wst.Items {
		go func(
			projectName string,
			i models.Item,
		) {
			defer wg.Done()
			Generate(projectName, i)
		}(projectName, i)

	}
	wg.Wait()
}
