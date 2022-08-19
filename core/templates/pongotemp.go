package templates

import (
	"github.com/flosch/pongo2"
)

func RenderPongoTemplate(
	templatePath string,
	m map[interface{}]interface{},
) string {
	t := getRawTemplate(templatePath)
	tpl, err := pongo2.FromString(t)
	if err != nil {
		panic(err)
	}
	out, err := tpl.Execute(pongo2.Context{
		"data": m["data"],
	})
	if err != nil {
		panic(err)
	}
	return out
}
