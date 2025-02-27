package http

import (
	"fmt"
	"http_server/api/http/types"
	"http_server/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Object represents an HTTP handler for managing objects.
type Object struct {
	service usecases.Object
}

// Task represents an HTTP handler for managing tasks.
type Task struct {
	service usecases.Task
}

// NewHandler creates a new instance of Object.
func NewHandler(service usecases.Object) *Object {
	return &Object{service: service}
}

// NewTaskHandler creates a new instance of Task.
func NewTaskHandler(service usecases.Task) *Task {
	return &Task{service: service}
}

// @Summary Get an object
// @Description Get an object by its key
// @Tags object
// @Accept  json
// @Produce json
// @Param key query string true "Key of the object"
// @Success 200 {object} types.GetObjectHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Object not found"
// @Router /object [get]
func (s *Object) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	value, err := s.service.Get(req.Key)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{Value: value})
}

// @Summary Create or update an object
// @Description Create or update an object with the specified key and value
// @Tags object
// @Accept  json
// @Produce json
// @Param request body domain.Object true "Object data"
// @Success 200 {string} string "Object created or updated successfully"
// @Failure 400 {string} string "Bad request"
// @Router /object [put]
func (s *Object) putHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePutObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = s.service.Put(req.Key, req.Value)
	types.ProcessError(w, err, nil)
}

// @Summary Create an object
// @Description Create a new object with the specified key and value
// @Tags object
// @Accept  json
// @Produce json
// @Param request body domain.Object true "Object data"
// @Success 200 {string} string "Object created successfully"
// @Failure 400 {string} string "Bad request"
// @Router /object [post]
func (s *Object) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = s.service.Post(req.Key, req.Value)
	types.ProcessError(w, err, nil)
}

// @Summary Delete an object
// @Description Delete an object by its key
// @Tags object
// @Accept  json
// @Produce json
// @Param key query string true "Key of the object"
// @Success 200 {string} string "Object deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Object not found"
// @Router /object [delete]
func (s *Object) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateDeleteObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = s.service.Delete(req.Key)
	types.ProcessError(w, err, nil)
}

// WithObjectHandlers registers object-related HTTP handlers.
func (s *Object) WithObjectHandlers(r chi.Router) {
	r.Route("/object", func(r chi.Router) {
		r.Get("/", s.getHandler)
		r.Post("/", s.postHandler)
		r.Put("/", s.putHandler)
		r.Delete("/", s.deleteHandler)
	})
}

// @Summary Get task status
// @Description Retrieve the status of a task by its ID
// @Tags task
// @Accept  json
// @Produce json
// @Param task_id path string true "Task ID"
// @Success 200 {object} types.GetTaskStatusHandlerResponse "Task status received successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /task/{task_id}/status [get]
func (s *Task) getTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetTaskHandlerRequest(r)	
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	status, err := s.service.GetStatus(req.Task_id)
	types.ProcessError(w, err, &types.GetTaskStatusHandlerResponse{Status: *status})
}

// @Summary Get task result
// @Description Retrieve the result of a completed task by its ID
// @Tags task
// @Accept  json
// @Produce json
// @Param task_id path string true "Task ID"
// @Success 200 {object} types.GetTaskResultHandlerResponse "Task result received successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /task/{task_id}/result [get]
func (s *Task) getTaskResultHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetTaskHandlerRequest(r)	
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, err := s.service.GetResult(req.Task_id)
	types.ProcessError(w, err, &types.GetTaskResultHandlerResponse{Result: *result})
}

// @Summary Create a new task with generated task_id
// @Description Submit a new task for processing
// @Tags task
// @Accept  json
// @Produce json
// @Param request body types.PostTaskHandlerRequest true "Task data"
// @Success 200 {string} string "Task ID returned successfully"
// @Failure 400 {string} string "Bad request"
// @Router /task [post]
func (s *Task) postTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostTaskHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	task_id, err := s.service.Post(req.Data)
	_, _ = fmt.Fprint(w, task_id)
	types.ProcessError(w, err, nil)
}

// WithObjectHandlers registers task-related HTTP handlers.
func (s *Task) WithTaskHandlers(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
    	r.Post("/", s.postTaskHandler)               
    	r.Get("/{task_id}/status", s.getTaskStatusHandler) 
    	r.Get("/{task_id}/result", s.getTaskResultHandler)
	})
}
