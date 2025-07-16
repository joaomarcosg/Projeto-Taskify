package services

import (
	"testing"
	"time"

	"github.com/joaomarcosg/Projeto-Taskify/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct{}

func (m *MockTaskStore) CreateTask(title, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) GetTaskById(id int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       "Mock Test Task",
		Description: "Mock Test Task",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) ListTasks() ([]store.Task, error) {
	return []store.Task{{
		Id:          1,
		Title:       "Mock Test Task",
		Description: "Mock Test Task",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
		{
			Id:          2,
			Title:       "Mock Test Task2",
			Description: "Mock Test Task2",
			Priority:    -1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockTaskStore) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) DeleteTask(id int32) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	task, err := taskService.Store.CreateTask("Mock Test Task", "Mock Test Task", 1)

	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Task", task.Description)
	assert.Equal(t, int32(1), task.Priority)

}

func TestGetTaskById(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	task, err := taskService.Store.GetTaskById(1)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.Id)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Task", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestListTasks(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	tasks, err := taskService.Store.ListTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "Mock Test Task", tasks[0].Title)
	assert.Equal(t, "Mock Test Task2", tasks[1].Title)
}
