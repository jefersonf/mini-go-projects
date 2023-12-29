package storage

type MongoRepository struct{}

func (*MongoRepository) init() {}

func (s *MongoRepository) Add(d Data) error {
	return nil
}

func (s *MongoRepository) GetAll() []Data {
	return make([]Data, 0)
}
