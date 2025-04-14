package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

var todos []Todo

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// add todo
	r.POST("/todos", func(c *gin.Context) {
		task := c.PostForm("task")
		todo := Todo{
			ID:        len(todos) + 1,
			Task:      task,
			Completed: false,
		}
		todos = append(todos, todo)
		c.HTML(200, "todo-item.html", todo)
	})

	// toggle todo
	r.POST("/todos/:id/toggle", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(400)
			return
		}
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Completed = !todos[i].Completed
				c.HTML(200, "todo-item.html", todos[i])
				return
			}
		}
		c.Status(404)
	})

	// delete todo
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(400)
			return
		}
		for i := range todos {
			if todos[i].ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.Status(200)
				return
			}
		}
		c.Status(404)
	})

	r.Run(":8080")
}
