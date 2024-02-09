package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/EarvinKayonga/tasks/models"
)

type jsonFile struct {
	*sync.Mutex
	file string
}

func NewJSONFile(file string) (TaskDB, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt stat file")
	}

	if info.IsDir() {
		return nil, fmt.Errorf("%s is a folder", file)
	}

	return &jsonFile{
		&sync.Mutex{},
		file,
	}, nil
}

func (e *jsonFile) CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	e.Lock()
	defer e.Unlock()

	tasks, err := e.loadTasks()
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	task.ID = uuid.New().String()
	tasks = append(tasks, task)

	err = e.writeTasks(tasks)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	return &task, nil
}

func (e *jsonFile) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	e.Lock()
	defer e.Unlock()

	if id == "" {
		return nil, fmt.Errorf("empty id provided")
	}

	tasks, err := e.loadTasks()
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task %s not found", id)
}
func (e *jsonFile) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	e.Lock()
	defer e.Unlock()

	tasks, err := e.loadTasks()
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	return tasks, nil
}
func (e *jsonFile) UpdateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	e.Lock()
	defer e.Unlock()

	tasks, err := e.loadTasks()
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	for i := range tasks {
		if tasks[i].ID == task.ID {
			tasks[i] = task
			err = e.writeTasks(tasks)
			if err != nil {
				return nil, errors.Wrap(err, "couldnt load tasks from file")
			}

			return &task, nil
		}
	}

	return nil, fmt.Errorf("task %s not found", task.ID)
}

func (e *jsonFile) DeleteTaskByID(ctx context.Context, id string) (*models.Task, error) {
	e.Lock()
	defer e.Unlock()

	tasks, err := e.loadTasks()
	if err != nil {
		return nil, errors.Wrap(err, "couldnt load tasks from file")
	}

	filteredTasks := []models.Task{}
	for _, task := range tasks {
		if task.ID != id {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == len(tasks) {
		return nil, fmt.Errorf("task %s not found", id)
	}

	err = e.writeTasks(filteredTasks)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("task %s not found", id)
}

func (e *jsonFile) loadTasks() ([]models.Task, error) {
	file, err := os.Open(e.file)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt open file")
	}

	defer file.Close()

	tasks := []models.Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, errors.Wrap(err, "couldnt decode tasks from file")
	}

	return tasks, nil
}

func (e *jsonFile) writeTasks(tasks []models.Task) error {
	file, err := os.Create(e.file)
	if err != nil {
		return errors.Wrap(err, "couldnt open file")
	}

	defer file.Close()

	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.Wrap(err, "couldnt marshal tasks")
	}

	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, "couldnt write file")
	}

	return nil
}
