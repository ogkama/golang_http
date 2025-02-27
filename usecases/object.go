package usecases

type Object interface {
	Get(key string) (*string, error)
	Put(key string, value string) error 
	Post(key string, value string) error
	Delete(key string) error
}

type Task interface {
	Get(task_id string) (*[]string, error)
	Post(data string) (string, error)
	GetStatus(task_id string) (*string, error)
	GetResult(task_id string) (*string, error)
	SomeLogic(task_id string, data string)
	SetStatus(task_id string,status string) 
	SetResult(task_id string, status string) 
}