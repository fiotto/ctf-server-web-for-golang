package main

import (
  "os"
  "fmt"
  "strconv"
  "net/http"
  "text/template"
  "database/sql"
  _ "github.com/lib/pq"
)

type Page struct {
  Query string
  Result []User
}

type User struct {
  Id          int
  FirstName   string
  LastName    string
  Job         string
  DeleteFlag  bool
}

const PORT = 8080

func createHtml(w http.ResponseWriter, r *http.Request) error {
  db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=password dbname=ctf_db sslmode=disable")

  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  
  query := r.FormValue("q")
  rows, err := db.Query("SELECT * FROM users WHERE delete_flag=FALSE AND job LIKE '%" + query + "%'")

  if err != nil {
    return err
  }
  defer rows.Close()

  members := []User{}
  for rows.Next() {
    user := User{}
    err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Job, &user.DeleteFlag)
    if err != nil {
      return err
    }
    members = append(members, user)
  }

  page := Page{query, members}
  tmpl, err := template.ParseFiles("index.html") 
  if err != nil {
    return err
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    return err
  }
  
  return nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  err := createHtml(w, r)
  if( err != nil ){
    fmt.Fprint(os.Stderr, err)
    fmt.Fprintf(w, err.Error())
  }
}

func main() {
  http.HandleFunc("/", viewHandler)

  fmt.Println("serving at port", PORT)
  err := http.ListenAndServe(":" + strconv.Itoa(PORT), nil)
  if err != nil {
    fmt.Fprint(os.Stderr, "ListenAndServe failed.\n", err, "\n")
  }
}
