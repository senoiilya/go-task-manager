package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Инициализация базы данных
func initDB() {
	var err error

	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	// Создание таблицы задач, если она ещё не существует
	query := `CREATE TABLE IF NOT EXISTS tasks (
							id INTEGER PRIMARY KEY AUTOINCREMENT,
							title TEXT NOT NULL,
							description TEXT,
							done BOOLEAN NOT NULL CHECK (done IN (0, 1))
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Функция добавления новой задачи в бд
func addTask(title, description string) (int64, error) {
	query := `INSERT INTO tasks (title, description, done) VALUES (?, ?, ?)`
	result, err := db.Exec(query, title, description, false)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
