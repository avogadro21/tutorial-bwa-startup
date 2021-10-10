package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tutorial_bwa_startup/handler"
	"tutorial_bwa_startup/user"
)

func main() {
	dsn := "host=localhost user=postgres password=password123 dbname=bwa_startup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	router.Run()

	//input := user.RegisterUserInput{
	//	Name:       "Gibran3",
	//	Occupation: "Programmer",
	//	Email:      "gibran3@gmail.com",
	//	Password:   "test123",
	//}


}
