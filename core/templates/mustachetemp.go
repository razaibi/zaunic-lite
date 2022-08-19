package templates

import (
	"log"

	"github.com/cbroglie/mustache"
)

func RenderMustacheTemplate(
	templatePath string,
	m map[interface{}]interface{},
) string {
	t := getRawTemplate(templatePath)
	out, err := mustache.Render(t, map[string]interface{}{
		"data": m["data"],
	})
	if err != nil {
		log.Fatalln(err)
	}
	return out
}
