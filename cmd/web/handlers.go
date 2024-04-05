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
  // Initialize a slice containing the paths to the two files.
  // It's important to note that the file containing our base
  // tempalte must be the *first* file in the slice
  files := []string{
    "./ui/html/base.tmpl",
    "./ui/html/partials/nav.tmpl",
    "./ui/html/pages/home.tmpl",
  }

  // Use the template.ParseFiles() function to read the files
  // and store the templates ina  template set. Notice that 
  // we can pass the slice of file paths as a variadic parameter
  ts, err := template.ParseFiles(files...)
  if err != nil{
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }

  // Use the Executetemplate() method to write the content
  // of the base tempatle as the response body
  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
  }
}

func snippetView(w http.ResponseWriter, r *http.Request){
  // Extrac tthe value of the id wildcard from the request using 
  // r.PathValue() and try to convert it to an integer using the 
  // strconv.Atoi() function. If it can"t be converted to an integer
  // or the value is less than 1, we return a 404
  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  // use the fmt.Sprintf() function to interpolate the id value with a 
  // message, the nwrite it as the HTTP response.
  msg := fmt.Sprintf("Display a specifc snippet with ID %d", id)
  w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Create a new snippet"))
}
