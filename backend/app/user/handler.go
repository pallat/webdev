package user

type storage interface {
	User(id string) (UserEntity, error)
}

type Handler struct {
	storage storage
}

func New(s storage) *Handler {
	return &Handler{storage: s}
}
