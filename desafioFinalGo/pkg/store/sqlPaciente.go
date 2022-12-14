package store

import (
	"database/sql"
	"desafioFinal/internal/domain"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type sqlStorePaciente struct {
	db *sql.DB
}

func NewSQLStoreSqlPaciente() StoreInterfacePaciente {
	database, err := sql.Open("mysql", "root:2319@tcp(localhost:3306)/desafio_final")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStorePaciente{
		db: database,
	}
}

func (s *sqlStorePaciente) GetAll() []domain.Paciente {
	var paciente domain.Paciente
	var p []domain.Paciente
	rows, err := s.db.Query("SELECT * FROM pacientes")
	if err != nil {
		log.Println(err)
		return p
	}

	for rows.Next() {
		if errs := rows.Scan(
			&paciente.Id,
			&paciente.Nome,
			&paciente.Sobrenome,
			&paciente.RG,
			&paciente.DataCadastro,
		); errs != nil {
			return p
		}
		p = append(p, paciente)
	}
	return p
}

func (s *sqlStorePaciente) GetByID(id int) (domain.Paciente, error) {
	//var paciente domain.Paciente
	//rows, err := s.db.Query("SELECT * FROM pacientes WHERE id = ?", id)
	//if err != nil {
	//	log.Println(err)
	//	return paciente, err
	//}
	//
	//for rows.Next() {
	//	if errs := rows.Scan(
	//		&paciente.Id,
	//		&paciente.Nome,
	//		&paciente.Sobrenome,
	//		&paciente.RG,
	//		&paciente.DataCadastro,
	//	); errs != nil {
	//		log.Println(err.Error())
	//		return paciente, errs
	//	}
	//}
	//return paciente, nil
	var pacient domain.Paciente
	rows, err := s.db.Query("select * from pacientes where id = ?", id)
	if err != nil {
		log.Println(err)
		return pacient, nil
	}
	for rows.Next() {
		if err := rows.Scan(&pacient.Id, &pacient.Nome, &pacient.Sobrenome, &pacient.RG, &pacient.DataCadastro); err != nil {
			log.Println(err.Error())
			return pacient, nil
		}
	}
	return pacient, nil
}
func (s *sqlStorePaciente) GetById(id int) (domain.Paciente, error) {

	var paciente domain.Paciente
	rows, err := s.db.Query("SELECT (id, nome, sobrenome, rg, data_cadastro) FROM pacientes where id = ?", id)
	if err != nil {
		log.Println(err)
		return paciente, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&paciente.Id,
			&paciente.Nome,
			&paciente.Sobrenome,
			&paciente.RG,
			&paciente.DataCadastro,
		); err != nil {
			log.Println(err.Error())
			return paciente, err
		}
	}
	return paciente, nil
}

func (s *sqlStorePaciente) Update(pid int, p domain.Paciente) (domain.Paciente, error) {
	result, err := s.db.Exec("UPDATE pacientes SET nome= ?,  sobrenome= ?,  rg= ?, data_cadastro= ? WHERE id=?", p.Nome, p.Sobrenome, p.RG, p.DataCadastro, pid)
	if err != nil {
		fmt.Println("Failed")
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Rows Affected:", rowsAffected)
	}
	return p, nil
}

func (s *sqlStorePaciente) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM pacientes WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStorePaciente) Create(p domain.Paciente) (domain.Paciente, error) {
	// Insira a data no formato yyyy-mm-dd
	stmt, err := s.db.Prepare("INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES (?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, errResult := stmt.Exec(p.Nome, p.Sobrenome, p.RG, p.DataCadastro)
	if errResult != nil {
		return domain.Paciente{}, errResult
	}
	insertedId, _ := result.LastInsertId()
	p.Id = int(insertedId)
	return p, nil
}
