package storage

import (
	"errors"
	"morrah77.org/location_tracker/domain"
)

var errorNoEntity = `No entities found!`

type Storage struct {
	entities map[string][]*domain.Location
}

func NewStorage() *Storage {
	entities := make(map[string][]*domain.Location, 0)
	return &Storage{
		entities: entities,
	}
}

func (storage *Storage) Fetch(key string, depth int) (interface{}, error) {
	values, ok := storage.entities[key]
	if !ok {
		return nil, errors.New(errorNoEntity)
	}
	if depth == -1 {
		return values, nil
	}
	if depth > len(values){
		depth = len(values)
	}
	return values[:depth], nil
}

// TODO improve it!
func (storage *Storage) Store(key string, value interface{}) (error) {
	values, ok := storage.entities[key]
	if !ok {
		storage.entities[key] = make([]*domain.Location, 1)
		storage.entities[key][0] = value.(*domain.Location)
		return nil
	}
	newEntities := make([]*domain.Location, 1)
	newEntities[0] = value.(*domain.Location)
	storage.entities[key] = append(newEntities, values...)
	return nil
}

func (storage *Storage) Delete(key string) (error) {
	_, ok := storage.entities[key]
	if ok {
		delete(storage.entities, key)
		return nil
	}
	return errors.New(errorNoEntity)
}
