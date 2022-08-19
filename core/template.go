package core

import (
	"zaunic-lite/core/templates"
)

func renderTemplate(
	engine string,
	templatePath string,
	m map[interface{}]interface{},
) string {
	var renderedTemplate string
	switch engine {
	case "go":
		renderedTemplate = templates.RenderGoTemplate(templatePath, m)
	case "liquid":
		renderedTemplate = templates.RenderLiquidTemplate(templatePath, m)
	case "pongo":
		renderedTemplate = templates.RenderPongoTemplate(templatePath, m)
	case "mustache":
		renderedTemplate = templates.RenderMustacheTemplate(templatePath, m)
	case "handlebars":
		renderedTemplate = templates.RenderHandlebarsTemplate(templatePath, m)
	default:
		renderedTemplate = templates.RenderGoTemplate(templatePath, m)
	}
	return renderedTemplate
}
