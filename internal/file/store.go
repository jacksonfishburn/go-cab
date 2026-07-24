package file

type Store struct {
	dir string
}

func (s Store) Save(name string, data []byte) error {
	return nil
}

func (s Store) Get(name string) ([]byte, error) {
	return []byte(""), nil
}

func (s Store) Delete(name string) error {
	return nil
}