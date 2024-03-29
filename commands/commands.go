package commands

import (
	"github.com/urfave/cli/v2"
)

func Create() *cli.App {
	return &cli.App{
		Name: "tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "json_file",
				Aliases: []string{"f"},
				Value:   "tasks.json",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all tasks",
				Action: func(c *cli.Context) error {
					return ListAll(c)
				},
			},
			{
				Name:  "get",
				Usage: "Get a task by ID",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					return GetTaskByID(c)
				},
			},
			{
				Name:  "create",
				Usage: "Create a new task",
				Action: func(c *cli.Context) error {
					return CreateTask(c)
				},
			},
			{
				Name:  "update",
				Usage: "Update a task",
				Action: func(c *cli.Context) error {
					return UpdateTask(c)
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a task by ID",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					return DeleteTaskByID(c)
				},
			},
			{
				Name:  "http",
				Usage: "REST API",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   8000,
					},
					&cli.StringFlag{
						Name:  "host",
						Value: "127.0.0.1",
					},
				},
				Action: func(c *cli.Context) error {
					return RestServer(c)
				},
			},
			{
				Name:  "grpc",
				Usage: "GRPC API",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   8000,
					},
					&cli.StringFlag{
						Name:  "host",
						Value: "127.0.0.1",
					},
				},
				Action: func(c *cli.Context) error {
					return GRPCServer(c)
				},
			},
		},
	}
}
