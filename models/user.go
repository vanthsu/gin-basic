package models

import (
	"errors"

	"github.com/vanthsu/gin-basic/db"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"` // struct tag 使用空白分隔
}

func GetAllUsers() (*[]User, error) {
	query := "SELECT * FROM users ORDER BY id"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 類似 stmt.Close()，這邊也需呼叫 rows.Close() 以釋放資源

	var users []User
	for rows.Next() {
		var user User

		// 要去拿到單一 row 的資料欄位做法，是用 Scan()，每個欄位要放一個變數的指標進去給他存資料
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user User
	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Save() (*User, error) {
	// 和 Laravel ORM 跨db保持一致的方式不同，Golang 的 Postgres Driver 比較底層(?)
	// 需依照各 DB 的特性處理 :
	// 	PostgreSQL 的 placeholder 使用 $1, $2, $3
	// 	PostgreSQL 不支援 LastInsertId，需要使用 RETURNING id
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// 這邊不能使用 Exec()，而要改用 QueryRow() 的方式，把 id 拿回來
	row := stmt.QueryRow(u.Name)
	err = row.Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Update() (*User, error) {
	query := "UPDATE users SET name = $1 WHERE id = $2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.ID)
	if err != nil {
		return nil, err
	}

	counts, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if counts <= 0 {
		return nil, errors.New("Update failed")
	}

	return u, nil
}

func (u *User) Delete() error {
	query := "DELETE FROM users WHERE id = $1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.ID)
	if err != nil {
		return err
	}

	counts, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if counts <= 0 {
		return errors.New("Delete failed")
	}

	return nil
}
