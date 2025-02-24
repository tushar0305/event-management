package models

import "github.com/tushar0305/event-management/db"

type User struct{
	Id 			int64	`json:"id"`
	Email 		string	`json:"email" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
}

func (u *User) Save() (int64, error) {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return 0, err
	}

	UserId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	u.Id = UserId
	return UserId, nil
}