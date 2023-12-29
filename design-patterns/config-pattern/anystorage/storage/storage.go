package storage

type Data struct{}

type Repository interface {
	Add(Data) error
	GetAll() []Data
}

type Storage struct {
	StorageOption
}

func NewStorage(opts ...Option) *Storage {
	options := defaultStorageOptions()
	for _, opt := range opts {
		opt(&options)
	}
	return &Storage{options}
}

func defaultStorageOptions() StorageOption {
	return StorageOption{MemoryDriver: &MemoryRepository{}}
}
