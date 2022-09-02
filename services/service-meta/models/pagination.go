package models

import (
	"encoding/json"
	"io"
	"math"
)

type Pagination struct {
	Items      int64 `json:"items"`
	TotalItems int64 `json:"totalItems"`
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
}

type Pageable[T any] struct {
	Content    *[]T        `json:"content"`
	Pagination *Pagination `json:"pagination"`
}

func Paginate[T any](content *[]T, count int64, page int, size int) *Pageable[T] {
	return &Pageable[T]{
		Content: content,
		Pagination: &Pagination{
			Items:      int64(len(*content)),
			TotalItems: count,
			Page:       int64(page),
			TotalPages: int64(math.Ceil(float64(count) / float64(size))),
		},
	}
}

func (pageable *Pageable[T]) Serialize(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(pageable)
}
