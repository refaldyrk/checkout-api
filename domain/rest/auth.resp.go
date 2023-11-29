package rest

type AuthLoginResp struct {
	AccessToken string `json:"access_token"`
}

type AuthCheckTokenResp struct {
	Username string `json:"username"`
}
