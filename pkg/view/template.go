package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	tpIndex       = parseTemplate("root.tmpl", "index.tmpl")
	tpNews        = parseTemplate("root.tmpl", "news.tmpl")
	tpAdminLogin  = parseTemplate("root.tmpl", "admin/login.tmpl")
	tpAdminList   = parseTemplate("root.tmpl", "admin/list.tmpl")
	tpAdminCreate = parseTemplate("root.tmpl", "admin/create.tmpl")
	tpAdminEdit   = parseTemplate("root.tmpl", "admin/edit.tmpl")
)

const templateDir = "template"

func joinTemplateDir(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}
func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
