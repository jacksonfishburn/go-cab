package file

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

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
	BlobStore     BlobStore
	MetadataStore MetadataStore
}

func (s *Service) Ping() bool {
	return true
}

func (s *Service) Add(name string, data []byte) (Record, error) {
	record := Record{
		name:      name,
		size:      len(data),
		md5:       computeMd5(data),
		createdAt: time.Now().String(),
		updatedAt: time.Now().String(),
	}
	err := s.MetadataStore.Put(name, record)

	if err != nil {
		return Record{}, err
	}

	err = s.BlobStore.Save(name, data)

	if err != nil {
		return Record{}, err
	}

	return record, nil
}

func (s *Service) Grab(name string) ([]byte, error) {
	data, err := s.BlobStore.Get(name)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) Del(name string) error {
	err := s.MetadataStore.Delete(name)

	if err != nil {
		return err
	}

	err = s.BlobStore.Delete(name)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Peek() (map[string]Record, error) {
	records, err := s.MetadataStore.List()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func computeMd5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
