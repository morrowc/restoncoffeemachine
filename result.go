package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
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

func handle(w http.ResponseWriter, r *http.Request) {
	defer func(t time.Time) { fmt.Fprintf(w, "<address>Generated page in %s thnx!</address>", time.Since(t)) }(time.Now())
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

func main() {
	// Prepare port data from ENV.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8777"
		log.Printf("Default port: %v used", port)
	}

	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)

	log.Printf("Listening on port: %v ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
