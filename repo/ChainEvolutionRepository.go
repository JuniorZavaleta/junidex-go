package repo

import (
	"../entities"
	"database/sql"
)

func CreateEvolutionFromRow(rows *sql.Rows) (*entities.ChainEvolution, error) {
	var ChainEvolution entities.ChainEvolution
	var ChainId int
	var CurrentId int
	var EvolutionId int
	var Current string
	var Evolution string
	var EvolutionType string
	var Details string

	err := rows.Scan(
		&ChainId,
		&CurrentId,
		&Current,
		&EvolutionId,
		&Evolution,
		&EvolutionType,
		&Details,
	)

	if err == nil {

	} else {
		panic(err)
	}

	return &ChainEvolution, nil
}