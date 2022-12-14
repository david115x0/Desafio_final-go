package store

import (
	"database/sql"
	"desafioFinal/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type sqlStoreConsulta struct {
	db *sql.DB
}

func NewSQLStoreSqlConsulta() StoreInterfaceConsulta {
	database, err := sql.Open("mysql", "root:2319@tcp(localhost:3306)/desafio_final")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStoreConsulta{
		db: database,
	}
}

func (s *sqlStoreConsulta) GetByRg(rg int) (domain.ConsultaDTO, error) {
	query := "SELECT c.id, c.descricao, c.data, c.hora, c.paciente_id, c.dentista_id,p.id, p.nome, p.sobrenome, p.rg, p.data_cadastro,d.id, d.nome, d.sobrenome, d.matricula FROM consultas c INNER JOIN pacientes p on c.paciente_id = p.id INNER JOIN dentistas d on c.dentista_id = d.id WHERE p.rg = ? ORDER BY p.data_cadastro;"
	rows, err := s.db.Query(query, rg)
	if err != nil {
		return domain.ConsultaDTO{}, nil
	}
	var consultas domain.ConsultaDTO

	for rows.Next() {
		if err := rows.Scan(
			&consultas.Id,
			&consultas.Descricao,
			&consultas.Data,
			&consultas.Hora,
			&consultas.PacienteID,
			&consultas.DentistaID,
			&consultas.Paciente.Id,
			&consultas.Paciente.Nome,
			&consultas.Paciente.Sobrenome,
			&consultas.Paciente.RG,
			&consultas.Paciente.DataCadastro,
			&consultas.Dentista.Id,
			&consultas.Dentista.Nome,
			&consultas.Dentista.Sobrenome,
			&consultas.Dentista.Matricula,
		); err != nil {
			return consultas, err
		}
	}

	return consultas, nil
}
func (s *sqlStoreConsulta) GetByID(id int) (domain.Consulta, error) {
	var consult domain.Consulta
	rows, err := s.db.Query("select * from consultas where id = ?", id)
	if err != nil {
		log.Println(err)
		return consult, nil
	}

	for rows.Next() {
		if err := rows.Scan(
			&consult.Id,
			&consult.PacienteID,
			&consult.DentistaID,
			&consult.Data,
			&consult.Hora,
			&consult.Descricao,
		); err != nil {
			log.Println(err.Error())
			return consult, nil
		}
	}
	return consult, nil
}
func (s *sqlStoreConsulta) Update(pid int, c domain.Consulta) (domain.Consulta, error) {
	_, err := s.db.Exec("UPDATE consultas SET data= ?, hora = ?, descricao= ? WHERE id = ?", c.Data, c.Hora, c.Descricao, pid)
	if err != nil {
		log.Fatalln(err)
		return domain.Consulta{}, err
	}

	return c, nil
}

func (s *sqlStoreConsulta) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM consultas WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStoreConsulta) Create(c domain.Consulta) (domain.Consulta, error) {
	stmt, err := s.db.Prepare("INSERT INTO consultas(data, hora, descricao, paciente_id, dentista_id) VALUES (?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, errResult := stmt.Exec(
		c.Data,
		c.Hora,
		c.Descricao,
		c.PacienteID,
		c.DentistaID,
	)
	if errResult != nil {
		return domain.Consulta{}, errResult
	}
	insertedId, _ := result.LastInsertId()
	c.Id = int(insertedId)
	return c, nil
}

func (s *sqlStoreConsulta) GetByIDDTO(id int) (domain.ConsultaDTORgMatricula, error) {
	var consult domain.ConsultaDTORgMatricula
	rows, err := s.db.Query("select * from consultas where id = ?", id)
	if err != nil {
		log.Println(err)
		return consult, nil
	}

	for rows.Next() {
		if err := rows.Scan(
			&consult.Id,
			&consult.PacienteID,
			&consult.DentistaID,
			&consult.Data,
			&consult.Hora,
			&consult.Descricao,
		); err != nil {
			log.Println(err.Error())
			return consult, nil
		}
	}
	return consult, nil
}
