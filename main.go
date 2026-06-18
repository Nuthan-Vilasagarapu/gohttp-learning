package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/**
One Form Page in /add
[Add TODO] [SUBMIT]

Home -> / -> Should have All the Todo Items

A Route to enter a TODO Item
A Route to print a Single TODO Item"

A Way to search all TODO Items
DELETE A Todo Item
UPDATE A Status of TODO Item

Clear All Todo Items

TODO{
task: <- entered by user

_id <- A Unique TODO ID
_status <- "todo/done"
_created_at <- Added When Create
_updated_at <- On Every Update
}
*/

type ITodo struct {
	Task      string    `json:"task"`
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	router := gin.Default()

	todos := []ITodo{}

	router.POST("/adding", func(c *gin.Context) {
		task := c.Request.FormValue("t_name")
		id := uuid.New()

		todo := ITodo{
			Task:      task,
			Status:    "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Id:        id.String(),
		}

		todos = append(todos, todo)

		c.JSON(200, gin.H{
			"todo": todo,
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"todos": todos,
		})
	})

	router.GET("/:todo_id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("todo_id")

		var todoItem ITodo

		for _, todo := range todos {
			if todo.Id == id {
				todoItem = todo
			}
		}
		if todoItem == (ITodo{}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Task not Found",
			})
		} else {
			ctx.JSON(http.StatusFound, gin.H{
				"message": "Task Found",
				"todo":    todoItem,
			})
		}
	})

	router.GET("/add", func(ctx *gin.Context) {
		ctx.File("index.html")
	})

	router.Run(":8000")
}
