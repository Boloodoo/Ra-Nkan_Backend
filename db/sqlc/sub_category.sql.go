// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: sub_category.sql

package db

import (
	"context"
	"time"
)

const createSubCategory = `-- name: CreateSubCategory :one
INSERT INTO sub_category (
    name,
    category_id,
    category_name
) VALUES (
    $1, $2, $3) RETURNING id, name, category_id, category_name, created_at, updated_at
`

type CreateSubCategoryParams struct {
	Name         string `json:"name"`
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (q *Queries) CreateSubCategory(ctx context.Context, arg CreateSubCategoryParams) (SubCategory, error) {
	row := q.db.QueryRowContext(ctx, createSubCategory, arg.Name, arg.CategoryID, arg.CategoryName)
	var i SubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.CategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllSubCategories = `-- name: DeleteAllSubCategories :exec
DELETE FROM sub_category
`

func (q *Queries) DeleteAllSubCategories(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllSubCategories)
	return err
}

const deleteSubCategory = `-- name: DeleteSubCategory :exec
DELETE FROM sub_category WHERE id = $1
`

func (q *Queries) DeleteSubCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSubCategory, id)
	return err
}

const getSubCategoryByCategory = `-- name: GetSubCategoryByCategory :many
SELECT id, name, category_id, category_name, created_at, updated_at FROM sub_category WHERE category_name = $1 ORDER BY id
`

func (q *Queries) GetSubCategoryByCategory(ctx context.Context, categoryName string) ([]SubCategory, error) {
	rows, err := q.db.QueryContext(ctx, getSubCategoryByCategory, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SubCategory{}
	for rows.Next() {
		var i SubCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CategoryID,
			&i.CategoryName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubCategoryById = `-- name: GetSubCategoryById :one
SELECT id, name, category_id, category_name, created_at, updated_at FROM sub_category WHERE id = $1
`

func (q *Queries) GetSubCategoryById(ctx context.Context, id int64) (SubCategory, error) {
	row := q.db.QueryRowContext(ctx, getSubCategoryById, id)
	var i SubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.CategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSubCategoryByName = `-- name: GetSubCategoryByName :one
SELECT id, name, category_id, category_name, created_at, updated_at FROM sub_category WHERE name = $1
`

func (q *Queries) GetSubCategoryByName(ctx context.Context, name string) (SubCategory, error) {
	row := q.db.QueryRowContext(ctx, getSubCategoryByName, name)
	var i SubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.CategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllSubCategory = `-- name: ListAllSubCategory :many
SELECT id, name, category_id, category_name, created_at, updated_at FROM sub_category ORDER BY id LIMIT $1 OFFSET $2
`

type ListAllSubCategoryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllSubCategory(ctx context.Context, arg ListAllSubCategoryParams) ([]SubCategory, error) {
	rows, err := q.db.QueryContext(ctx, listAllSubCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SubCategory{}
	for rows.Next() {
		var i SubCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CategoryID,
			&i.CategoryName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSubCategory = `-- name: UpdateSubCategory :one
UPDATE sub_category SET name = $2, updated_at = $3 WHERE id = $1 RETURNING id, name, category_id, category_name, created_at, updated_at
`

type UpdateSubCategoryParams struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) UpdateSubCategory(ctx context.Context, arg UpdateSubCategoryParams) (SubCategory, error) {
	row := q.db.QueryRowContext(ctx, updateSubCategory, arg.ID, arg.Name, arg.UpdatedAt)
	var i SubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.CategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
