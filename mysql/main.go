package main

import (
	"database/sql"
	"fmt"

	// 前面的_ 表示导入包但是不使用里面的方法，只会执行 init
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type stu struct {
	id   int
	name string
	age  int
}

func initDb() (err error) {
	dsn := "root:m3479735881@tcp(127.0.0.1:3306)/go_study"
	// Open 时不校验用户名和密码,而是检查dsn的格式，返回值db就是一个数据库的连接池
	// 这里一定不是 :=
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 这里校验
	err = db.Ping()
	if err != nil {
		return
	}
	// 可以设置db这个数据库连接池中连接的最大个数
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数
	// db.SetMaxIdleConns(8)
	return
}

// 查询一行
func query(i int) {
	var stu1 stu
	// 1. 查询
	sqlStr := "select id,name,age from student where id=?;"
	// 从连接池中取出一个连接去查询。
	ret := db.QueryRow(sqlStr, i)
	// 取出结果放到结构体里,同时Scan内部会释放连接！！所以一般必须调用Scan释放连接
	ret.Scan(&stu1.id, &stu1.name, &stu1.age)
	fmt.Printf("%#v\n", stu1)
}

func queryMore(i int) {
	sqlStr := "select id,name,age from student where id > ?;"
	rets, err := db.Query(sqlStr, i)
	if err != nil {
		fmt.Printf("Query error:%#v\n", err)
		return
	}
	// 这个要自己手动关闭！
	defer rets.Close()
	// 循环取值
	var stu1 stu
	for rets.Next() {
		rets.Scan(&stu1.id, &stu1.name, &stu1.age)
		fmt.Printf("%#v\n", stu1)
	}
}

func insert(id int, name string, age int) {
	sqlStr := "insert into student values (?, ?, ?)"
	req, err := db.Exec(sqlStr, id, name, age)
	if err != nil {
		fmt.Printf("Exec error:%#v\n", err)
		return
	}
	// 获取插入数据的ID
	req.LastInsertId()
	// 获取受影响的行数
	req.RowsAffected()
}

func update() {
	sqlStr := "update student set age=1111 where id=1;"
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("Exec error:%#v\n", err)
		return
	}
	// 获取修改数据的ID
	res.LastInsertId()
	// 获取受影响的行数
	res.RowsAffected()
}

func prepareInsert() {
	sqlStr := "insert into student values (?,?,?)"
	// 预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error:%#v\n", err)
		return
	}
	defer stmt.Close()
	// 插入,填入之前的占位符即可
	req, err := stmt.Exec(10, "asdfsadf", 123)
	if err != nil {
		fmt.Printf("prepare error:%#v\n", err)
		return
	}
	req.LastInsertId()
	// req.RowsAffected()
}

func transaction() {
	// 开始一个事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Begin error:%#v\n", err)
		return
	}
	// 执行多条SQL
	sqlStr1 := "update student set age = 1 where id = 1"
	sqlStr2 := "update student set age = 2 where id = 2"

	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 出错了要回滚
		tx.Rollback()
		fmt.Printf("Exec error:%#v\n", err)
		return
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 出错了要回滚
		tx.Rollback()
		fmt.Printf("Exec error:%#v\n", err)
		return
	}
	// 所有的SQL都成功了提交本次事务
	err = tx.Commit()
	if err != nil {
		// 出错了要回滚
		tx.Rollback()
		fmt.Printf("Exec error:%#v\n", err)
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init error:%#v\n", err)
	}
	fmt.Println("连接成功")
	// query(2)
	// queryMore(1)
	// insert(5, "tianqi", 990)
	// update()
	// prepareInsert()
	transaction()
}
