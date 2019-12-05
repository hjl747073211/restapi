package main


import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "api/model"
    "database/sql"
	 _ "github.com/go-sql-driver/mysql"

)

func main(){

    //连接数据库
    model.DB, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
    defer model.DB.Close()
    r := mux.NewRouter()
    //路由
    r.HandleFunc("/api/getall", model.GetAll).Methods("GET")
    r.HandleFunc("/api/getblog/{id}", model.GetBlog).Methods("GET")
    r.HandleFunc("/api/createblog", model.CreateBlog).Methods("POST")
    r.HandleFunc("/api/updateblog/{id}", model.UpdateBlog).Methods("PUT")
    r.HandleFunc("/api/deleteblog/{id}", model.DeleteBlog).Methods("DELETE")
   
    log.Fatal(http.ListenAndServe(":4200",r))


}



