package storage

import (
	"errors"

	"github.com/david1312/todo2/model"
)

type structDb struct{}

// func (o structDb) CreateTCus() error {
// 	db := model.ConnectDb()
// 	_, err := db.Exec("CREATE TABLE customers (id int,name varchar(200),address varchar(200)," +
// 		"phone varchar(20),birthdate date)")
// 	return err
// }
func (o structDb) List() ([]model.Customer, error) {
	db := model.ConnectDb()
	rows, err := db.Query("SELECT id,name, address, phone, birthdate FROM customers ORDER BY id asc LIMIT 10;")
	if err != nil {
		return nil, err
	}
	defer rows.Close() //defer supaya row gk ke close sampe fungsi selesai

	res := make([]model.Customer, 0)
	for rows.Next() {
		var cus model.Customer //deklarasi variable namanya cus dengan tipedata Todo
		//scan berguna menampung data yang diambil dari database kedalam variable
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone, &cus.BirthDate)
		if err != nil { //kalo scan error return null
			return nil, err
		}
		//kalo scan sukses
		res = append(res, cus)
	}
	return res, nil
}

func (o structDb) Create(obj model.Customer) error {
	db := model.ConnectDb()

	//buat create table
	// _, err2 := db.Exec("CREATE TABLE customers (id int,name varchar(200),address varchar(200)," +
	// 	"phone varchar(20),birthdate date)")
	// menghindari sql Injection
	_, err2 := db.Exec("INSERT INTO customers (id, name, address, phone, birthdate) VALUES "+
		"($1, $2, $3, $4, $5);",
		obj.ID,
		obj.Name,
		obj.Address,
		obj.Phone,
		obj.BirthDate)
	return err2
}
func (o structDb) Detail(id int32) (model.Customer, error) {
	db := model.ConnectDb()
	//siapkan
	rows, err := db.Query("SELECT id,name,address,phone,birthdate FROM customers WHERE id = $1;", id)
	if err != nil {
		return model.Customer{}, err
	}
	defer rows.Close()
	// res := make(model.Customer{}, 0) //inisialisasi variable slice, dengan size 0

	for rows.Next() {
		var cus model.Customer
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone, &cus.BirthDate)
		// err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone)
		if err != nil {
			return model.Customer{}, err
		}
		return cus, nil
	}
	return model.Customer{}, errors.New("Customercl tidak ditemukan")
}
func (o structDb) Delete(id int32) error {
	db := model.ConnectDb()
	_, err := db.Exec("DELETE from customers WHERE id = $1",
		id)
	//$ -> cara ngedefine variable di postgreSQL
	return err //kalau g

}

func (o structDb) Update(cust model.Customer, id int32) error {
	db := model.ConnectDb()
	_, err := db.Exec("UPDATE customers SET name = $1, address = $2, phone = $3, birthdate = $4 WHERE id = $5",
		cust.Name,
		cust.Address,
		cust.Phone,
		cust.BirthDate,
		id)
	//$ -> cara ngedefine variable di postgreSQL
	return err //kalau gagal return error
}

//coba git
//this is develop code
