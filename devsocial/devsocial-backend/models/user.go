package models

type user struct {
	id         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json: "email"`
	Password   string `"json:password_hashed"`
	Created_at string `json: "created_at"`
	Updated_at string `json: ""updated_at`
}
