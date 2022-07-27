package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//创建user结构体
type user struct {
	Loginid  string `db:"loginid"`
	Loginpwd string `db:"loginpwd"`
	Lickname string `db:"nickname"`
	Tel      string `db:"tel"`
	Pic      string `db:"pic"'`
}

//结构体（struct）里面的成员变量首字母是区分大小写的，若首字母大写则该成员为公有成员(对外可见)，否则是私有成员(对外不可见)。
var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:wl5201314@tcp(127.0.0.1:3306)/mytest")
	if err != nil {
		fmt.Println("mysql open failed", err)
		return
	}
	Db = database
	//defer Db.Close()不知道为啥数据库直接被关了
}
func main() {
	//执行sql语句
	/*	r, err := Db.Exec("insert into user(loginid,loginpwd,nickname,tel,pic)values(?,?,?,?,?)", "131", "110", "111", "112", "113")
		if err != nil {
			fmt.Println("exec failed", err)
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Println("insert failed", err)
			return
		}
		fmt.Println("insert success:", id)*/

	//查询功能
	/*	var User []user
		err := Db.Select(&User, "select * from user where loginid=?", 123)
		if err != nil {
			fmt.Println("exec failed", err)
			return
		}
		fmt.Println("select success:", User)*/

	//更新操作
	/*	res, err := Db.Exec("update user set nickname=? where loginId=?", "test", "131")
		if err != nil {
			fmt.Println("exec failed:", err)
			return
		}
		row, err := res.RowsAffected()
		if err != nil {
			fmt.Println("row failed", err)

		}
		fmt.Println("update success", row)*/

	//删除操作
	res, err := Db.Exec("delete from user where loginid=?", 130)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete success: ", row)
}
