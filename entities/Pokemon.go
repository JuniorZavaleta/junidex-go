package entities

import "./base"

type PokemonType struct {
	Id   int	`json:",omitempty"`
	Name string `json:",omitempty"`
}

type Pokemon struct {
	Id        	int
	Name      	string
	TypeOneId 	int
	TypeTwoId 	base.IntNull
	TypeOne   	*PokemonType
	TypeTwo   	*PokemonType `json:",omitempty"`
	HasPreEvol 	bool
}
