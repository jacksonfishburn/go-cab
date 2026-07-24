package file

type MetadataStore interface {
    Put(name string, record FileRecord) error
    Get(name string) (FileRecord, error)
    List() (map[string]FileRecord, error)
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
