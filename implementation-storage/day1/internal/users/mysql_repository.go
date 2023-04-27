package users

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)


type MySQLRepository struct{
	Database *sql.DB
}

func (repository *MySQLRepository) GetById(id int) (user *User, err error)  {
	//generate the query
	query:= `
		SELECT
		id, userName, email
		FROM users
		WHERE id = ?
	`
	//execute the query
	row:=repository.Database.QueryRow(query,id)

	//map the result
	err:= row.Scan(&user.Username,&user.Email,&user.ID)
	if err != nil{
		switch err{
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return
	} 
	return
}

func (repository *MySQLRepository) Create(user *User) (err error)  {
	statement, err :=repository.Database.Prepare(`
	INSERT INTO users (username, email)
		VALUES (?,?)
		`)
	if err != nil{
		return
	}
	defer statement.Close()

	result,err := statement.Exec(
		user.userName,
		user.email,
	)
	if err != nil{
		//cast to MYSQL error
		mysqlError, ok:= err.(*mysql.MySQLError)
		if !ok {
			return
		}

		switch mysqlError{
		case 1062:
			err = ErrAlreadyExists
		case 1586:
			err = ErrAlreadyExists
		}

		return
	}

	//Map the result
	lastId, err := result.LastInsertId()

	if err != nil{
		return
	}

	//return the result
	user.Id = lastId
	return
}

func (repository *MySQLRepository) Update(user *User) (err error)  {
	//prepare
	statement, err:= repository.Database.Prepare(`
		UPDATE users
		SET username = ?, email=?
		WHERE id = ?	
	`)
	if err != nil{
		return
	}

	//execute 
	_, err = statement.Exec(
		user.Username,
		user.Email,
		user.ID,
	)
	if err != nil{
		return
	}
	return
}

func (repository *MySQLRepository) Delete(id int) (err error)  {
	// generate the statement
	statement,err := repository.Database.Prepare(`
		DELETE FROM users
		WHERE id = ?
	`)
	if err != nil{
		return
	}

	_,err = statement.Exec(id)
	if err != nil{
		return
	}
	return
}

