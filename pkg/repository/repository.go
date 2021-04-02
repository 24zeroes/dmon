package repository

type Authorization interface {
}

type DmonList interface {
}

type DmonItem interface {
}

type Repository struct {
	Authorization
	DmonList
	DmonItem
}

func NewRepository() *Repository {
	return &Repository{}
}
