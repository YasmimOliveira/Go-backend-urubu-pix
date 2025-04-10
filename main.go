// main.go

package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "backend-urubu-do-pix/handlers"
    "backend-urubu-do-pix/utils"
)

func main() {
    utils.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/deposit", handlers.Deposit).Methods("POST")
    r.HandleFunc("/balance/{id}", handlers.GetBalance).Methods("GET")
    r.HandleFunc("/withdraw/{id}", handlers.Withdraw).Methods("POST")

    log.Println("Servidor rodando na porta 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}