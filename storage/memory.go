package storage

type Memory struct {
	currentId int
	Persons   map[int]model.Person
}
