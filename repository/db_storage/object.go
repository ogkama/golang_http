package db_storage

import (
	"errors"
)

type Object struct {
	data map[string][]string
}

func NewObject() *Object {
	return &Object{
		data: make(map[string][]string),
	}
	
}

func (rs *Object) Get(task_id string) (*[]string, error) {
	data, exists := rs.data[task_id]
	if !exists {
		return &[]string{}, nil
	}
	return &data, nil
}

func (rs *Object) GetStatus(task_id string) (*string, error) {
	data, exists := rs.data[task_id]
	if !exists {
		empty := ""
		return &empty, errors.New("task_id does nor exists")
	}
	status := data[0]
	return &status, nil
}

func (rs *Object) GetResult(task_id string) (*string, error) {
	data, exists := rs.data[task_id]
	if !exists {
		empty := ""
		return &empty, errors.New("task_id does nor exists")
	}
	result := data[1]
	return &result, nil
}


func (rs *Object) Post(task_id string, data string) (string, error) {
	if _, exists := rs.data[task_id]; !exists {
		rs.data[task_id] = []string{"","",""}
	}
	rs.data[task_id][2] = data 
	return task_id, nil
}

func (rs *Object) SetStatus(task_id string, status string) error{
	if _, exists := rs.data[task_id]; !exists {
		rs.data[task_id] = []string{"","",""}
	}
	rs.data[task_id][0] = status
	return nil
}

func (rs *Object) SetResult(task_id string, result string) error{
	rs.data[task_id][1] = result 
	return nil
}