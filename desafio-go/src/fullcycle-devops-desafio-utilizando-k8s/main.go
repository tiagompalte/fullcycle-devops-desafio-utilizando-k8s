package main

import (
    "net/http"
)

func greeting(message string) string {
  return "<b>" + message + "</b>"
}

func index(w http.ResponseWriter, req *http.Request) {
  message := greeting("Code.education Rocks!")

  w.WriteHeader(200)
  w.Write([]byte(message))
}

func main() {

    http.HandleFunc("/", index)

    http.ListenAndServe(":8000", nil)
}
