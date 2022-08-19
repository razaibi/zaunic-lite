package templates

import (
	"log"

	"github.com/aymerick/raymond"
)

func RenderHandlebarsTemplate(
	templatePath string,
	m map[interface{}]interface{},
) string {
	t := getRawTemplate(templatePath)
	data := map[string]interface{}{
		"data": m["data"],
	}

	out, err := raymond.Render(t, data)
	if err != nil {
		log.Fatalln(err)
	}

	return out
}
