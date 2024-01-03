package storage

// MemoryRepository provides access to memory storage.
type MemoryRepository struct {
	data []Data
}

// WithMemoryRepository is a storage option that enables storage in memory.
func WithMemoryRepository() Option {
	return func(o *StorageOption) error {
		return nil
	}
}

// init function initialize the in-memory repository.
func (s *MemoryRepository) init() {
	if s.data == nil {
		s.data = make([]Data, 0)
	}
}

// Add persists some data in memory.
func (s *MemoryRepository) Add(d Data) error {
	s.init()
	s.data = append(s.data, d)
	return nil
}

// GetAll returns all data saved on storage.
func (s *MemoryRepository) GetAll() []Data {
	s.init()
	return s.data
}
