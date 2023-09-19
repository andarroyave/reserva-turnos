package main

import (
	"database/sql"

	"github.com/andarroyave/reserva-turnos/cmd/server/handler"
	"github.com/andarroyave/reserva-turnos/docs"
	"github.com/andarroyave/reserva-turnos/internal/turn"
	"github.com/andarroyave/reserva-turnos/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title PI Reserva Turnos
// @version 1.0
// @description This API Handle Products
// @license.name Apache 2.0

func main() {
	datasource := "root:password@tcp(localhost:3306)/reserva-turnos"
	TurnsDB, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err = TurnsDB.Ping(); err != nil {
		panic(err)
	}
	storage := store.SqlStore{TurnsDB}
	turnsRepo := turn.Repository{&storage}
	turnService := turn.Service{&turnsRepo}
	turnHandler := handler.TurnHandler{&turnService}

	server := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8085"
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	turnServer := server.Group("/turns")
	turnServer.GET("/:id", turnHandler.GetTurnById)
	turnServer.POST("/", turnHandler.CreateTurn)
	turnServer.PUT("/:id", turnHandler.UpdateTurn)
	turnServer.PATCH("/:id", turnHandler.UpdateTurnFields)
	turnServer.DELETE("/:id", turnHandler.DeleteTurn)
	turnServer.GET("/", turnHandler.GetTurnByDNI)
	server.Run(":8085")

}
