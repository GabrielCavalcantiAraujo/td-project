package models

import "td-project/db"

func GetTasks() (tasks []Task, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM TASKS`)
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Employee, &task.Done)
		if err != nil {
			continue
		}

		tasks = append(tasks, task)
	}
	
	return

}