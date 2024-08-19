package repository

import (
	"errors"
	"qualifood-solutions-api/internal/domain"
)

type MemoryBookRepository struct {
	books []domain.Book
}

func NewMemoryBookRepository() *MemoryBookRepository {
	return &MemoryBookRepository{
		books: []domain.Book{},
	}
}

func (r *MemoryBookRepository) Save(book domain.Book) error {
	r.books = append(r.books, book)
	return nil
}

func (r *MemoryBookRepository) GetAll() ([]domain.Book, error) {
	return r.books, nil
}

func (r *MemoryBookRepository) GetByID(id int) (*domain.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (r *MemoryBookRepository) Update(book domain.Book) error {
	for i, b := range r.books {
		if b.ID == book.ID {
			r.books[i] = book
			return nil
		}
	}
	return errors.New("book not found")
}

func (r *MemoryBookRepository) Delete(id int) error {
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
