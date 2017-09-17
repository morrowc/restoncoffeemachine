package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"google.golang.org/appengine"
)

type results []struct {
	Color       string
	Schema      string
	Schema2     string
	Exclamation string
	Broken      string
	Image       string
}

var (
	t    = template.Must(template.New("template").Parse(mainPage))
	data = results{{
		Color:       "red",
		Schema:      "https",
		Schema2:     "https",
		Exclamation: "NO",
		Broken:      "BROKEN",
		Image:       "no-coffee.jpg",
	}, {
		Color:       "green",
		Schema:      "https",
		Schema2:     "https",
		Exclamation: "YES",
		Broken:      "ON",
		Image:       "coffee-works.jpg",
	}}

	randPool = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	// Start the appengine main.
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)

	appengine.Main()
	log.Print("Listening on port 8777")
	log.Fatal(http.ListenAndServe(":8777", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// var result bytes.Buffer
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := t.Execute(w, data[randPool.Intn(len(data))])
	if err != nil {
		log.Fatal("executing the template failed: %v", err)
	}
	// fmt.Fprintf(w, string(result))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
