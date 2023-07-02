package renderers

import (
	"bytes"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package.
func NewTemplate(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		// get the template cache from AppConfig
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatalln("cannot create template cache:", err)
		}
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("cannot get template from template cache")
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all files named *.page.tmpl from ./templates/
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
