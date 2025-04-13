package mail

import (
	"arseniyms/suppliers/server/internal/controllers/companies"
	"encoding/json"
	"net/http"
	"net/smtp"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SendToMail(w http.ResponseWriter, r *http.Request) {
	var sendMail SendMail
	if err := json.NewDecoder(r.Body).Decode(&sendMail); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	smtpHost := os.Getenv(ENV_SMTP_HOST)
	smtpPort := os.Getenv(ENV_SMTP_PORT)
	from := os.Getenv(ENV_ADDRESS_FROM)
	password := os.Getenv(ENV_ADDRESS_PASSWORD_FROM)
	to := []string{sendMail.Mail}

	c, compErr := companies.GetCompanyById(w, sendMail.ExtId)

	if compErr != nil {
		if compErr == mongo.ErrNoDocuments {
			http.Error(w, "No found items", http.StatusNotFound)
			return
		}
		http.Error(w, "Error finding company", http.StatusInternalServerError)
		return
	}

	subject := "Subject: Информация о компании \n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body :=
		"<html><body>" +
			"<h1>" + c.CompanyName + "</h1>" +
			"<ul>" +
			createLi("Интегратор/Вендор: ", c.CompanyType) +
			createLi("ИНН: ", c.INN) +
			createLi("Сайт: ", c.Website) +
			createLi("ФИО: ", c.People) +
			createLi("Сотовые: ", c.Phones) +
			createLi("Эл. Почта: ", c.Emails) +
			createLi("Адрес: ", c.Address) +
			createLi("Тип ИТ оборудования: ", c.ITEquipment) +
			createLi("Наименование ПО: ", c.SoftwareName) +
			createLi("Наличие в реестре Минпромторга: ", c.IsMinPromTorg) +
			createLi("Наличие в реестре Минцифр: ", c.IsMincifr) +
			createLi("Краткое описание ИТ-решения: ", c.Description) +
			"</ul>" +
			"</body></html>"

	message := []byte(subject + mime + body)

	addr := smtpHost + ":" + smtpPort
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		http.Error(w, "Error sending email", http.StatusBadRequest)
		return
	}

	success := MailSuccess{
		Code: http.StatusOK,
		Msg:  "Mail was successfully sent",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}

func createLi(title string, info string) string {
	return "<li><b>" + title + "</b>" + info + "</li>"
}
