package commands

import (
	"context"

	"github.com/EarvinKayonga/tasks/database"
	"github.com/EarvinKayonga/tasks/models"
	"github.com/urfave/cli/v2"
)

func CreateTask(c *cli.Context) error {
	store, err := database.NewJsonFile(c.String("json_file"))
	if err != nil {
		return err
	}

	task, err := deserialize(c.Args().First())
	if err != nil {
		return err
	}

	if task.Status == "" {
		task.Status = models.StatusPending
	}

	task, err = store.CreateTask(context.Background(), *task)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}
