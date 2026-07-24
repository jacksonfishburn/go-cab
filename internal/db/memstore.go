package db

import (
	"errors"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type MemStore struct {
	records map[string]file.Record
}

func (m MemStore) Put(name string, record file.Record) error {
	if _, taken := m.records[name]; taken {
		return errors.New("Name Already Taken")
	}
	m.records[name] = record
	return nil
}

func (m MemStore) Get(name string) (file.Record, error) {
	v, ok := m.records[name]
	if !ok {
		return file.Record{}, errors.New("No such Record")
	}
	return v, nil
}

func (m MemStore) List() (map[string]file.Record, error) {
	return m.records, nil
}

func (m MemStore) Delete(name string) error {
	_, ok := m.records[name]
	if !ok {
		return errors.New("No such Record")
	}
	delete(m.records, name)
	return nil
}