package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"service-consult/service"
	"strings"
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
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", Username, Password, Hostname, DBName)
}

func OpenConnection() (*sql.DB, error) {
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

func (tx *repo) ConsultNegativacoes(data string) ([]*service.Data, error) {

	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, "-", "", -1)
	data = strings.Replace(data, " ", "", -1)
	log.Println(data)

	query := `SELECT
		id,
		companydocument,
		companyname,
		customerdocument,
		value,
		contract,
		debtdate,
		inclusiondate FROM testeDB.negativacoes WHERE customerdocument = ?;`

	rows, err := tx.db.Query(query, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	ngs := []*service.Data{}

	for rows.Next() {
		ng := &service.Data{}

		err := rows.Scan(&ng.ID, &ng.CompanyDocument, &ng.CompanyName, &ng.CustomerDocument, &ng.Value, &ng.Contract, &ng.DebtDate, &ng.InclusionDate)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan(): %w", err)
		}
		ngs = append(ngs, ng)
	}

	log.Println(ngs)

	return ngs, nil

}
