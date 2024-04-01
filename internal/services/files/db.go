package files

// ---------------------------------------------------------------------------------------------------------------------
// Init DB Repo
// ---------------------------------------------------------------------------------------------------------------------

type fileManager interface {
	add(file *File) (*File, error)
	find() ([]*File, error)
	findByID(id int) (*File, error)
}

type dbRepo struct{}

var db fileManager = (*dbRepo)(nil)
var _ fileManager = db // TODO: quickfix for now since db will get used later!!!

// -----------------------------------------------------------------------------
// DB Repo Functions
// -----------------------------------------------------------------------------

func (m *dbRepo) add(file *File) (*File, error) {
	return nil, nil
}

func (m *dbRepo) find() ([]*File, error) {
	return nil, nil
}

func (m *dbRepo) findByID(id int) (*File, error) {
	return nil, nil
}
