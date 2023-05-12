package database

import (
	"database/sql"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/database/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
	_ "github.com/mattn/go-sqlite3"
)

type DBController struct {
	scheme string
	db     *sql.DB
	log    *logger.DBLogger
}

func New() *DBController {
	return &DBController{scheme: "test.db",
		db: nil, log: logger.New()}
}

func (d *DBController) Create(user models.User) {
	var err error
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
	reqSql := fmt.Sprintf("INSERT INTO users (lastname,firstname,thirdname,DOB,phone,email,hash,role) VALUES"+
		" ('%s', '%s','%s','%s','%s', '%s', '%s', '%s' );", user.LastName, user.FirstName, user.ThirdName, user.DOB,
		user.Phone, user.Email, user.Hash, user.Role)
	_, err = d.db.Exec(reqSql)
	if err != nil {
		d.log.Error(err.Error())
	}
	defer d.db.Close()
}

func (d *DBController) Read(users []models.User) ([]models.User, int) {
	var err error
	var resultUsers []models.User
	var temp models.User
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
	defer d.db.Close()

	for _, usr := range users {
		reqSql := fmt.Sprintf("SELECT * FROM users WHERE (email='%s' AND hash='%s');", usr.Email, usr.Hash)

		rows, err := d.db.Query(reqSql, nil)
		if err != nil {
			d.log.Error(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&temp.ID, &temp.LastName, &temp.FirstName, &temp.ThirdName, &temp.DOB,
				&temp.Phone, &temp.Email, &temp.Hash, &temp.Role)
			if err != nil {
				d.log.Error(err.Error())
			}
			resultUsers = append(resultUsers, temp)
		}
	}
	defer d.db.Close()
	return resultUsers, len(resultUsers)
}

func (d *DBController) Update(old models.User, new models.User) {
	var err error
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
	reqSql := fmt.Sprintf("UPDATE users SET lastname = '%s', firstname = '%s',"+
		"thirdname = '%s', DOB = '%s', phone= '%s', email = '%s', hash = '%s', WHERE email= '%s' AND hash='%s';",
		new.LastName, new.FirstName, new.ThirdName, new.DOB, new.Phone, new.Email, new.Hash, old.Email, old.Hash)
	_, err = d.db.Exec(reqSql)
	if err != nil {
		d.log.Error(err.Error())
	}
	defer d.db.Close()
}

func (d *DBController) Delete(user models.User) {
	var err error
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
	sqlReq := fmt.Sprintf("DELETE FROM users WHERE id = '%d'", user.ID)
	_, err = d.db.Exec(sqlReq)
	if err != nil {
		d.log.Error(err.Error())
	}
	defer d.db.Close()
}

func (d *DBController) DBHealth() {
	var err error
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
	err = d.db.Ping()
	if err != nil {
		d.log.Error(err.Error())
	} else {
		d.log.Log("Good")
	}
	defer d.db.Close()
}
