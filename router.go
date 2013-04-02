package btcreg

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

func RunRouter() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/q/{email}", QueryHandler).Methods("GET")
    r.HandleFunc("/j/{email}", JsonQueryHandler).Methods("GET")
    r.HandleFunc("/n/{uuid}", NewAddressFormHandler).Methods("GET")
    r.HandleFunc("/n/{uuid}", NewAddressHandler).Methods("POST")
    r.HandleFunc("/d/{uuid}", DeleteAddressHandler).Methods("DELETE")
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))

    // Form methods
    r.HandleFunc("/add/", AddHandler).Methods("GET")
    r.HandleFunc("/add/", AddHandlerPost).Methods("POST")
    r.HandleFunc("/delete/", DeleteHandler).Methods("GET")
    r.HandleFunc("/delete/", DeleteHandlerPost).Methods("POST")
    r.HandleFunc("/about/", AboutHandler).Methods("GET")

    fmt.Println("router initialized")
    fmt.Println("listening on port 8080...")
    http.ListenAndServe(":8080", r)
}
