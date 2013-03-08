package btcreg

import (
    "net/http"
    "fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Home handler called!")
}

func QueryHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Query handler called!")
}

func JsonQueryHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Json Query handler called!")
}

func NewAddressFormHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("New Address Form Handler called!")
}

func NewAddressHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("New Address Handler called!")
}

func DeleteAddressHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Delete address handler called!")
}
