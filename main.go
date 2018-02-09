package main

import (
  "github.com/joho/godotenv"
  "log"
  "net/http"
)

func main() {
  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  log.Print("Started server")

  router := CreateRouter()

  log.Fatal(http.ListenAndServe(":8000", router))
}
