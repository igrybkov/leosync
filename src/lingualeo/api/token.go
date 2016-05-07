package api

type token struct {
	Username string `json:"nickname"`
	Id       int    `json:"user_id"`
	Key      string `json:"autologin_key"`
}

type loginResp struct {
	ErrorMsg string `json:"error_msg"`
	Token    token  `json:"user"`
}
