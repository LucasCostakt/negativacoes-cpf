package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"service-import-data/service"
	"time"

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
func DNS(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Hostname, DBName)
}

func OpenConnection() (*sql.DB, error){
	db, err := sql.Open("mysql", DNS(DBName))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", DBName)

	return db, nil
}

func (tx *repo)SaveJson(data []*service.Data) error{

	query := `INSERT INTO testeDB.negativacoes (
		companydocument,
		companyname,
		customerdocument,
		value,
		contract,
		debtdate,
		inclusiondate) VALUES (?, ?, ?, ?, ?, ?, ?) ;`

	for _, i := range data{
		_, err := tx.db.Exec(query, i.CompanyDocument, i.CompanyName, i.CustomerDocument, i.Value, i.Contract, i.DebtDate, i.InclusionDate)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil

}