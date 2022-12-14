package store

import (
	"database/sql"
	"desafioFinal/internal/domain"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type sqlStore struct {
	db *sql.DB
}

func NewSQLStoreSql() StoreInterfaceDentista {
	database, err := sql.Open("mysql", "root:2319@tcp(localhost:3306)/desafio_final")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}
func (s *sqlStore) GetAll() []domain.Dentista {
	var dentist domain.Dentista
	var dentists []domain.Dentista
	rows, err := s.db.Query("SELECT * FROM dentistas")
	if err != nil {
		log.Println(err)
		return dentists
	}

	for rows.Next() {
		if err := rows.Scan(
			&dentist.Id,
			&dentist.Nome,
			&dentist.Sobrenome,
			&dentist.Matricula,
		); err != nil {
			return dentists
		}
		dentists = append(dentists, dentist)
	}
	return dentists
}

func (s *sqlStore) GetByID(id int) (domain.Dentista, error) {
	var dentist domain.Dentista
	rows, err := s.db.Query("SELECT * FROM dentistas WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return dentist, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&dentist.Id,
			&dentist.Nome,
			&dentist.Sobrenome,
			&dentist.Matricula,
		); err != nil {
			log.Println(err.Error())
			return dentist, err
		}
	}
	return dentist, nil
}

func (s *sqlStore) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	_, err := s.db.Exec("UPDATE dentistas SET nome = ? ,sobrenome = ?,matricula = ? WHERE id = ?", d.Nome, d.Sobrenome, d.Matricula, id)
	if err != nil {
		log.Fatalln(err)
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *sqlStore) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM dentistas WHERE id=?", id)
	if err != nil {
		return err
	}
	count, err := result.LastInsertId()
	if count == 0 {
		return errors.New("dentist not found at database")
	}
	return nil
}
func (s *sqlStore) GetById(id uint64) (domain.Dentista, error) {

	var dentista domain.Dentista
	rows, err := s.db.Query("SELECT (id, nome, sobrenome, matricula) FROM dentistas where id = ?", id)
	if err != nil {
		log.Println(err)
		return dentista, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&dentista.Id,
			&dentista.Nome,
			&dentista.Sobrenome,
			&dentista.Matricula,
		); err != nil {
			log.Println(err.Error())
			return dentista, err
		}
	}
	return dentista, nil
}

func (s *sqlStore) Create(d domain.Dentista) (domain.Dentista, error) {
	result, err := s.db.Exec("INSERT INTO dentistas (nome,sobrenome,matricula) VALUES (?,?,?)", d.Nome, d.Sobrenome, d.Matricula)
	if err != nil {
		fmt.Println(err.Error())
		return domain.Dentista{}, nil
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Dentista{}, nil
	}
	d.Id = int(lastInsertId)
	return d, nil
}
