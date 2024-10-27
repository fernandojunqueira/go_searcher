package sqlite

import (
	"database/sql"
	"fmt"
	"go_searcher/participant"
	"log"
	"strings"
)

type SQLite struct {
	db *sql.DB
}

func NewSqlite(db *sql.DB) participant.Repository {
	return &SQLite{
		db: db,
	}
}

func (r *SQLite) List() ([]*participant.Participant, error) {
	rows, err := r.db.Query("SELECT * FROM inscritos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []*participant.Participant
	for rows.Next() {
		var p participant.Participant
		if err := rows.Scan(&p.ID, &p.Name, &p.IsPresent); err != nil {
			return nil, err
		}
		participants = append(participants, &p)
	}

	return participants, nil
}

func (r *SQLite) Update(id string) (string, error) {
	stmt, err := r.db.Prepare(`UPDATE inscritos SET isPresent=1 WHERE id = ?`)
	if err != nil {
		return "nil", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "Nenhum registro atualizado", nil
	}

	return "Registro atualizado com sucesso", nil
}

func (r *SQLite) Search(query string) ([]*participant.Participant, error) {
	stmt, err := r.db.Prepare(`
		Select 
			id, 
			name, 
			isPresent 
		from inscritos 
		where name like ?
		order by name asc
		limit 5`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var people []*participant.Participant
	query = "%" + strings.ToLower(query) + "%"
	rows, err := stmt.Query(query, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p participant.Participant
		err = rows.Scan(&p.ID, &p.Name, &p.IsPresent)
		if err != nil {
			return nil, err
		}
		people = append(people, &p)
	}
	if len(people) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return people, nil
}

func (r *SQLite) Get(id participant.ID) (*participant.Participant, error) {
	stmt, err := r.db.Prepare(`select id, name, isPresent from inscritos where id = ?`)
	if err != nil {
		return nil, err
	}
	var p participant.Participant
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}
	err = rows.Scan(&p.ID, &p.Name, &p.IsPresent)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *SQLite) Clean() {
	updateQuery := `UPDATE inscritos SET isPresent = 0`

	result, err := r.db.Exec(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Número de registros atualizados: %d\n", rowsAffected)
}
