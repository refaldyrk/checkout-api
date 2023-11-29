package rest

type BaseJSONResp struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
}
