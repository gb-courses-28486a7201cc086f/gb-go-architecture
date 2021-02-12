package repository

import (
	"errors"

	"gb-go-architecture/lesson-1/shop/models"
)

var (
	testItems = map[string]*models.Item{
		"someName1": {Name: "someName1", Price: 10},
		"someName2": {Name: "someName2", Price: 20},
		"someName3": {Name: "someName3", Price: 30},
		"someName4": {Name: "someName4", Price: 40},
		"someName5": {Name: "someName5", Price: 50},
		"someName6": {Name: "someName6", Price: 60},
	}
)

type mapDBErrorMock struct{}

// NewMapDBErrorMock creates test instance of Repository which
// retuns error on any mathod call for testing purposes
func NewMapDBErrorMock() Repository {
	return &mapDBErrorMock{}
}

func (m *mapDBErrorMock) CreateItem(item *models.Item) (*models.Item, error) {
	return nil, errors.New("mock: test error")
}

func (m *mapDBErrorMock) ListItems(filter *ItemFilter) ([]*models.Item, error) {
	return []*models.Item{}, errors.New("mock: test error")
}

func (m *mapDBErrorMock) GetItem(ID int32) (*models.Item, error) {
	return nil, errors.New("mock: test error")
}

func (m *mapDBErrorMock) DeleteItem(ID int32) error {
	return errors.New("mock: test error")
}

func (m *mapDBErrorMock) UpdateItem(item *models.Item) (*models.Item, error) {
	return nil, errors.New("mock: test error")
}

// NewMapDBMock creates in-memory instance of Repository
// for testing purposes. Repository contains some test data
func NewMapDBMock() Repository {
	// our MapDB already woks as mock;) well, let's just use it for tests
	// if we will changed realization of Repository using real DB,
	// current MapDB will be moved here
	repo := NewMapDB()

	// fill repo with data
	for _, item := range testItems {
		repo.CreateItem(item)
	}

	return repo
}
