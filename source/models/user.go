package models

//user struct define
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Authority `json:"authority"`
}

type Authority int

const (
	Admin Authority = iota
	Merchant
	Customer
	NotRegistered
)

//seed data
var Users = []User{
	{ID: 1, FirstName: "Ahmet", LastName: "Soran", Age: 24, Authority: Admin},
	{ID: 2, FirstName: "Burak", LastName: "Soran", Age: 28, Authority: Merchant},
	{ID: 3, FirstName: "Batuhan", LastName: "Demircan", Age: 23, Authority: Customer},
	{ID: 5, FirstName: "Cihan", LastName: "Özhan", Age: 34, Authority: Merchant},
	{ID: 8, FirstName: "İsim", LastName: "Soyisim", Age: 13, Authority: Customer},
	{ID: 8, FirstName: "İsim2", LastName: "Soyisim2", Age: 14, Authority: NotRegistered},
	{ID: 8, FirstName: "İsim3", LastName: "Soyisim3", Age: 19, Authority: NotRegistered},
}
