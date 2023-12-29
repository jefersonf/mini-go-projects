package storage

type MySQLRepository struct{}

func (*MySQLRepository) init() {}

func (s *MySQLRepository) Add(d Data) error {
	return nil
}

func (s *MySQLRepository) GetAll() []Data {
	return make([]Data, 0)
}
