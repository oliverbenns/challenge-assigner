package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
  router := mux.NewRouter()

  // Views - @TODO: Fix weird trailing slash and multiple file servers.
  admin := http.FileServer(http.Dir("./views/admin"))
  start := http.FileServer(http.Dir("./views/start"))
  router.Handle("/admin/", http.StripPrefix("/admin/", admin))
  router.Handle("/start/", http.StripPrefix("/start/", start))

  // Api
  router.HandleFunc("/api/challenge/create", CreateChallenge).Methods("POST")
  router.HandleFunc("/api/challenge/start", StartChallenge).Methods("POST")

  return router
}
