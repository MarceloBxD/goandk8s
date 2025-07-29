package book

import "github.com/MarceloBxD/goandk8s/internal/domain/book"

type CreateBookUseCase struct {
	Repo book.Repository
}

func NewCreateBookUseCase(repo book.Repository) *CreateBookUseCase {
	return &CreateBookUseCase{Repo: repo}
}

func (uc *CreateBookUseCase) Execute(b *book.Book) error {
	if(len(b.Title) < 3) {
		return &InvalidTitleError{}
	}

	return uc.Repo.Create(b)
}

type InvalidTitleError struct{}

func (e *InvalidTitleError) Error() string {
	return "Título deve ter no mínimo 3 caracteres"
}