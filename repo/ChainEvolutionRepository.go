package repo

import (
	"../entities"
	"database/sql"
	"encoding/json"
)

func CreateEvolutionFromRow(rows *sql.Rows) (*entities.ChainEvolution, error) {
	var ChainEvolution entities.ChainEvolution

	var ChainId int
	var CurrentId int
	var EvolutionId int
	var TypeId int
	var TypeName string
	var Details string

	err := rows.Scan(
		&ChainId,
		&CurrentId,
		&EvolutionId,
		&TypeId,
		&TypeName,
		&Details,
	)

	if err == nil {
		ChainEvolution.Id = ChainId
		ChainEvolution.CurrentId = CurrentId
		ChainEvolution.EvolutionId = EvolutionId
		ChainEvolution.Type = &entities.EvolutionType{Id: TypeId, Name: TypeName}
		ChainEvolution.Evolution, _ = FindPokemon(EvolutionId)
		ChainEvolution.Current, _ = FindPokemon(CurrentId)
		json.Unmarshal([]byte(Details), &ChainEvolution.Details)
	} else {
		panic(err)
	}

	return &ChainEvolution, nil
}

func GetChainEvolution(Id int) ([]entities.ChainEvolution, error) {
	rows, err := db.Query("CALL junidex.get_complete_chain_evolution(?)", Id)

	if err != nil {
		panic(err)
	} else {
		var ChainEvolutions []entities.ChainEvolution

		for rows.Next() {
			ChainEvolution, _ := CreateEvolutionFromRow(rows)
			ChainEvolutions = append(ChainEvolutions, *ChainEvolution)
		}

		return ChainEvolutions, nil
	}

	return []entities.ChainEvolution{}, err
}
