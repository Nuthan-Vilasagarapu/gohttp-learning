package main

import (
	"net/http"
	"slices"
	"strings"
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

	router.PUT("/:id", func(ctx *gin.Context) {
		task_id := ctx.Params.ByName("id")
		task_body := ctx.Request.FormValue("task")
		status_body := ctx.Request.FormValue("status")

		for i, todo := range todos {
			if todo.Id == task_id {
				if task_body != "" {
					todo.Task = task_body
				}
				if status_body != "" && (status_body == "todo" || status_body == "done") {
					todo.Status = status_body
				}
				todo.UpdatedAt = time.Now()
				todos[i] = todo
				ctx.JSON(200, gin.H{
					"message": "Task updated successfully",
					"tasks":   todos,
				})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Task not Found",
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

	router.GET("/todos/search", func(ctx *gin.Context) {
		task := ctx.Query("task")
		var todoItems []ITodo

		for _, todo := range todos {
			if strings.Contains(todo.Task, task) {
				todoItems = append(todoItems, todo)
			}
		}
		if len(todoItems) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Task not Found",
			})
		} else {
			ctx.JSON(http.StatusFound, gin.H{
				"message": "Task Found",
				"todo":    todoItems,
			})
		}
	})

	router.DELETE("/:ID", func(ctx *gin.Context) {
		task := ctx.Params.ByName("ID")
		pos := -1
		for i, todo := range todos {
			if todo.Id == task {
				pos = i
				break
			}
		}
		if pos == -1 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Task not Found",
			})
		} else {
			todos = slices.Delete(todos, pos, pos+1)
			ctx.JSON(200, gin.H{
				"message": "Task deleted successfully",
				"tasks":   todos,
			})
		}
	})

	router.GET("/add", func(ctx *gin.Context) {
		ctx.File("index.html")
	})

	router.Run(":8000")
}
