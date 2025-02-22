package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"
	"to-do-list/database"
	"to-do-list/models"
	"to-do-list/repository"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo *models.TodoRequest

	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Printf("Error while binding json %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if database.DB == nil {
		log.Print("Error in Database Connection")
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", todo.Date)
	if err != nil {
		log.Printf("Error while parsing date %v", err)
	}

	todoId, err := repository.CreateTodo(todo.Name, todo.Description, parsedDate)

	if err != nil {
		log.Print("Error while inserting")
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": todoId})
}

func GetAllTodos(c *gin.Context) {
	if database.DB == nil {
		log.Print("Error in Database Connection")
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
		return
	}

	todos, err := repository.GetAllTodos()

	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("No Todos Found")
			c.JSON(http.StatusNoContent, gin.H{"res": "No Todos Found"})
		} else {
			log.Printf("Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todos})
}

func EditTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.TodoRequest

	intParamId, err := strconv.Atoi(id)

	if err != nil {
		log.Print("Error while parsing param id")
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Print("Error while parsing body")
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", todo.Date)

	if err != nil {
		log.Print("Error while parsing date")
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if database.DB == nil {
		log.Print("Error in Database Connection")
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
		return
	}

	exists := repository.CheckTodoExistsById(intParamId)

	if !exists {
		log.Print("Todo does not exists")
		c.JSON(http.StatusNoContent, gin.H{"res": "Todo does not exists"})
		return
	}

	err = repository.EditTodo(intParamId, todo.Name, todo.Description, parsedDate)

	if err != nil {
		log.Print("Error while Updating")
		c.JSON(http.StatusInternalServerError, gin.H{"res": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Updated"})
}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")

	intParamId, err := strconv.Atoi(id)

	if err != nil {
		log.Print("Error while parsing param id")
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	exists := repository.CheckTodoExistsById(intParamId)

	if !exists {
		log.Print("Todo does not exists")
		c.JSON(http.StatusNoContent, gin.H{"res": "Todo does not exists"})
		return
	}

	err = repository.DeleteTodoById(intParamId)

	if err != nil {
		log.Printf("Error while deleting %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Deleted..."})
}
