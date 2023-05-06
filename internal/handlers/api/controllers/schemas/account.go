package schemas

type LoginRequest struct {
	UserName string `json:"userName" validate:"required" example:"username"`
	Password string `json:"password" validate:"required" example:"mypassword"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
}

type CreateUserRequest struct {
	UserName        string `json:"userName" validate:"required" example:"username"`
	Email           string `json:"email" validate:"required,email" example:"myemail@mail.com"`
	Password        string `json:"password" validate:"required,min=8,containsany=!@#?*$" example:"yourpassword!@#$"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
