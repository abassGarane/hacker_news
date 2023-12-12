package services

type Repository interface {
	GetRecent() (news any, err error)
}
