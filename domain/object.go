package domain

type Object struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Task struct {
	Task_id string `json:"-"`
	Status  string `json:"-"`
	Result  string `json:"-"`
	Data    string `json:"data"`
}