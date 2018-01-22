package repositories

import (
	"github.com/yurisouza/labs-go-with-docker/models"
	"github.com/yurisouza/labs-go-with-docker/util"
)

func GetAllUsers() ([]models.User, error) {

	db, errDb := GetInstanceDB()

	if errDb != nil {
		return nil, errDb
	}

	users := []models.User{}

	query := `SELECT * FROM Users`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Name, &user.Mail)
		users = append(users, user)
	}

	return users, err
}

func GetUser(id string) (*models.User, error) {

	db, errDb := GetInstanceDB()

	if errDb != nil {
		return nil, errDb
	}

	query := `SELECT * FROM Users where id='` + id + `'`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Name, &user.Mail)
		return &user, nil
	}

	return nil, err
}

func InsertUser(user *models.User) error {

	db, errDb := GetInstanceDB()

	if errDb != nil {
		return errDb
	}

	if preparedQuery, err := db.Prepare("INSERT INTO Users (id, name, mail) VALUES ($1,$2,$3);"); err != nil {
		return err
	} else if _, err := preparedQuery.Exec(util.NewUUID(), &user.Name, &user.Mail); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	db, errDb := GetInstanceDB()

	if errDb != nil {
		return errDb
	}

	if preparedQuery, err := db.Prepare("UPDATE Users SET name=$2, mail=$3 where id = $1;"); err != nil {
		return err
	} else if _, err := preparedQuery.Exec(&user.Id, &user.Name, &user.Mail); err != nil {
		return err
	}
	return nil
}

func RemoveUser(id string) error {
	db, errDb := GetInstanceDB()

	if errDb != nil {
		return errDb
	}

	if preparedQuery, err := db.Prepare("DELETE FROM Users where id = $1;"); err != nil {
		return err
	} else if _, err := preparedQuery.Exec(id); err != nil {
		return err
	}
	return nil
}

func CreateStructure() error {
	db, errDb := GetInstanceDB()

	if errDb != nil {
		return errDb
	}

	proc := `
	CREATE TABLE Users(
		id varchar,
		name varchar,
		mail varchar
	  )
	`

	if preparedQuery, err := db.Prepare(proc); err != nil {
		return err
	} else if _, err := preparedQuery.Exec(); err != nil {
		return err
	}
	return nil
}
