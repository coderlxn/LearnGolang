package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main()  {
	db, err := sql.Open("mysql", "joywok:joywok2048@(127.0.0.1:3306)/jwassistant") //NOTE 注意数据库的地址格式
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	stmtIns, err := db.Prepare("INSERT INTO godemo VALUES (?, ?)")  // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	stmtOut, err := db.Prepare("SELECT squareNumber from godemo WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	//for i := 0; i < 25; i++ {
	//	_, err = stmtIns.Exec(i, i * i)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//}

	var squareNum int  // 存储查询结果
	err = stmtOut.QueryRow(13).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)

	err = stmtOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}
