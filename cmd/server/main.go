package main

import (
	"database/sql"
	"os"

	"github.com/andarroyave/reserva-turnos/cmd/server/handler"
	"github.com/andarroyave/reserva-turnos/cmd/server/middlewares"
	"github.com/andarroyave/reserva-turnos/docs"
	"github.com/andarroyave/reserva-turnos/internal/patient"
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
// @description This API Handle Products
// @license.name Apache 2.0

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	publicKey := "public"

	datasource := "root:password@tcp(localhost:3306)/reserva-turnos"
	TurnsDB, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err = TurnsDB.Ping(); err != nil {
		panic(err)
	}

	authMid := middlewares.NewAuth(publicKey, secretKey)
	server := gin.Default()

	storage := store.SqlStore{TurnsDB}
	turnsRepo := turn.Repository{&storage}
	turnService := turn.Service{&turnsRepo}
	turnHandler := handler.TurnHandler{&turnService}

	repo := patient.NewRepository(&storage)
	service := patient.NewService(repo)
	patientHandler := handler.NewPatientHandler(service)

	

	docs.SwaggerInfo.Host = "localhost:8085"
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	patients := server.Group("/patients")
	{
		patients.GET("/getByID/:id", patientHandler.GetById())
		patients.GET("/getAll", patientHandler.GetAllPatients())
		patients.POST("", patientHandler.Post())
		patients.DELETE("/:id", patientHandler.Delete())
		patients.PUT("/:id", patientHandler.Put())
	}



	turnServer := server.Group("/turns")
	turnServer.GET("/:id", turnHandler.GetTurnById)
	turnServer.POST("/", authMid.AuthHeader, turnHandler.CreateTurn)
	turnServer.PUT("/:id", authMid.AuthHeader, turnHandler.UpdateTurn)
	turnServer.PATCH("/:id", authMid.AuthHeader, turnHandler.UpdateTurnFields)
	turnServer.DELETE("/:id", authMid.AuthHeader, turnHandler.DeleteTurn)
	turnServer.GET("/", turnHandler.GetTurnByDNI)
	server.Run(":8085")

}
