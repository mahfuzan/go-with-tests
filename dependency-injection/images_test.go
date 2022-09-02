package main

import (
	"testing"
)

type MockSqlRepository struct{}

func NewMockSqlRepository() ImageRepository {
	return MockSqlRepository{}
}

func (mock MockSqlRepository) Store(image Image) error {
	return nil
}

func (mock MockSqlRepository) GetList() ([]Image, error) {
	return []Image{}, nil
}

func (mock MockSqlRepository) Get(id int64) (Image, error) {
	return Image{}, nil
}

func TestMainTDD(t *testing.T) {
	mockRepository := NewMockSqlRepository()
	imageService := NewService(mockRepository)

	t.Run("insert", func(t *testing.T) {
		if err := imageService.Insert("testing", "testing", "testing"); err != nil {
			t.Errorf("Got %v, expect nil when insert mock data", err.Error())
		}
	})

	t.Run("get list", func(t *testing.T) {
		_, err := imageService.GetImageList()
		if err != nil {
			t.Errorf("Got %v, expect nil when get mock data", err.Error())
		}
	})

	t.Run("get by id", func(t *testing.T) {
		_, err := imageService.GetById(1)
		if err != nil {
			t.Errorf("Got %v, expect nil when get mock data", err.Error())
		}
	})
}
