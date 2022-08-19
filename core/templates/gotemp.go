package templates

import (
	"bytes"
	"os"
	"text/template"
)

func getProcessedTemplate(templatePath string) template.Template {
	templateContent, _ := os.ReadFile(templatePath)
	rawTemplate := template.Must(template.New("").Parse(string(templateContent)))
	return *rawTemplate
}

func RenderGoTemplate(templatePath string, m map[interface{}]interface{}) string {
	t := getProcessedTemplate(templatePath)
	var templateBuffer bytes.Buffer
	if err := t.Execute(&templateBuffer, m); err != nil {
		panic(err)
	}
	result := templateBuffer.String()
	return result
}
