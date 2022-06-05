package domain

type MyData struct {
}

type IdentifiedNamedObject struct {
	Id        int64
	Name      string
	Code      string
	CreatedOn int64
	CreatedBy int64
	UpdatedOn int64
	UpdatedBy int64
	Deletedon int64
	DeletedBy int64
	IsDeleted bool
}

type User struct {
	IdentifiedNamedObject
	Password string
	Email    string
	Points   int
}

type Role struct {
	IdentifiedNamedObject
}

type Permission struct {
	IdentifiedNamedObject
}

type Game struct {
	IdentifiedNamedObject
}

type Channel struct {
	IdentifiedNamedObject
}

type Rules struct {
	IdentifiedNamedObject
}

type WishTicket struct {
	IdentifiedNamedObject
	Content string
}
