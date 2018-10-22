package entities

type EvolutionType struct {
	Id int
	Name string
}

type ChainEvolution struct {
	Id int
	CurrentId int
	Current Pokemon
	EvolutionId int
	Evolution Pokemon
	Type EvolutionType
}
