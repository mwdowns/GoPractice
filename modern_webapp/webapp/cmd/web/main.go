package main

import (
	"fmt"
	"log"
	"modern_webapp/webapp/pkg/config"
	"modern_webapp/webapp/pkg/constants"
	"modern_webapp/webapp/pkg/render"
	"modern_webapp/webapp/pkg/routes"
	"net/http"
)

// main serves the app
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/about", routes.About)

	fmt.Println("listing on port", constants.PortNumber)
	http.ListenAndServe(constants.PortNumber, nil)
}
