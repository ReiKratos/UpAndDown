package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "sobidesce/lobby"
)

// our main function
func main() {
  router := lobby.NewRouter()

  log.Fatal(http.ListenAndServe(":8000", router))
}
