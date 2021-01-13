package main
import (
      "fmt"
      //"strings"
      "database/sql"
      _ "github.com/lib/pq"
      "net/http"
      "log"
      "encoding/json"
      "reflect"
      "github.com/gorilla/mux"
      "strconv"
      )

type product struct{
            id int
            name string
            ccal int
}
func DB() []product {
  connStr := "user=admin password=228322 dbname=calority sslmode=disable"
  db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    result, err := db.Query("select * from calority;")

    if err != nil{
       panic(err)
   }
   defer result.Close()
   //fmt.Println("res",result)
    products := []product{}

    for result.Next(){
        p := product{}
        err := result.Scan(&p.id, &p.name, &p.ccal)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }

   /*for _, p := range products{
         fmt.Println(p.id, p.name,p.ccal)
       }
       */
    fmt.Println()
    return products
  }
func homePage(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"welocme to Homepage")
  fmt.Println("Endpoint Hit:homePage")
}
type Article struct {
  Id      int    `json:"Id"`
  Name string  `json:"Name"`
  Ccal int `json:"Ccal"`
}
var Articles []Article
func main(){

  /*
    Articles = []Article{
        Article{Id: "1", name: "Hello", ccal: "228"},
        Article{Id: "2", name: "Hello 2", ccal: "322"},
    }*/
    sl:=DB()
    Articles =make( []Article,len(sl))

    k:=0
    for _,i:=range sl{
      Articles[k].Id = i.id
      Articles[k].Name = i.name
      Articles[k].Ccal = i.ccal

      k++
    }
    //Articles[1]=Article{Id:1,name:"pivo",ccal:1488}
    fmt.Println(reflect.TypeOf(Articles).Kind())
    //fmt.Println(Articles,Articles[0])
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/",homePage)
    myRouter.HandleFunc("/articles",returnAllArticles)
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    fmt.Println("Rest Api Mux")
    log.Fatal(http.ListenAndServe(":9001", myRouter))
  }
func returnSingleArticle(w http.ResponseWriter,r *http.Request){
  fmt.Println("singleart")
  vars:=mux.Vars(r)
  key :=vars["id"]
  //fmt.Fprintf(w,"Key:"+key)
  for _, article := range Articles {
        if strconv.Itoa(article.Id) == key {
          fmt.Println(article,reflect.TypeOf(article).Kind())
            json.NewEncoder(w).Encode(article)
        }
    }

}
func returnAllArticles(w http.ResponseWriter, r *http.Request){
  fmt.Println("art",Articles[0])
  json.NewEncoder(w).Encode(Articles)
}
