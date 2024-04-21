package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"subscription/controller"
	model "subscription/models"
)

func Subscribeapi(w http.ResponseWriter, r *http.Request) {
	var jsonData model.DataPaylod
	var sendData model.EncodeData
	_ = json.NewDecoder(r.Body).Decode(&jsonData)
	sendData = model.EncodeData{Code: "200", Message: "Email Already exist in DB"}
	if controller.CheckIfAlreadyExit(jsonData.Email) {
		json.NewEncoder(w).Encode(sendData)
		return
	}
	controller.InsertIntoDatabase(jsonData)
	sendData = model.EncodeData{Code: "200", Message: "SuccessFully Inserted in DB"}
	json.NewEncoder(w).Encode(sendData)
}
func Sendnotification(w http.ResponseWriter, r *http.Request) {
	result := controller.GetAllData()
	subject := "New Content Added"
	body := fmt.Sprintf("Hello User,\n\nNew content has been added to our website: \n\nCheck it out now!\n\nBest regards,\nYour Website Team")
	message := "Subject:" + subject + "\n\n" + body
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	senderEmail := "pranjal.kumar91406@gmail.com"
	senderPassword := "rdlj klmj bvsz wjzt"
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	fmt.Println(auth)
	for _, data := range result {
		smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, senderEmail, []string{data["email"].(string)}, []byte(message))

	}
}
