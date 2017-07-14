package main

import (
  "encoding/json"
	// "fmt"
  "html/template"
  "net/http"
	"io/ioutil"
	"io"
	"log"
  // "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  log.Println("GET Request made for /view")
  res := getResults("./linux.json")
  t, _ := template.ParseFiles("templates/view.html")
  t.Execute(w, res.Profiles)
}

func generate(w http.ResponseWriter, r *http.Request) {
  log.Println("POST Request made for /generate")
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
      panic(err)
  }
  if err := r.Body.Close(); err != nil {
      panic(err)
  }
  res := Results{}
  json.Unmarshal(body, &res);
  t, _ := template.ParseFiles("templates/view.html") //setp 1
  t.Execute(w, res.Profiles) //step 2
}

func main() {
  // file := os.Args[1]

  server := http.Server{
    Addr: "0.0.0.0:8080",
  }
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/view", handler)
  http.HandleFunc("/generate", generate)
  server.ListenAndServe()
}
