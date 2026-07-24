package file

type MetadataStore interface {
    Put(name string, record Record) error
    Get(name string) (Record, error)
    List() (map[string]Record, error)
    Delete(name string) error
}

type BlobStore interface {
    Save(name string, data []byte) error
    Get(name string) ([]byte, error)
    Delete(name string) error
}

type Service struct {
	BlobStore BlobStore
	MetadataStore MetadataStore
}

func (m *Service) Ping() bool {
	return true
}

func (m *Service) Add() error {
	return nil
}

func (m *Service) Grab() error {
	return nil
}

func (m *Service) Del() error {
	return nil
}

func (m *Service) Peek() error {
	return nil
}
