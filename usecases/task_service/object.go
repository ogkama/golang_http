package task_service

import (
	"http_server/repository"
	"time"

	"log"

	"github.com/google/uuid"
)

type Task struct {
	repo repository.Task
}

func NewObject(repo repository.Task) *Task {
	return &Task{
		repo: repo,
	}
}

func (rs *Task) Get(task_id string) (*[]string, error) {
	return rs.repo.Get(task_id)
}

func (rs *Task) GetStatus(task_id string) (*string, error) {
	return rs.repo.GetStatus(task_id)
}

func (rs *Task) GetResult(task_id string) (*string, error) {
	return rs.repo.GetResult(task_id)
}

func (rs *Task) Post(data string) (string, error) {
	task_id := uuid.New().String()	
	log.Printf("Generated task_id: %s", task_id)
	rs.SomeLogic(task_id, data)
	return rs.repo.Post(task_id, data)
}

func (rs *Task) SetStatus(task_id string, status string) {
	rs.repo.SetStatus(task_id, status)
} 

func (rs *Task) SetResult(task_id string, result string) {
	rs.repo.SetResult(task_id, result)
} 

func (rs *Task) SomeLogic(task_id string, data string) {
	go func ()  {
		rs.SetStatus(task_id, "in progress")
		time.Sleep(30 * time.Second)
		rs.SetResult(task_id, "ABOBA")
		rs.SetStatus(task_id, "done")
	}()
}