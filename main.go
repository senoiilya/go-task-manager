package main

import "github.com/gin-gonic/gin"

func main() {
	// Инициализация БД и закрытие
	initDB()
	defer db.Close()
	// Создание нового Gin роутера
	r := gin.Default()

}
