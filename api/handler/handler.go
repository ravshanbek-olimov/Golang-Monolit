package handler

import (
	"github.com/ravshanbek-olimov/Golang-Monolit/storage"
)

type Handler struct {
	storage storage.StorageI
}

func NewHandler(storage storage.StorageI) *Handler {
	return &Handler{
		storage: storage,
	}
}
