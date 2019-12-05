package model

import (
	"database/sql"
	 _ "github.com/go-sql-driver/mysql"
	 "log"
	 "fmt"
	 "encoding/json"
	 "net/http"
	 // "math/rand"
	 "strconv"
	 "github.com/gorilla/mux"
)

type blog struct {
	Id int `json:"id"`
	Img string `json:"img"`
	Title string `json:"title"`
	Remark string `json:"remark"`
	Content string `json:"content"`
	Creater string `json:"creater"`
	Category string `json:"category"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	LookNum string `json:"look_num"`
	TalkNum string `json:"talk_num"`
	LikeNum string `json:"like_num"`
}
var DB *sql.DB

//获取所有数据
func GetAll(w http.ResponseWriter, r *http.Request){
	rows, err := DB.Query("SELECT * FROM blog")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)

	var blogs []blog
	
    for rows.Next(){
		var s blog
		err=rows.Scan(&s.Id,&s.Img,&s.Title,&s.Remark,&s.Content,&s.Creater,&s.Category,&s.CreateTime,&s.UpdateTime,&s.LookNum,&s.TalkNum,&s.LikeNum)
        if err != nil {
            log.Fatal(err)
        }
		blogs=append(blogs,s)
		fmt.Println(s)

	}
	rows.Close()
	// fmt.Println(blogs)

	w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(blogs)
	
}


//获取单条数据
func GetBlog(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) //Get params 获取id
	// fmt.Println(params)

	id:=params["id"]

	var s blog
	err := DB.QueryRow("SELECT * FROM blog where id=?",id).Scan(&s.Id,&s.Img,&s.Title,&s.Remark,&s.Content,&s.Creater,&s.Category,&s.CreateTime,&s.UpdateTime,&s.LookNum,&s.TalkNum,&s.LikeNum)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(s)
}


//新建
func CreateBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	r.ParseForm()
	b:=r.Form

	//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
	stmt,err := DB.Prepare("insert into `blog`(img,title,remark)values(?,?,?)")
	if err!=nil{
		fmt.Println("预处理失败:",err)
		return         
	}
	//使用Stmt对象执行预处理参数
	result,err := stmt.Exec(b["img"][0],b["title"][0],b["remark"][0])
	var s blog
	if err!=nil{
		fmt.Println("执行预处理失败:",err)
		return         
	}else{
		//获取新增id
		id,_ := result.LastInsertId()
		//添加成功，返回当前信息
		_ = DB.QueryRow("SELECT * FROM blog where id=?",id).Scan(&s.Id,&s.Img,&s.Title,&s.Remark,&s.Content,&s.Creater,&s.Category,&s.CreateTime,&s.UpdateTime,&s.LookNum,&s.TalkNum,&s.LikeNum)
		json.NewEncoder(w).Encode(s)

	}
}

//更新
func UpdateBlog(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
	id:=params["id"]
	r.ParseForm()
	b:=r.Form
// fmt.Println(b)
// return
	//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
	stmt,err := DB.Prepare("update `blog` set img=?,title=?,remark=? where id=?")
	if err!=nil{
		fmt.Println("预处理失败:",err)
		return         
	}
	//使用Stmt对象执行预处理参数
	result,err := stmt.Exec(b["img"][0],b["title"][0],b["remark"][0],id)
	var s blog
	if err!=nil{
		fmt.Println("执行预处理失败:",err)
		return         
	}else{
		num,err := result.RowsAffected()
		if num!=1{
			fmt.Println("更新失败:",err)
			return    
		}
		//更新成功，返回当前信息
		_ = DB.QueryRow("SELECT * FROM blog where id=?",id).Scan(&s.Id,&s.Img,&s.Title,&s.Remark,&s.Content,&s.Creater,&s.Category,&s.CreateTime,&s.UpdateTime,&s.LookNum,&s.TalkNum,&s.LikeNum)
		json.NewEncoder(w).Encode(s)

	}

}


//删除
func DeleteBlog(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
	id:=params["id"]
	
	stmt, _ := DB.Prepare(`DELETE FROM blog WHERE id=?`)
	res, _ := stmt.Exec(id)
	num, _ := res.RowsAffected()
	if num!=1{
		data:=[]string{strconv.Itoa(http.StatusNotFound), "No Blog Found"}
		json.NewEncoder(w).Encode(data)
	}else{
		data:=[]string{strconv.Itoa(http.StatusOK), "OK"}
		json.NewEncoder(w).Encode(data)
	}


}





