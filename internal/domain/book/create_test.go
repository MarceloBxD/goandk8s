package book_test

import (
	"testing"

	"github.com/MarceloBxD/bookstore/internal/domain/book"
	book "github.com/MarceloBxD/goandk8s/internal/domain/book"
	usecase "github.com/MarceloBxD/goandk8s/internal/usecase/book"
	"github.com/stretchr/testify/mock"
	"github.com/zeebo/assert"
)

type mockRepo struct {
	mock.Mock 
}

func (m *mockRepo) Create(b *book.Book) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *mockRepo) FindByID(id uint) (*book.Book, error) { return nil, nil}
func (m *mockRepo) FindAll() ([]book.Book, error)         { return nil, nil }
func (m *mockRepo) Update(b *book.Book) error             { return nil }
func (m *mockRepo) Delete(id uint) error                  { return nil }

func TestCreateBookUseCase(t *testing.T){
	repo := new(mockRepo)
	uc := usecase.NewCreateBookUseCase(repo)

	bookData := &book.Book{Title: "Go Patterns", Author: "Bracet", Price: 49.9}
	repo.On("Create", bookData).Return(nil)

	err := uc.Execute(bookData)

	assert.Nil(t, err)
	repo.AssertExpectations(t)
}