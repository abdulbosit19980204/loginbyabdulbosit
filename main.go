package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQLga ulanish
	db, err := sql.Open("mysql", "root:1101@tcp(127.0.0.1:3306)/telegram")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Gin server yaratish
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "index.html")
	// Login formasi uchun endpoint yaratish
	router.GET("/login", func(c *gin.Context) {

		c.HTML(200, "login.html", gin.H{"message": "login.html"})

	})

	// Login formasi uchun endpoint yaratish
	router.POST("/login", func(c *gin.Context) {
		// Foydalanuvchi kiritgan ma'lumotlarni olish
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Ma'lumotlar bazasidagi foydalanuvchilar jadvalidan foydalanuvchini tekshirish
		var result string
		err = db.QueryRow("SELECT username FROM users WHERE username=? AND password=?", username, password).Scan(&result)
		if err != nil {
			// Foydalanuvchi aniqlanmadi
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}

		// Foydalanuvchi aniqlandi
		//c.JSON(200, gin.H{"message": fmt.Sprintf("Welcome, %s!", result)})
		c.HTML(200, "index.html", gin.H{"message": "index.html"})
	})

	// Serverni ishga tushurish
	router.Run()
}
