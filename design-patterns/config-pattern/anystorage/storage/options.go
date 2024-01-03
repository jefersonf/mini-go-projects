package storage

// Option is a function that represents a storage option setter.
type Option func(*StorageOption) error

// StorageOption represents the storage options.
type StorageOption struct {
	MemoryDriver *MemoryRepository
	MongoDriver  *MongoRepository
	MySQLDriver  *MySQLRepository
}
