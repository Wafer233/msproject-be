package model

type RegisterReq struct {
	Email     string
	Name      string
	Password  string
	Password2 string
	Mobile    string
	Captcha   string
}
