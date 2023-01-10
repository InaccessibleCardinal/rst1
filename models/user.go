package models

type User struct {
	Firstname string `db:"firstname" json:"firstname" validate:"required,min=2,max=32"`
	Lastname  string `db:"lastname" json:"lastname" validate:"required,min=2,max=32"`
	Email     string `db:"email" json:"email" validate:"required,min=4,max=50"`
	Password  string `db:"password" json:"password" validate:"required,min=4,max=120"`
	Id        int    `db:"id" json:"id" validate:"required"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Update(u *User) (int64, error)
}
