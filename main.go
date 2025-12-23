package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates map[string]*template.Template

func loadTemplates() {
	templates = make(map[string]*template.Template)
	pages := []string{
		"home.html", "about.html", "academics.html", "admissions.html",
		"research.html", "student-life.html", "news.html", "directory.html", "contact.html",
	}

	for _, page := range pages {
		tmpl, err := template.ParseFiles(filepath.Join("templates", "base.html"), filepath.Join("templates", page))
		if err != nil {
			log.Fatalf("Error parsing template %s: %v", page, err)
		}
		templates[page] = tmpl
	}
}

func render(w http.ResponseWriter, page string, path string) {
	tmpl, ok := templates[page]
	if !ok {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	data := map[string]string{
		"Path": path,
	}

	// Execute the base template (which executes "content" block)
	err := tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Error executing template %s: %v", page, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		render(w, "home.html", "/")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about.html", "/about")
	})

	mux.HandleFunc("/academics", func(w http.ResponseWriter, r *http.Request) {
		render(w, "academics.html", "/academics")
	})

	mux.HandleFunc("/admissions", func(w http.ResponseWriter, r *http.Request) {
		render(w, "admissions.html", "/admissions")
	})

	mux.HandleFunc("/research", func(w http.ResponseWriter, r *http.Request) {
		render(w, "research.html", "/research")
	})

	mux.HandleFunc("/student-life", func(w http.ResponseWriter, r *http.Request) {
		render(w, "student-life.html", "/student-life")
	})

	mux.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		render(w, "news.html", "/news")
	})

	mux.HandleFunc("/directory", func(w http.ResponseWriter, r *http.Request) {
		render(w, "directory.html", "/directory")
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		render(w, "contact.html", "/contact")
	})

	return mux
}

func main() {
	loadTemplates()
	mux := setupRouter()

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
