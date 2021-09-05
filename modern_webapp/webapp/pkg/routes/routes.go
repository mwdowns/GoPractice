package routes

import (
	"fmt"
	"modern_webapp/webapp/pkg/render"
	"net/http"
)

// Home route
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the home page")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About route
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is the about page")
	render.RenderTemplate(w, "about.page.tmpl")
}
