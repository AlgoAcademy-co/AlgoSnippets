package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/{$}", home) // restrict this route to exact matches on / only
  mux.HandleFunc("/snippet/view/{id}", snippetView)
  mux.HandleFunc("/snipet/create", snippetCreate)

  log.Println("Starting server on :8080")

  err := http.ListenAndServe(":8080", mux)
  log.Fatal(err)
}
