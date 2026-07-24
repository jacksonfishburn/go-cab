package db

import (
	"errors"
	"maps"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type MemStore struct {
	Records map[string]file.Record
}

func (m MemStore) Put(name string, record file.Record) error {
	if _, taken := m.Records[name]; taken {
		return errors.New("Name Already Taken")
	}
	m.Records[name] = record
	return nil
}

func (m MemStore) Get(name string) (file.Record, error) {
	v, ok := m.Records[name]
	if !ok {
		return file.Record{}, errors.New("No such Record")
	}
	return v, nil
}

func (m MemStore) List() (map[string]file.Record, error) {
	return maps.Clone(m.Records), nil
}

func (m MemStore) Delete(name string) error {
	_, ok := m.Records[name]
	if !ok {
		return errors.New("No such Record")
	}
	delete(m.Records, name)
	return nil
}