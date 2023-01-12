package storage

import (
	"context"

	"github.com/ravshanbek-olimov/Golang-Monolit/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Category() CategoryRepoI
}

type BookRepoI interface {
	Insert(context.Context, *models.CreateBook) (string, error)
	GetByID(context.Context, *models.BookPrimeryKey) (*models.Book, error)
}

type CategoryRepoI interface {
	Insert(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimeryKey) (*models.Category, error)
}
