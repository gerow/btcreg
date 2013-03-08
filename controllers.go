package btcreg

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Home handler called!")
}

func QueryHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("Query handler called!")
    fmt.Println("Got email " + vars["email"])
}

func JsonQueryHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("Json Query handler called!")
    fmt.Println("Got email " + vars["email"])
}

func NewAddressFormHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("New Address Form Handler called!")
    fmt.Println("Got uuid " + vars["uuid"])
}

func NewAddressHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("New Address Handler called!")
    fmt.Println("Got uuid " + vars["uuid"])
}

func DeleteAddressHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("Delete address handler called!")
    fmt.Println("Got uuid " + vars["uuid"])
}
