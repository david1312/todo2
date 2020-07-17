package main

import (
	"log"

	"github.com/david1312/todo2/model"
	"github.com/david1312/todo2/storage"
	"github.com/kyokomi/emoji"
)

func main() {
	//bikin mock (fake implementation) this is LATIHAN IMPLEMENTASIIN DESIGN PATTERN FACTORY , SAMA ADAPTER
	// var store = getStorage(StorageTypeMock)

	//inisialisasi
	//var memStore = storage.NewMemory()
	var memStore = storage.GetStorage(storage.StorageTypeDatabase)
	//
	// if err := memStore.Create
	//CREATE
	if err := memStore.Create(model.Customer{
		ID:        2,
		Name:      "Mr B",
		Address:   "WEWORK NOBLE HOUSE",
		Phone:     "081279393023",
		BirthDate: "2020-10-08",
	}); err != nil {
		log.Fatal(err)
	}
	emoji.Println("new Customer added to db :sunglasses:")

	//DETAIL  (GET)
	var obj model.Customer
	obj, err := memStore.Detail(7)
	if err != nil {
		log.Fatal(err)
	}
	emoji.Println("Customer di temukan :blush:namanya : " + obj.Name)

	//UPDATE
	if err := memStore.Update(model.Customer{
		Name:      "Mr F UPDATED",
		Address:   "Mock 2 UPDATED",
		Phone:     "081279393023",
		BirthDate: "2020-10-08",
	}, 3); err != nil {
		log.Fatal(err)
	}
	emoji.Println("Update Success :grin:")

	//DELETE
	// var statDelete = memStore.Delete(5)
	// if statDelete != nil {
	// 	log.Fatal(statDelete)
	// }
	// log.Println("Delete success")

	//LIST
	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}
	var number int = 1
	for _, val := range list {
		log.Printf("Customer %d Nama : %s, HP : %s \n", number, val.Name, val.Phone)
		number++
	}

}
