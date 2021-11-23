package entity

type UserFull struct {
	Password string `db:"password"`
	User
}
