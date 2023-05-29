package entity

type User struct {
	ID string
	Name string
	Age int32
	Email string
	Password string
	Role string
}

type Login struct {
	Email string
	Password string
}