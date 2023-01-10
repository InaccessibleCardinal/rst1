package db

import (
	"database/sql"
	"rst1/models"
)

type UserDBRepository struct {
	db *sql.DB
}

func (usersRepo *UserDBRepository) FindAll() ([]models.User, error) {
	var users []models.User
	rows, err := usersRepo.db.Query("select * from users1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userResponse models.User
		if err := rows.Scan(
			&userResponse.Firstname,
			&userResponse.Lastname,
			&userResponse.Email,
			&userResponse.Password,
			&userResponse.Id); err != nil {
			return nil, err
		}
		users = append(users, userResponse)
	}
	return users, nil
}

func (usersRepo *UserDBRepository) FindById(id int) (models.User, error) {
	row := usersRepo.db.QueryRow("select * from users1 where id=?", id)
	userResponse := models.User{}
	var err error
	if err = row.Scan(
		&userResponse.Firstname,
		&userResponse.Lastname,
		&userResponse.Email,
		&userResponse.Password,
		&userResponse.Id); err == sql.ErrNoRows {
		println("Id not found")
		return userResponse, err
	}
	return userResponse, err
}

func (usersRepo *UserDBRepository) Update(u *models.User) (int64, error) {
	res, err := usersRepo.db.Exec(
		"update users1 set firstname = ?, lastname = ?, email = ?, password = ? where id = ?",
		u.Firstname,
		u.Lastname,
		u.Email,
		u.Password,
		u.Id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func NewUserRepository(db *sql.DB) models.UserRepository {
	return &UserDBRepository{db: db}
}
