package service

import (
	"http_server/repository"
)

type Object struct {
	repo repository.Object
}

func NewObject(repo repository.Object) *Object {
	return &Object{
		repo: repo,
	}
}

func (rs *Object) Get(key string) (*string, error) {
	return rs.repo.Get(key)
}

func (rs *Object) Put(key string, value string) error {
	return rs.repo.Put(key, value)
}

func (rs *Object) Post(key string, value string) error {
	return rs.repo.Post(key, value)
}

func (rs *Object) Delete(key string) error {
	return rs.repo.Delete(key)
}
