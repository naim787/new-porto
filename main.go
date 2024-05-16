package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func parseTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t := template.Must(template.ParseFiles("views/" + tmpl))
	t.ExecuteTemplate(w, tmpl, data)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parseTemplate(w, "home.html", nil)
}

func projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parseTemplate(w, "projects.html", nil)
}

func adminlogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parseTemplate(w, "login.html", nil)
}


func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	parseTemplate(w, "error.html", nil)
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Println("ini di login")
	parseTemplate(w, "login.html", nil)
}

func main() {
	app := httprouter.New()

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../Public")))

	app.GET("/static/*filepath", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		staticHandler.ServeHTTP(w, r)
	})

	app.GET("/", home)
	app.GET("/Projects", projects)
	app.GET("/Admin", adminlogin)
	app.GET("/Login", login)

	app.NotFound = http.HandlerFunc(handleNotFound)

	fmt.Println("Server berjalan di http://localhost:5000")

	http.ListenAndServe(":5000", app)
}
