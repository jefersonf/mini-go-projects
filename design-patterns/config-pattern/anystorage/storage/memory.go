package storage

type MemoryRepository struct {
	data []Data
}

func WithMemoryRepository() Option {
	return func(o *StorageOption) error {
		return nil
	}
}

func (s *MemoryRepository) init() {
	if s.data == nil {
		s.data = make([]Data, 0)
	}
}

func (s *MemoryRepository) Add(d Data) error {
	s.init()
	s.data = append(s.data, d)
	return nil
}

func (s *MemoryRepository) GetAll() []Data {
	s.init()
	return s.data
}
