package types

import (
	"encoding/json"
	"fmt"
	"http_server/domain"
	"http_server/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetObjectHandlerRequest struct {
	Key string `json:"key"`
}

func CreateGetObjectHandlerRequest(r *http.Request) (*GetObjectHandlerRequest, error) {
	key := r.URL.Query().Get("key")
	if key == "" {
		return nil, fmt.Errorf("missing key")
	}
	return &GetObjectHandlerRequest{Key: key}, nil
}

type PutObjectHandlerRequest struct {
	domain.Object
}

func CreatePutObjectHandlerRequest(r *http.Request) (*PutObjectHandlerRequest, error) {
	var req PutObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type DeleteObjectHandlerRequest struct {
	Key string `json:"key"`
}

func CreateDeleteObjectHandlerRequest(r *http.Request) (*DeleteObjectHandlerRequest, error) {
	key := r.URL.Query().Get("key")
	if key == "" {
		return nil, fmt.Errorf("missing key")
	}
	return &DeleteObjectHandlerRequest{Key: key}, nil
}

type PostObjectHandlerRequest struct {
	domain.Object
}

func CreatePostObjectHandlerRequest(r *http.Request) (*PostObjectHandlerRequest, error) {
	var req PostObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type GetObjectHandlerResponse struct {
	Value *string `json:"value"`
}

func ProcessError(w http.ResponseWriter, err error, resp any) {
	if err == repository.NotFound {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if resp != nil {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
}

//===========
type GetTaskStatusHandlerResponse struct {
	Status  string `json:"status"`
}

type GetTaskResultHandlerResponse struct {
	Result  string `json:"result"`
}

type GetTaskHandlerRequest struct {
	Task_id string `json:"task_id"`
}

func CreateGetTaskHandlerRequest(r *http.Request) (*GetTaskHandlerRequest, error) {
	//task_id := r.URL.Query().Get("task_id")
	task_id := chi.URLParam(r, "task_id")
	if task_id == "" {
		return nil, fmt.Errorf("missing task_id")
	}
	return &GetTaskHandlerRequest{Task_id: task_id}, nil
}

type PostTaskHandlerRequest struct {
	Data string `json:"data"`
}

func CreatePostTaskHandlerRequest(r *http.Request) (*PostTaskHandlerRequest, error) {
	var req PostTaskHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
    
    return &req, nil
}