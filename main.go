package main

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  // check if the current request URL path exactly matches "/". If it doesn't,
  // use the http.NotFound() function to send a 404 response to the client. 
  // importantly, we then return from the handler. if we don't return the 
  // handler would keep executing and also write the "Hello from AlgoAcademy"
  // message.
  if r.URL.Path != "/"{
    http.NotFound(w, r)
    return
  }

 w.Write([]byte("Hello from AlgoAcademy"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("display the specific snippet"))
}

func snippetCreate( w http.ResponseWriter, r *http.Request){
  w.Write([]byte("lets create a snippet"))
}

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("snippet/create", snippetCreate)

  log.Println("Starting server on port 8080")
  err := http.ListenAndServe(":8080", mux)

  log.Fatal(err)
}
