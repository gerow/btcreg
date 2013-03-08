package btcreg

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "html/template"
    "bytes"
)

type BaseData struct {
    Title string
    Nav template.HTML
    Content template.HTML
}

type NavData struct {
    QueryActive bool
    AddActive bool
    DeleteActive bool
    AboutActive bool
}

type QueryData struct {
    Email string
    Address string
}

func buildBase(title string, nav template.HTML, content template.HTML) ([]byte, error) {
    var buf bytes.Buffer
    t, err := template.ParseFiles("templates/base.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&buf, BaseData{title, nav, content})
    if err != nil {
      return nil, err
    }
    return buf.Bytes(), nil
}

func buildNav(active string) ([]byte, error) {
    var nd NavData
    if active == "Query" {
      nd.QueryActive = true
    } else if active == "Add" {
      nd.AddActive = true
    } else if active == "Delete" {
      nd.DeleteActive = true
    } else {
      nd.AboutActive = true
    }

    var buf bytes.Buffer
    t, err := template.ParseFiles("templates/nav.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&buf, nd)
    if err != nil {
      return nil, err
    }
    return buf.Bytes(), nil
}

func buildSuccessfulQuery(email string, address string) ([]byte, error) {
    var buf bytes.Buffer
    t, err := template.ParseFiles("templates/successful_query.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&buf, QueryData{email, address})
    if err != nil {
      return nil, err
    }
    return buf.Bytes(), nil
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Home handler called!")
}

func QueryHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("Query handler called!")
    fmt.Println("Got email " + vars["email"])

    nav, err := buildNav("Query")
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    content, err := buildSuccessfulQuery("gerow.mike@gmail.com", "thisisatest")
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    title := "query for gerow.mike@gmail.com"

    data, err := buildBase(title, template.HTML(nav), template.HTML(content))
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }

    w.Write(data)
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

func AddHandler(w http.ResponseWriter, req *http.Request) {
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
}
