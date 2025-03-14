package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Pray", Completed: false},
	{ID: "5", Item: "Work ", Completed: false},
	{ID: "10", Item: "Rest", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByd(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not found"})
		return
	}

	context.IndentedJSON(http.StatusCreated, todo)
}

func getTodoByd(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggleStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByd((id))

	if err != nil {
		{
			context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "To do No found"})
			return
		}
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusCreated, todo)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleStatus)
	router.Run("localhost:9090")
}
