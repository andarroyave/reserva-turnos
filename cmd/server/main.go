package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	//"github.com/joho/godotenv"
	"github.com/andarroyave/reserva-turnos/cmd/server/handler"
	"github.com/andarroyave/reserva-turnos/internal/turn"
	"github.com/andarroyave/reserva-turnos/pkg/store"
)

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
	turnServer := server.Group("/turns")
	turnServer.GET("/:id", turnHandler.GetTurnById)
	turnServer.POST("/", turnHandler.CreateTurn)
	turnServer.PUT("/:id", turnHandler.UpdateTurn)
	turnServer.PATCH("/:id", turnHandler.UpdateTurnFields)
	turnServer.DELETE("/:id", turnHandler.DeleteTurn)
	turnServer.GET("/", turnHandler.GetTurnByDNI)
	server.Run(":8085")

}
