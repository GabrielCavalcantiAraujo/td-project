package models

type Task struct {
	ID          int64   `json:"id"`
	Title   	string  `json:"title"`
	Description	string  `json:"description"`
	Employee	string	`json:"employee"`
	Done		bool 	`json:"done"`
}