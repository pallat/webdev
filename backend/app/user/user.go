package user

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserEntity struct {
	*User
	CreatedDate string
	UpdatedDate string
	DeletedDate string
}
