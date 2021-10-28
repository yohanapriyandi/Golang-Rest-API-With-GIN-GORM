package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookUpdate BookUpdate) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {

	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// convert json.Number to int64
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Title,
		Author:      bookRequest.Author,
		Rating:      int(rating),
		Discount:    int(discount),
		Price:       int(price), //convert again using int
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookUpdate BookUpdate) (Book, error) {
	book, err := s.repository.FindByID(ID)

	// convert json.Number to int64
	price, _ := bookUpdate.Price.Int64()
	rating, _ := bookUpdate.Rating.Int64()
	discount, _ := bookUpdate.Discount.Int64()

	book.Title = bookUpdate.Title
	book.Description = bookUpdate.Description
	book.Author = bookUpdate.Author
	book.Rating = int(rating)
	book.Discount = int(discount)
	book.Price = int(price) //convert again using int

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
