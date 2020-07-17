package storage

import (
	"errors"

	"github.com/david1312/todo2/model"
)

type memory struct {
	store map[int32]model.Customer
}

func newMemory() memory {
	return memory{
		store: make(map[int32]model.Customer),
	}
}

func (o memory) List() ([]model.Customer, error) {
	//looping
	result := []model.Customer{}
	for _, v := range o.store {
		result = append(result, v)
	}
	return result, nil
}
func (o memory) Create(obj model.Customer) error {
	//menyimpan
	o.store[obj.ID] = obj
	return nil
}
func (o memory) Detail(id int32) (model.Customer, error) {
	//untuk ngecek apakah id tertentu exist dalam sebuah array atau map untuk menyambungkan if pakai semicolon
	//ini gaperlu dicasting intnya karena datanya static
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Customer{}, errors.New("Customer tidak ditemukan")
}
func (o memory) Delete(id int32) error {
	delete(o.store, id)
	return nil
}
func (o memory) Update(cust model.Customer, id int32) error {
	return nil
}
