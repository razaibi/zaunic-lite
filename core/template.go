package core

import (
	"bytes"
	"log"
	"text/template"

	"github.com/osteele/liquid"
)

func renderTemplate(
	engine string,
	t template.Template,
	m map[interface{}]interface{},
) string {
	var renderedTemplate string
	switch engine {
	case "go":
		renderedTemplate = renderGoTemplate(t, m)
	case "liquid":
		renderedTemplate = renderLiquidTemplate(t, m)
	default:
		renderedTemplate = renderGoTemplate(t, m)
	}
	return renderedTemplate
}

func renderGoTemplate(t template.Template, m map[interface{}]interface{}) string {
	var templateBuffer bytes.Buffer
	if err := t.Execute(&templateBuffer, m); err != nil {
		panic(err)
	}
	result := templateBuffer.String()
	return result
}

func renderLiquidTemplate(
	t template.Template,
	m map[interface{}]interface{},
) string {
	engine := liquid.NewEngine()
	template := `<h1>{{ page.title }}</h1>`
	data := map[string]interface{}{
		"data": m["data"],
	}
	out, err := engine.ParseAndRenderString(template, data)
	if err != nil {
		log.Fatalln(err)
	}
	return out
}
