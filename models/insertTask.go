package models

import (
	"td-project/db"
)

func InsertTask(task Task) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO tasks(tittle, description, employee, done) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(sql, task.Title, task.Description, task.Employee, task.Done).Scan(&id)
	
	return
}