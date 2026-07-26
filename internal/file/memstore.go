package file

import "errors"

type MemStore struct {
	Data map[string][]byte
}

func (s *MemStore) Save(name string, data []byte) error {
	if _, exists := s.Data[name]; exists {
		return errors.New("Name Already Taken")
	}
	s.Data[name] = data
	return nil
}

func (s *MemStore) Get(name string) ([]byte, error) {
	val, ok := s.Data[name]
	if !ok {
		return nil, errors.New("No such bytes")
	}
	return val, nil
}

func (s *MemStore) Delete(name string) error {
	_, ok := s.Data[name]
	if !ok {
		return errors.New("No such bytes")
	}
	delete(s.Data, name)
	return nil
}
