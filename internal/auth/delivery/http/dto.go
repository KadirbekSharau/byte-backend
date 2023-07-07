package http

type signUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}
