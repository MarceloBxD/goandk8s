package book

type Repository interface {
	Create(book *Book) error 
	FindByID(id uint) (*Book, error)
	FindAll() ([]Book, error)
	Update(book *Book) (*Book, error)
	Delete(id uint) error
}