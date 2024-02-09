package commands

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"github.com/EarvinKayonga/tasks/database"
	"github.com/EarvinKayonga/tasks/models"
)

func RestServer(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := store.GetAllTasks(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		task, err := store.GetTaskByID(c.Request.Context(), c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.POST("/tasks", func(c *gin.Context) {
		newTask := models.Task{}
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task, err := store.CreateTask(c.Request.Context(), newTask)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		newTask := models.Task{
			ID: c.Param("id"),
		}
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task, err := store.UpdateTask(c.Request.Context(), newTask)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		task, err := store.DeleteTaskByID(c.Request.Context(), c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		c.JSON(http.StatusOK, task)
	})

	socket := fmt.Sprintf("%s:%d",
		c.String("host"),
		c.Int64("port"),
	)

	return r.Run(socket)
}
