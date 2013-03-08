package btcreg

import (
    "net/http"
    "github.com/gorilla/mux"
)

func Main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/q/{email}", QueryHandler).Methods("GET")
    r.HandleFunc("/j/{email}", JsonQueryHandler).Methods("GET")
    r.HandleFunc("/n/{uuid}", NewAddressFormHandler).Methods("GET")
    r.HandleFunc("/n/{uuid}", NewAddressHandler).Methods("POST")
    r.HandleFunc("/d/{uuid}", DeleteAddressHandler).Methods("DELETE")
    http.ListenAndServe(":8080", r)
}
