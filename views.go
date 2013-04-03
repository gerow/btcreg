package btcreg

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "html/template"
    "bytes"
    "strings"
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

type FailedQueryData struct {
    Email string
}

type AddData struct {
    Error bool
    DefaultValue string
}

type Delete struct {
    Error bool
    DefaultValue string
}

type AddSuccessData struct {
    Email string
}

type DeleteSuccessData struct {
    Email string
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
    var query_content bytes.Buffer
    t, err := template.ParseFiles("templates/successful_query.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&query_content, QueryData{email, address})
    if err != nil {
      return nil, err
    }
    nav, err := buildNav("Query")
    if err != nil {
      return nil, err
    }
    data, err := buildBase(email + "|BtcReg", template.HTML(nav), template.HTML(query_content.Bytes()))
    if err != nil {
      return nil, err
    }

    return data, nil
}

func buildFailedQuery(email string) ([]byte, error) {
    var query_content bytes.Buffer
    t, err := template.ParseFiles("templates/failed_query.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&query_content, FailedQueryData{email})
    if err != nil {
      return nil, err
    }
    nav, err := buildNav("Query")
    if err != nil {
      return nil, err
    }
    data, err := buildBase(email + "|BtcReg", template.HTML(nav), template.HTML(query_content.Bytes()))
    if err != nil {
      return nil, err
    }

    return data, nil
}

func buildAdd() ([]byte, error) {
  return _buildAdd(false, "")
}

func buildAddWithError(defaultValue string) ([]byte, error) {
  return _buildAdd(true, defaultValue)
}

func _buildAdd(yesError bool, defaultValue string) ([]byte, error) {
  var add_data bytes.Buffer
    t, err := template.ParseFiles("templates/add.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&add_data, AddData{yesError, defaultValue})
    if err != nil {
      return nil, err
    }
    nav, err := buildNav("Add")
    if err != nil {
      return nil, err
    }
    data, err := buildBase("Add|BtcReg", template.HTML(nav), template.HTML(add_data.Bytes()))
    if err != nil {
      return nil, err
    }

    return data, nil
}

func buildAddSuccess(email string) ([]byte, error) {
  var addSuccessData bytes.Buffer
  t, err := template.ParseFiles("templates/add_success.html")
  if err != nil {
    return nil, err
  }

  err = t.Execute(&addSuccessData, AddSuccessData{email})
  if err != nil {
    return nil, err
  }
  nav, err := buildNav("Add")
  if err != nil {
    return nil, err
  }
  data, err := buildBase("Add|BtcReg", template.HTML(nav), template.HTML(addSuccessData.Bytes()))
  if err != nil {
    return nil, err
  }

  return data, nil
}

func buildDelete() ([]byte, error) {
  return _buildDelete(false, "")
}

func buildDeleteWithError(defaultValue string) ([]byte, error) {
  return _buildDelete(true, defaultValue)
}

func _buildDelete(yesError bool, defaultValue string) ([]byte, error) {
  var delete_data bytes.Buffer
    t, err := template.ParseFiles("templates/delete.html")
    if err != nil {
      return nil, err
    }

    err = t.Execute(&delete_data, AddData{yesError, defaultValue})
    if err != nil {
      return nil, err
    }
    nav, err := buildNav("Delete")
    if err != nil {
      return nil, err
    }
    data, err := buildBase("Delete|BtcReg", template.HTML(nav), template.HTML(delete_data.Bytes()))
    if err != nil {
      return nil, err
    }

    return data, nil
}

func buildDeleteSuccess(email string) ([]byte, error) {
  var deleteSuccessData bytes.Buffer
  t, err := template.ParseFiles("templates/delete_success.html")
  if err != nil {
    return nil, err
  }

  err = t.Execute(&deleteSuccessData , DeleteSuccessData{email})
  if err != nil {
    return nil, err
  }
  nav, err := buildNav("Delete")
  if err != nil {
    return nil, err
  }
  data, err := buildBase("Delete|BtcReg", template.HTML(nav), template.HTML(deleteSuccessData.Bytes()))
  if err != nil {
    return nil, err
  }

  return data, nil
}


func buildQueryForm() ([]byte, error) {
  var queryFormData bytes.Buffer
  t, err := template.ParseFiles("templates/query.html")
  if err != nil {
    return nil, err
  }

  err = t.Execute(&queryFormData, nil)
  if err != nil {
    return nil, err
  }
  nav, err := buildNav("Query")
  if err != nil {
    return nil, err
  }
  data, err := buildBase("Query|BtcReg", template.HTML(nav), template.HTML(queryFormData.Bytes()))
  return data, nil
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Home handler called!")
    data, err := buildQueryForm()
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    w.Write(data)
}

func QueryHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    fmt.Println("Query handler called!")
    fmt.Println("Got email " + vars["email"])

    addr, err := LoadAddressByEmail(vars["email"])
    if err != nil {
      fmt.Println("failed to find address")
      data, err := buildFailedQuery(vars["email"])
      if err != nil {
        fmt.Println("got error " + err.Error())
        w.WriteHeader(500)
        return
      }
      w.WriteHeader(404)
      w.Write(data)
      return
    }

    data, err := buildSuccessfulQuery(addr.Email, addr.Address)
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

func AddHandler(w http.ResponseWriter, req *http.Request) {
  data, err := buildAdd()
  if err != nil {
    fmt.Println("Got error " + err.Error())
    w.WriteHeader(500)
    return
  }

  w.Write(data)
}

func AddHandlerPost(w http.ResponseWriter, req *http.Request) {
  email := req.FormValue("email")
  fmt.Println("Got email " + email)
  if !strings.Contains(email, "@") {
    // It doesn't look like an email, so send them back
    data, err := buildAddWithError(email)
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    w.Write(data)
  } else {
    data, err := buildAddSuccess(email)
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    w.Write(data)
  }
}

func DeleteAddressHandler(w http.ResponseWriter, req *http.Request) {
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
  data, err := buildDelete()
  if err != nil {
    fmt.Println("Got error " + err.Error())
    w.WriteHeader(500)
    return
  }

  w.Write(data)
}

func DeleteHandlerPost(w http.ResponseWriter, req *http.Request) {
  email := req.FormValue("email")
  fmt.Println("Got email " + email)
  if !strings.Contains(email, "@") {
    // It doesn't look like an email, so send them back
    data, err := buildDeleteWithError(email)
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    w.Write(data)
  } else {
    data, err := buildDeleteSuccess(email)
    if err != nil {
      fmt.Println("Got error " + err.Error())
      w.WriteHeader(500)
      return
    }
    w.Write(data)
  }
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
}
