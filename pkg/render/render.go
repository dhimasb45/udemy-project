package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

// package render

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"path/filepath"
// )

// var templateCache = map[string]*template.Template{}

// // LoadTemplates akan mem-parsing semua template dan menyimpannya dalam cache
// func LoadTemplates() {
// 	pages, err := filepath.Glob("./templates/*.html")
// 	if err != nil {
// 		fmt.Println("error finding templates:", err)
// 		return
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		tmpl, err := template.ParseFiles(page)
// 		if err != nil {
// 			fmt.Println("error parsing template:", err)
// 			continue // Lanjutkan ke template berikutnya jika parsing gagal
// 		}
// 		templateCache[name] = tmpl
// 	}
// }

// // RenderTemplate akan menggunakan template yang sudah di-cache untuk dirender
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	t, ok := templateCache[tmpl]
// 	if !ok {
// 		http.Error(w, "Template not found", http.StatusInternalServerError)
// 		fmt.Println("template not found:", tmpl)
// 		return
// 	}

// 	err := t.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error executing template:", err)
// 		http.Error(w, "Error rendering template", http.StatusInternalServerError)
// 		return
// 	}
// }
