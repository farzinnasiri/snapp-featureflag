package feature

type QueryHandler struct {
	Repository
}

func NewQueryHandler(repository Repository) QueryHandler {
	return QueryHandler{repository}
}
