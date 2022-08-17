package core

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/osteele/liquid"
)

func getProcessedTemplate(templatePath string) template.Template {
	templateContent, _ := os.ReadFile(templatePath)
	rawTemplate := template.Must(template.New("").Parse(string(templateContent)))
	return *rawTemplate
}

func getRawTemplate(templatePath string) string {
	templateRaw, _ := os.ReadFile(templatePath)
	return string(templateRaw)
}

func renderTemplate(
	engine string,
	templatePath string,
	m map[interface{}]interface{},
) string {
	var renderedTemplate string
	switch engine {
	case "go":
		renderedTemplate = renderGoTemplate(templatePath, m)
	case "liquid":
		renderedTemplate = renderLiquidTemplate(templatePath, m)
	default:
		renderedTemplate = renderGoTemplate(templatePath, m)
	}
	return renderedTemplate
}

func renderGoTemplate(templatePath string, m map[interface{}]interface{}) string {
	t := getProcessedTemplate(templatePath)
	var templateBuffer bytes.Buffer
	if err := t.Execute(&templateBuffer, m); err != nil {
		panic(err)
	}
	result := templateBuffer.String()
	return result
}

func renderLiquidTemplate(
	templatePath string,
	m map[interface{}]interface{},
) string {
	engine := liquid.NewEngine()
	data := map[string]interface{}{
		"data": m["data"],
	}
	t := getRawTemplate(templatePath)
	out, err := engine.ParseAndRenderString(t, data)
	if err != nil {
		log.Fatalln(err)
	}
	return out
}
