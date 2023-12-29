package storage

type Option func(*StorageOption) error

type StorageOption struct {
	MemoryDriver *MemoryRepository
	MongoDriver  *MongoRepository
	MySQLDriver  *MySQLRepository
}
