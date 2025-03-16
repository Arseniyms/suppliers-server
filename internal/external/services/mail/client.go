package mail

const (
	MAIL_PATH = "/mail/"

	ENV_ADDRESS_FROM          = "SMTP_ADDRESS_FROM"
	ENV_ADDRESS_PASSWORD_FROM = "SMTP_PASSWORD_FROM"
	ENV_SMTP_HOST             = "SMTP_HOST"
	ENV_SMTP_PORT             = "SMTP_PORT"
)

type MailSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SendMail struct {
	Mail  string `json:"mail"`
	ExtId string `json:"extId"`
}
