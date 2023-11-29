package rest

type PostLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostCheckTokenReq struct {
	Token string `json:"token"`
}

type PostSetRoleFromTokenReq struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type PostRegisterReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
