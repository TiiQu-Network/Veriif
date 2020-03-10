package templates

import "github.com/TiiQu-Network/Veriif/app/paths"

func GetTemplates(layout string, name string) (string, string) {
	return paths.LayoutDir + layout + ".html", paths.TemplateDir + name + ".html"
}

func GetTemplate(name string) string {
	return paths.TemplateDir + name + ".html"
}