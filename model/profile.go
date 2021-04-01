package model

type AuthenRequestBody struct {
	Token string `json:"token"`
}
type UserSSO struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Userid       string `json:"lastname"`
	Usermail     string `json:"usermail"`
	Userstatus   string `json:"userstatus"`
	Userparentid string `json:"userparentid"`
	Usertype     string `json:"usertype"`
}
