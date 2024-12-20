package queries

const (
	// Создание таблицы задач, если она ещё не существует
	InitTasks = `CREATE TABLE IF NOT EXISTS tasks (
							id INTEGER PRIMARY KEY AUTOINCREMENT,
							title TEXT NOT NULL,
							description TEXT,
							done BOOLEAN NOT NULL CHECK (done IN (0, 1))
	);`
	// Вставка новой задачи
	InsertTask = `INSERT INTO tasks (title, description, done) VALUES (?, ?, ?);`
	// Получение всех задач
	GetAllTasks = `SELECT * FROM tasks;`
	// Получение одной задачи по айди
	GetTaskByID = `SELECT * FROM tasks WHERE id = ?;`
	// Обновление задачи по айди
	UpdateTask = `UPDATE tasks SET title = ?, description = ?, done = ?, WHERE id = ?`
	// Удаление задачи по айди
	DeleteTask = `DELETE FROM tasks WHERE id = ?`
)
