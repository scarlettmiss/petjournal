package baseRepo

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

/**
implements the repo methods that communicate with the db
*/
