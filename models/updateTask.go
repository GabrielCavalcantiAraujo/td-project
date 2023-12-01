package models

import "td-project/db"

func UpdateTask(id int64, task Task) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE tasks SET title=$1, description=$2, employee=$3, done=$4 WHERE id=$5`,
	task.Title, task.Description, task.Employee, task.Done, task.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()	
}