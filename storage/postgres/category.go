package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ravshanbek-olimov/Golang-Monolit/models"

	"github.com/google/uuid"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Insert(ctx context.Context, category *models.CreateCategory) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO category (
			id,
			name,
			updated_at
		) VALUES ($1, $2, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		category.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *categoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimeryKey) (*models.Category, error) {

	query := `
		SELECT	
			id,
			name,
			created_at,
			updated_at
		FROM category
		WHERE id = $1
	`

	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   models.GetListCategoryResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			description
		FROM books
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListCategoryResponse{}, err
	}

	for rows.Next() {
		var category models.Category

		err = rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Name,
		)

		if err != nil {
			return &models.GetListCategoryResponse{}, err
		}
		resp.Categories = append(resp.Categories, &category)

	}
	return &resp, nil
}

func (r *categoryRepo) Update(ctx context.Context, category *models.UpdateCategory) error {
	query := `
		UPDATE 
			books 
		SET 
			name = $2,
			price = $3,
			description = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		category.Id,
		category.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimeryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM book_category WHERE books_id  = $1 ", req.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, "DELETE FROM books WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
