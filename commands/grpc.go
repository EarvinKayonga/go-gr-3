package commands

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/EarvinKayonga/tasks/database"
	v1 "github.com/EarvinKayonga/tasks/gen/protocol/v1"
	"github.com/EarvinKayonga/tasks/gen/protocol/v1/tasksv1connect"
)

func GRPCServer(c *cli.Context) error {
	_, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	impl := &impl{}

	mux := http.NewServeMux()
	path, handler := tasksv1connect.NewTaskServiceHandler(impl)
	mux.Handle(path, handler)

	socket := fmt.Sprintf("%s:%d",
		c.String("host"),
		c.Int64("port"),
	)

	return http.ListenAndServe(
		socket,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type impl struct{}

func (e *impl) GetTasks(context.Context, *connect.Request[v1.GetTasksRequest]) (*connect.Response[v1.Task], error) {

	return nil, fmt.Errorf("implmented")
}

// Get a specific task
func (e *impl) GetTask(context.Context, *connect.Request[v1.GetTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, fmt.Errorf("implmented")
}

// Create a new task
func (e *impl) CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, fmt.Errorf("implmented")
}

// Update an existing task
func (e *impl) UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, fmt.Errorf("implmented")
}

// Delete a task
func (e *impl) DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, fmt.Errorf("implmented")
}
