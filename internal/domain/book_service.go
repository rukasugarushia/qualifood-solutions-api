package domain

type BookService interface {
	GetAll() ([]Book, error)
	GetByID(id int) (Book, error)
	Create(book Book) error
	Update(id int, book Book) error
	Delete(id int) error
}

type bookService struct {
	books []Book
}

func NewBookService() BookService {
	return &bookService{
		books: []Book{},
	}
}

func (s *bookService) GetAll() ([]Book, error) {
	return s.books, nil
}

func (s *bookService) GetByID(id int) (Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, nil // handle error in real case
}

func (s *bookService) Create(book Book) error {
	s.books = append(s.books, book)
	return nil
}

func (s *bookService) Update(id int, updatedBook Book) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books[i] = updatedBook
			return nil
		}
	}
	return nil // handle error in real case
}

func (s *bookService) Delete(id int) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return nil // handle error in real case
}
