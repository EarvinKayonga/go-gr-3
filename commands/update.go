package commands

import (
	"context"
	"encoding/json"

	"github.com/EarvinKayonga/tasks/database"
	"github.com/EarvinKayonga/tasks/models"
	"github.com/urfave/cli/v2"
)

func UpdateTask(c *cli.Context) error {
	store, err := database.NewJsonFile(c.String("json_file"))
	if err != nil {
		return err
	}

	task, err := deserialize(c.Args().First())
	if err != nil {
		return err
	}

	task, err = store.UpdateTask(context.Background(), *task)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}

func deserialize(taskStr string) (*models.Task, error) {
	var task models.Task

	err := json.Unmarshal([]byte(taskStr), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
