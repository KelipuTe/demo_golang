// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"demo-golang/microservice/internal/repo"
	"demo-golang/microservice/internal/repo/dao"
	"demo-golang/microservice/internal/service"
	"demo-golang/microservice/internal/web"
	"demo-golang/microservice/ioc"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitServer() *gin.Engine {
	db := ioc.InitGorm()
	userDao := dao.NewUserDao(db)
	userRepo := repo.NewUserRepo(userDao)
	userService := service.NewUserService(userRepo)
	userHandler := web.NewUserHandler(userService)
	engine := ioc.InitWebServer(userHandler)
	return engine
}
