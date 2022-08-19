package templates

import (
	"log"

	"github.com/osteele/liquid"
)

func RenderLiquidTemplate(
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
