package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация БД и закрытие
	initDB()
	defer db.Close()
	// Создание нового Gin роутера
	r := gin.Default()

	// Роуты API
	r.Get("tasks/", getAllTasksHandle)
	r.GET("/tasks/:id", getTaskByIDHandle)
	r.POST("/tasks", createTaskHandle)
	r.PUT("/tasks/:id", updateTaskHandle)
	r.DELETE("/tasks/:id", deleteTaskHandle)

	// Запуск сервера
	r.Run(":8080")
}

// Получение всех задач
func getAllTasksHandle(c *gin.Context) {
	tasks, err := getTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Получение задачи по ID
func getTaskByIDHandle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный ID"})
		return
	}

	task, err := getTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Создание новой задачи
func createTaskHandle(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавление задачи в базу данных
	id, err := addTask(newTask.Title, newTask.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = int(id)
	c.JSON(http.StatusCreated, newTask)
}

// Обновление существующей задачи
func updateTaskHandle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный ID"})
		return
	}

	var updTask Task
	if err := c.ShouldBindJSON(&updTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = updateTask(id, updTask.Title, updTask.Description, updTask.Done)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	updTask.ID = id
	c.JSON(http.StatusOK, updTask)
}

// Удаление задачи
func deleteTaskHandle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный ID"})
		return
	}

	err = deleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
