package main

import (
	"github.com/SBEKzy/testTask2/config"
	"github.com/SBEKzy/testTask2/handler"
	"github.com/SBEKzy/testTask2/repository"
	"github.com/SBEKzy/testTask2/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.New()
	db := repository.ConnectDB()
	repo := repository.New(db)
	s := service.New(repo)
	h := handler.New(s)
	r := gin.Default()
	r.POST("user", h.CreateUser)
	r.GET("user/:email", h.GetUser)
	r.PUT("user", h.UpdateUser)
	r.DELETE("user/:email", h.DeleteUser)
	r.Run(":" + viper.GetString("PORT"))
}
