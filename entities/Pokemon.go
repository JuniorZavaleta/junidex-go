package entities

type PokemonType struct {
	Id   int
	Name string
}

type Pokemon struct {
	Id        int
	Name      string
	TypeOneId int
	TypeTwoId int
	TypeOne   PokemonType
	TypeTwo   PokemonType
}
