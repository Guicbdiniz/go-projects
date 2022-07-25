package user

type CreateUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
