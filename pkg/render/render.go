package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var (
	// TemplateCache to store parsed templates
	tc = make(map[string]*template.Template)
	// Mutex for safe concurrent access
	mu sync.RWMutex
)

// RenderTemplate checks the cache for the template and renders it
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Acquire read lock to check if the template is in the cache
	mu.RLock()
	_, inMap := tc[t]
	mu.RUnlock()

	if !inMap {
		// Template not in cache, so create it
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error creating template", http.StatusInternalServerError)
			return
		}
	} else {
		// Template found in cache
		log.Println("Using cached template")
	}

	// Acquire read lock again to safely access the template in cache
	mu.RLock()
	tmpl = tc[t]
	mu.RUnlock()

	// Render the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// createTemplateCache parses and adds the template to the cache
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the template files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Acquire write lock to safely update the cache
	mu.Lock()
	tc[t] = tmpl
	mu.Unlock()

	return nil
}
