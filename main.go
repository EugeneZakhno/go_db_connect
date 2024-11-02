package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	fmt.Println("Работа с MySql")
	db, err := sql.Open("mysql", "root:Decent@1Ljcnjqysq@tcp(34.56.246.79:3306)/forGolangPrj")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to MySQL")
	//Установка данных
	insert, err := db.Query("INSERT INTO users (name, age) VALUES ('Diana',19)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	fmt.Println("Успешно добавлен пользователь!")

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
