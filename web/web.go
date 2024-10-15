// Author: pat@patfairbank.com (Patrick Fairbank)
// Modified for Isotope Robotics by: Ethen Brandenburg

package web

// Handles all Web Server Information

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/Isotope-Robotics/Isotope-Scouting-System/model"
)

const (
	sessionTokenCookie = "session_token"
	adminUser          = "admin"
)

type Web struct {
	templateHelpers template.FuncMap
}

func NewWeb() *Web {
	web := &Web{}

	// Helper functions that can be used inside templates.
	web.templateHelpers = template.FuncMap{
		// Allows sub-templates to be invoked with multiple arguments.
		"dict": func(values ...any) (map[string]any, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("invalid dict call")
			}
			dict := make(map[string]any, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"add": func(a, b int) int {
			return a + b
		},
		"itoa": func(a int) string {
			return strconv.Itoa(a)
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"seq": func(count int) []int {
			seq := make([]int, count)
			for i := 0; i < count; i++ {
				seq[i] = i + 1
			}
			return seq
		},
		"toUpper": func(str string) string {
			return strings.ToUpper(str)
		},
	}

	return web
}

// Starts the webserver and blocks, waiting on requests. Does not return until the application exits.
func (web *Web) ServeWebInterface(port int) {
	http.Handle("/static/", http.StripPrefix("/static/", addNoCacheHeader(http.FileServer(http.Dir("static/")))))
	http.Handle("/", web.newHandler())
	log.Printf("Serving HTTP requests on port %d", port)

	// Start Server
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// Serves the root page of Isotope Scouting System.
func (web *Web) indexHandler(w http.ResponseWriter, r *http.Request) {
	template, err := web.parseFiles("templates/index.html", "templates/base.html")
	if err != nil {
		handleWebErr(w, err)
		return
	}
	data := struct {
		*model.EventSettings
	}{web.arena.EventSettings}
	err = template.ExecuteTemplate(w, "base", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}
}

// Adds a "Cache-Control: no-cache" header to the given handler to force browser validation of last modified time.
func addNoCacheHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache")
		handler.ServeHTTP(w, r)
	})
}

func (web *Web) newHandler() {
	mux := http.NewServeMux()
	return mux
}

// Writes the given error out as plain text with a status code of 500.
func handleWebErr(w http.ResponseWriter, err error) {
	log.Printf("HTTP request error: %v", err)
	http.Error(w, "Internal server error: "+err.Error(), 500)
}

// Prepends the base directory to the template filenames.
func (web *Web) parseFiles(filenames ...string) (*template.Template, error) {
	var paths []string
	for _, filename := range filenames {
		paths = append(paths, filepath.Join(model.BaseDir, filename))
	}

	template := template.New("").Funcs(web.templateHelpers)
	return template.ParseFiles(paths...)
}
