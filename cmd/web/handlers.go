package main

import (
  "fmt"
  "net/http"
  "strconv"
  "log"
  "html/template"
)

func home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" {
    http.NotFound(w,r)
    return
  }
  // Use the template.ParseFiles() function to read the template file into a 
  // template set. If there's an error, we log the detailed error message
  // and use the http.Errr() function to send a generic 500 internal server
  // error response to the user
  ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }

  // We then use the Execute() method on the template set to write the 
  // tempalte content as the response body. The last parameter to
  // Execute() represents any dynamic data that we want to pass in, which
  // for now we leave as nil.
  err = ts.Execute(w,nil) 
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
  }
}

func snippetView(w http.ResponseWriter, r *http.Request){
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  fmt.Fprintf(w, "Display a specific snippet with ID %d..", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Create a new snippet"))
}
