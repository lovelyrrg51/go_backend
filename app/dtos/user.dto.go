package dtos

// Normal User Response
type NormalUserResponse struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	IsEmailConfirm bool   `json:"is_email_confirm"`
}

// Register Request & Response
type AuthRegisterRequest struct {
	Username  string `json:"username" validate:"nonzero,regexp=^[0-9a-z]*$"`
	FirstName string `json:"first_name" validate:"nonzero"`
	LastName  string `json:"last_name" validate:"nonzero"`
	Email     string `json:"email" validate:"nonzero,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password  string `json:"password" validate:"nonzero"`
}
type AuthRegisterResponse struct {
	Status string             `json:"status"`
	Data   NormalUserResponse `json:"data"`
}

// Login Request & Response
type AuthLoginRequest struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}
type AuthLoginResponse struct {
	Status    string             `json:"status"`
	Data      NormalUserResponse `json:"data"`
	AuthToken string             `json:"auth_token"`
}

// Get Profile Response
type UserProfileResponse struct {
	Status    string             `json:"status"`
	Data      NormalUserResponse `json:"data"`
	AuthToken string             `json:"auth_token"`
}
