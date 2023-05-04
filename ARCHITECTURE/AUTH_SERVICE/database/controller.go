package database

import (
	"database/sql"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/database/logger"
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

func (d *DBController) OpenConn() {
	var err error
	d.db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		d.log.Error(err.Error())
	}
}

func (d *DBController) Controller() {
	_, err := d.db.Exec("INSERT INTO users VALUES (1,'Иванов', 'Иван','Иваныч','21.01.1900','788888', 'gg@mail.ru', '123123', '1');")
	if err != nil {
		d.log.Error(err.Error())
	}
}

func (d *DBController) Create() {
	_, err := d.db.Exec("INSERT INTO users VALUES (1,'Иванов', 'Иван','Иваныч','21.01.1900','788888', 'gg@mail.ru', '123123', '1');")
	if err != nil {
		d.log.Error(err.Error())
	}
}

func (d *DBController) Read() {
	_, err := d.db.Exec("INSERT INTO users VALUES (1,'Иванов', 'Иван','Иваныч','21.01.1900','788888', 'gg@mail.ru', '123123', '1');")
	if err != nil {
		d.log.Error(err.Error())
	}
}

func (d *DBController) Update() {
	_, err := d.db.Exec("INSERT INTO users VALUES (1,'Иванов', 'Иван','Иваныч','21.01.1900','788888', 'gg@mail.ru', '123123', '1');")
	if err != nil {
		d.log.Error(err.Error())
	}
}

func (d *DBController) Delete() {
	_, err := d.db.Exec("INSERT INTO users VALUES (1,'Иванов', 'Иван','Иваныч','21.01.1900','788888', 'gg@mail.ru', '123123', '1');")
	if err != nil {
		d.log.Error(err.Error())
	}
}
