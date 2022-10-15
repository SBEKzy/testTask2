package model

type User struct {
	Name     *string `bson:"name,omitempty"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
}

type UpdateUser struct {
	Email  string
	Update User
}
