package core

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"zaunic-lite/models"

	"gopkg.in/yaml.v2"
)

func getTemplate(templatePath string) template.Template {
	templateContent, _ := os.ReadFile(templatePath)
	rawTemplate := template.Must(template.New("").Parse(string(templateContent)))
	return *rawTemplate
}

func getTemplateData(templateData []byte) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(templateData), &m); err != nil {
		panic(err)
	}
	return m
}

func renderTemplate(t template.Template, m map[interface{}]interface{}) string {
	var templateBuffer bytes.Buffer
	if err := t.Execute(&templateBuffer, m); err != nil {
		panic(err)
	}
	result := templateBuffer.String()
	return result
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
	wsi models.Item,
) {
	templateDataFilePath := filepath.Join(
		"data",
		wsi.Data,
	)

	templateData, _ := os.ReadFile(templateDataFilePath)
	templateFilePath := filepath.Join(
		"templates",
		wsi.Template,
	)
	t := getTemplate(templateFilePath)
	m := getTemplateData(templateData)
	renderedTemplate := renderTemplate(t, m)

	outputPath := filepath.Join(
		"output",
		wsi.Output,
	)

	writeOutput(outputPath, renderedTemplate)
}
