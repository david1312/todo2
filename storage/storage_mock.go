package storage

import (
	"errors"

	"github.com/david1312/todo2/model"
)

type mock struct{}

func (o mock) List() ([]model.Customer, error) {
	result := []model.Customer{}
	result = append(result, model.Customer{
		ID:        1,
		Name:      "Mock Title",
		Address:   "Mock Description",
		Phone:     "081279393023",
		BirthDate: "2020-10-08",
	}, model.Customer{
		ID:        2,
		Name:      "Second mock Title",
		Address:   "Mock 2",
		Phone:     "081279393023",
		BirthDate: "2020-10-08",
	})
	return result, nil
}
func (o mock) Create(obj model.Customer) error {
	return nil
}
func (o mock) Detail(id int32) (model.Customer, error) {
	if id == 1 {
		return model.Customer{
			ID:      2,
			Name:    "Mock Title",
			Address: "test",
		}, nil
	}
	return model.Customer{}, errors.New("Todo tidak ditemukan")
}
func (o mock) Delete(id int32) error {
	return nil
}
func (o mock) Update(cust model.Customer, id int32) error {
	return nil
}
