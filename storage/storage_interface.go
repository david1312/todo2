package storage

import "github.com/david1312/todo2/model"

//sebuah interface dengan 3 signature func
type Storage interface {
	//func create menerima object Todo dari model
	Create(obj model.Customer) error
	Detail(id int32) (model.Customer, error)
	Delete(id int32) error

	//tugas bikin fungsi list buat nyimpen datanya ke array lalu di print dalam looping
	//yang storage 2 data yg mock juga 2 data
	List() ([]model.Customer, error)
	Update(obj model.Customer, id int32) error
	//coba implementasi db
	// ListDatabase([]model.Customer, error)
}
