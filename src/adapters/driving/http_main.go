package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"hexhello/src/adapters/driven"
	"hexhello/src/app"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "World"
	}
	presenter := driven.NewHTTPPresenter(w)
	svc := app.NewGreetUserService(presenter)
	svc.Greet(name)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, `<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Hello Hexagon</title>
    <style>
      body { font-family: -apple-system, system-ui, Segoe UI, Roboto, Ubuntu, Cantarell, Noto Sans, sans-serif; margin: 2rem; }
      form { display: flex; gap: .5rem; }
      input { padding: .5rem .75rem; font-size: 1rem; }
      button { padding: .5rem .75rem; font-size: 1rem; }
    </style>
  </head>
  <body>
    <h1>What's your name?</h1>
    <form action="/greet" method="POST">
      <input type="text" name="name" placeholder="Your name" />
      <button type="submit">Greet</button>
    </form>
    <p>Or try: <code>/greet?name=Alice</code></p>
  </body>
</html>`)
}

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/greet", greetHandler)
	log.Printf("HTTP server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
