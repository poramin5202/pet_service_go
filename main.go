package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/rs/cors"
)

var name string
var phone string
var email string
var aca string
var dca string
var edt string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/endpoint", sendDataHandler)

	// ตั้งค่า CORS
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8081", handler)
}

func sendDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// แสดงข้อมูลที่รับมาใน Console
	//fmt.Println("Received data:", data)

	email = data["email"].(string)
	name = data["name"].(string)
	phone = data["phone"].(string)
	aca = data["aca"].(string)
	dca = data["dca"].(string)
	edt = data["edt"].(string)

	fmt.Println("Email:", email)
	fmt.Println("Name:", name)
	fmt.Println("Phone:", phone)
	fmt.Println("Aca:", aca)
	fmt.Println("Dca:", dca)
	fmt.Println("Edt:", edt)

	sendSuccessEmail()

	// ส่งข้อความกลับไปยัง client
	fmt.Fprintf(w, "Data received and processed")

}

func sendSuccessEmail() {
	from := "poramin5202@gmail.com"
	password := "cnusugspxifbtnlv"
	to := "poramin5202@gmail.com"

	msg := "Subject: Success\n\nData received and processed successfully."

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg),
	)

	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}
