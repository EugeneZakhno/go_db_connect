package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	fmt.Println("Работа с Postgres")
	db, err := sql.Open("postgres", "postgresql://godbtest_user:UosIYcNeEmC04atz5u3Mo8VzExRvuA6H@dpg-ct841upu0jms73auig7g-a.oregon-postgres.render.com/godbtest")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to Postgres")

	//Установка данных
	/*
		insert, err := db.Query("INSERT INTO users (name, age) VALUES ('Connection check',01)")
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		fmt.Println("Успешно добавлен пользователь!") */

	// Выборка из Базы данных
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
