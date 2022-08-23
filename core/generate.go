package core

import (
	"log"
	"os"
	"path/filepath"
	"zaunic-lite/core/models"

	"gopkg.in/yaml.v2"
)

func getTemplateData(templateData []byte) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(templateData), &m); err != nil {
		panic(err)
	}
	return m
}

func writeOutput(filepath string, content string) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(content)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func Generate(
	projectName string,
	wsi models.Item,
) {
	templateDataFilePath := filepath.Join(
		"_projects",
		projectName,
		"data",
		wsi.Data,
	)

	templateData, _ := os.ReadFile(templateDataFilePath)
	templatePath := filepath.Join(
		"_projects",
		projectName,
		"templates",
		wsi.Template,
	)

	m := getTemplateData(templateData)
	m["secrets"] = wsi.Secrets

	renderedTemplate := renderTemplate(
		wsi.Engine,
		templatePath,
		m,
	)

	outputPath := filepath.Join(
		"_projects",
		projectName,
		"output",
		wsi.Output,
	)

	writeOutput(outputPath, renderedTemplate)
}
