package repository

import (
	"context"
	"log"
	"time"
	"to-do-list/database"
	"to-do-list/models"
)

func CreateTodo(todoName, todoDescription string, todoDate time.Time) (int, error) {
	query := `INSERT INTO todo(name, description, date) VALUES($1, $2, $3) RETURNING id;`

	var id int

	err := database.DB.QueryRow(context.Background(), query, todoName, todoDescription, todoDate).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetAllTodos() ([]models.TodoResponse, error) {

	var todos []models.TodoResponse
	query := `SELECT id, name, description, date from todo;`

	rows, err := database.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.TodoResponse

		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Date); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func EditTodo(id int, name, description string, date time.Time) error {
	query := `UPDATE todo SET name=$1, description=$2, date=$3 WHERE id=$4;`

	result, err := database.DB.Exec(context.Background(), query, name, description, date, id)

	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return err
	}

	log.Printf("Updated: %v", result)
	return nil
}

func CheckTodoExistsById(id int) bool {
	query := `SELECT EXISTS(SELECT 1 FROM todo where id=$1);`
	var exists bool

	err := database.DB.QueryRow(context.Background(), query, id).Scan(&exists)

	if err != nil {
		return false
	}

	return exists
}

func DeleteTodoById(id int) error {
	query := `DELETE FROM todo where id=$1;`

	result, err := database.DB.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	log.Printf("Success: %v", result)

	return nil
}
