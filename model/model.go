package model

// type Todo struct {
// 	ID          int32
// 	Title       string
// 	Description string
// 	CreatedAt   time.Time
// }

//buat dulu penampung struct buat supaya fieldnya sesuai dengan database
type Customer struct {
	ID        int32
	Name      string
	Address   string
	Phone     string
	BirthDate string
}
