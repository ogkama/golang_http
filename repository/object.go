package repository

type Object interface {
	Get(key string) (*string, error)
	Put(key string, value string) error
	Post(key string, value string) error
	Delete(key string) error
}

type Task interface {
	Get(task_id string) (*[]string, error)
	Post(task_id string, data string) (string, error)
	GetStatus(task_id string) (*string, error)
	GetResult(task_id string) (*string, error)
	SetStatus(task_id string,status string) error
	SetResult(task_id string, status string) error
}
