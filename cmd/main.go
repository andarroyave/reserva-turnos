package main

import (
	"database/sql"
	"os"

	"github.com/andarroyave/reserva-turnos/cmd/server/handler"
	"github.com/andarroyave/reserva-turnos/cmd/server/middlewares"
	"github.com/andarroyave/reserva-turnos/docs"
	"github.com/andarroyave/reserva-turnos/internal/dentist"
	"github.com/andarroyave/reserva-turnos/internal/turn"
	"github.com/andarroyave/reserva-turnos/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title PI Reserva Turnos
// @version 1.0
// @description This API Handles Products
// @license.name Apache 2.0

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	publicKey := "public"

	datasource := "root:password@tcp(localhost:3306)/reserva-turnos"
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// Configurar el almacenamiento y el servicio para Dentistas
	dentistStorage := store.NewSqlStore(db)
	dentistRepo := dentist.NewSqlRepository(dentistStorage)
	dentistService := dentist.NewService(dentistRepo)
	dentistHandler := handler.DentistHandler{DentistService: dentistService}

	// Configurar el almacenamiento y el servicio para Turnos
	turnStorage := store.NewSqlStore(db)
	turnRepo := turn.NewRepository(turnStorage)
	turnService := turn.NewService(turnRepo)
	turnHandler := handler.TurnHandler{TurnService: turnService}

	authMid := middlewares.NewAuth(publicKey, secretKey)
	server := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8085"
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rutas para Dentistas
	dentistServer := server.Group("/dentists")
	dentistServer.GET("/:id", dentistHandler.GetDentistById)
	dentistServer.POST("/", authMid.AuthHeader, dentistHandler.createDentist)
	dentistServer.PUT("/:id", authMid.AuthHeader, dentistHandler.PutDentistById)
	dentistServer.DELETE("/:id", authMid.AuthHeader, dentistHandler.DeleteDentistById)

	// Rutas para Turnos
	turnServer := server.Group("/turns")
	turnServer.GET("/:id", turnHandler.GetTurnById)
	turnServer.POST("/", authMid.AuthHeader, turnHandler.CreateTurn)
	turnServer.PUT("/:id", authMid.AuthHeader, turnHandler.UpdateTurn)
	turnServer.PATCH("/:id", authMid.AuthHeader, turnHandler.UpdateTurnFields)
	turnServer.DELETE("/:id", authMid.AuthHeader, turnHandler.DeleteTurn)
	turnServer.GET("/", turnHandler.GetTurnByDNI)

	server.Run(":8085")
}
