package handler

import (
	"html/template"
	"log"
	"net/http"
)

var (
	// Templates for handlers
	templates *template.Template
)

// InitHandlers initializes OAuth package
func InitHandlers() {

	// Templates
	tmpls, err := template.ParseGlob("template/*.html")
	if err != nil {
		log.Fatalf("Error while parsing templates: %v", err)
	}
	templates = tmpls
}

func getCurrentUserID(r *http.Request) string {
	c, _ := r.Cookie(userIDCookieName)
	if c != nil {
		return c.Value
	}
	return ""
}
