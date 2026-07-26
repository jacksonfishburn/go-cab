package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jacksonfishburn/go-cab/internal/caberr"
)

type Store struct {
	dir string
}

func Open(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, caberr.Internal("open file store", err)
	}
	return &Store{dir: dir}, nil
}

func (s *Store) Save(name string, data []byte) error {
	path := filepath.Join(s.dir, name)

	_, err := os.Stat(path)
	if err == nil {
		return caberr.AlreadyExists(fmt.Sprintf("Filename '%s' is already taken", name))
	}
	if !errors.Is(err, os.ErrNotExist) {
		return caberr.Internal("file save: stat", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return caberr.Internal("file save: write", err)
	}

	return nil
}

func (s *Store) Get(name string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(s.dir, name))
	if errors.Is(err, os.ErrNotExist) {
		return nil, caberr.NotFound(fmt.Sprintf("No file with name '%s' found", name))
	}
	if err != nil {
		return nil, caberr.Internal("file get", err)
	}

	return data, nil
}

func (s *Store) Delete(name string) error {
	err := os.Remove(filepath.Join(s.dir, name))
	if errors.Is(err, os.ErrNotExist) {
		return caberr.NotFound(fmt.Sprintf("No file with name '%s' found", name))
	}
	if err != nil {
		return caberr.Internal("file delete", err)
	}

	return nil
}
