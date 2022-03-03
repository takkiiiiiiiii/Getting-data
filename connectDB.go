package main

import(
   "database/sql"
   "fmt"
   "database/conf"
   _ "github.com/go-sql-driver/mysql"
)

type Users struct {
   Id        int
   Account   string
   Name      string
   Passwd    string
   Create    string
}

func connectionDB() *sql.DB{
   confDB, err := conf.ReadConfDB()
   if err != nil {
       fmt.Println("Read Failure!!")
   }
   //接続文字列の生成
   conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)
   db, err := sql.Open("mysql", conStr)
   if err != nil{
       fmt.Println("Error!!")
   }
   return db
}

func getData(db *sql.DB) *sql.Rows{
   data, err := db.Query("SELECT * FROM M_User")
   if err != nil{
      fmt.Println("Error!!")
      panic(err.Error)
   }
   return data
}

func main() {
    db := connectionDB()
    defer db.Close()
    data := getData(db)
        // 行データ取得
    users := Users{}
    for data.Next() {
        error := data.Scan(&users.Id, &users.Account, &users.Name, &users.Passwd, &users.Create)
        if error != nil {
            fmt.Println("scan error")
        } else {
              fmt.Println(users)
        }
    }
}
