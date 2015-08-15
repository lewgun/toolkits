package gists

http://studygolang.com/articles/1741

func unescaped (x string) interface{} { return template.HTML(x) }

func renderTemplate(w http.ResponseWriter, tmpl string, view *Page) {
	t := template.New("")
	t = t.Funcs(template.FuncMap{"unescaped": unescaped})
	t, err := t.ParseFiles("view.html", "edit.html")
	err = t.ExecuteTemplate(w, tmpl + ".html", view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


{{printf "%s" .Body | unescaped}} //[]byte
{{.Body | unescaped}} //string