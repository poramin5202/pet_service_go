package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var name string
var phone string
var email string
var aca string
var dca string
var edt string

type FormData struct {
	Name  string `form:"name"`
	Email string `form:"email"`
	Phone string `form:"phone"`
}

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.POST("/api/endpoint", sendDataHandler)
	router.GET("/api/endpoint2", getDataHandler)
	router.Run(":8085")

}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func getDataHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func sendDataHandler(c *gin.Context) {
	var formData FormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	name := formData.Name
	email := formData.Email
	phone := formData.Phone


	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Phone:", phone)


	sendSuccessEmail(name, email, phone)

	c.JSON(http.StatusOK, gin.H{"message": "Data received and processed"})
}

func sendSuccessEmail(name, email, phone string) {
	from := "poramin@gmail.com"
	password := "poramin" // แทนที่ด้วยรหัสแอปพาสเวิร์ดจริงของคุณ
	to := "poramin@gmail.com"

	msg := fmt.Sprintf("Subject: Get free estimate\n\nName: %s\nPhone: %s\nEmail: %s", name, phone, email)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Get free estimate")
	m.SetBody("text/plain", msg)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
	}
}
