package templates

import "os"

func getRawTemplate(templatePath string) string {
	templateRaw, _ := os.ReadFile(templatePath)
	return string(templateRaw)
}
