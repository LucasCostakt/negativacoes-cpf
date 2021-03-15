package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Storage {
	return &repo{
		db: db,
	}
}

func (tx *repo) CreatePrincipalTable() error {

	query := `CREATE TABLE IF NOT EXISTS testeDB.negativacoes (
		id INT NOT NULL AUTO_INCREMENT,
		companydocument VARCHAR(255) NULL,
		companyname VARCHAR(255) NULL,
		customerdocument VARCHAR(255) NULL,
		value FLOAT NULL,
		contract VARCHAR(255) NULL,
		debtdate DATETIME NOT NULL,
		inclusiondate DATETIME NOT NULL,
		PRIMARY KEY (id));`

	_, err := tx.db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
