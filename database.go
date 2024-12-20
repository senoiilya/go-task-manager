package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/senoiilya/go-task-manager/queries"
)

var db *sql.DB

// Инициализация базы данных
func initDB() {
	var err error

	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(queries.InitTasks)
	if err != nil {
		log.Fatal(err)
	}
}

// Функция добавления новой задачи в бд
func addTask(title, description string) (int64, error) {
	result, err := db.Exec(queries.InsertTask, title, description, false)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getTasks() ([]Task, error) {
	rows, err := db.Query(queries.GetAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func getTaskByID(id int) (Task, error) {
	var task Task
	row := db.QueryRow(queries.GetTaskByID, id)
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
		return Task{}, err
	}
	return task, nil
}

func updateTask(id int, title, description string, done bool) error {
	_, err := db.Exec(queries.UpdateTask, title, description, done, id)
	return err
}

func deleteTask(id int) error {
	_, err := db.Exec(queries.DeleteTask, id)
	return err
}

func patchTask(id int, patchedTask Task) error {
	var setValues []string
	var args []interface{}

	if patchedTask.Title != "" {
		setValues = append(setValues, "title = ?")
		args = append(args, patchedTask.Title)
	}

	if patchedTask.Description != "" {
		setValues = append(setValues, "description = ?")
		args = append(args, patchedTask.Description)
	}

	args = append(args, id)

	if len(setValues) == 0 {
		return nil // no fields to update
	}

	_, err := db.Exec(fmt.Sprintf(queries.PatchTask, strings.Join(setValues, ", ")), args...)
	return err
}
