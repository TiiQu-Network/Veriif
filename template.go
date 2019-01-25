package main

func GetTemplates(layout string, name string) (string, string) {
	return layoutDir + layout + ".html", templateDir + name + ".html"
}

func GetTemplate(name string) string {
	return templateDir + name + ".html"
}