package main

import (
	"fmt"
	"log"
	"modern_webapp/webapp/pkg/config"
	"modern_webapp/webapp/pkg/constants"
	"modern_webapp/webapp/pkg/handlers"
	"modern_webapp/webapp/pkg/render"
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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("listing on port", constants.PortNumber)
	http.ListenAndServe(constants.PortNumber, nil)
}
