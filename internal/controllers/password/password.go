package password

type Admin struct {
	Password string `json:"password"`
}

type Success struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
