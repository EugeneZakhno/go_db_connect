package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	fmt.Println("Работа с MySql")
	db, err := sql.Open("mysql", "root:Decent@1Ljcnjqysq@tcp(34.56.246.79:3306)/forGolangPrj")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to MySQL")

	//Установка данных
	/*
		insert, err := db.Query("INSERT INTO users (name, age) VALUES ('Diana',19)")
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		fmt.Println("Успешно добавлен пользователь!")
	*/

	res, err := db.Query("select name, age from users")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Print(fmt.Sprintf("User: %s with age: %d", user.Name, user.Age))
	}
}
