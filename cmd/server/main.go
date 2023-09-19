package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andarroyave/reserva-turnos/cmd/server/handler"
	"github.com/andarroyave/reserva-turnos/internal/patient"
	"github.com/andarroyave/reserva-turnos/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Digital-21@tcp(localhost:3306)/clinica-odont")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
	} else {
		fmt.Println("Conexión exitosa a la base de datos")
	}

	storage := store.SqlStore(db)
	repo := patient.NewRepository(storage)
	service := patient.NewService(repo)
	patientHandler := handler.NewPatientHandler(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	patients := r.Group("/patients")
	{
		patients.GET("/getByID/:id", patientHandler.GetByID())
		patients.GET("/getAll", patientHandler.GetAllPatients())
		patients.POST("", patientHandler.Post())
		patients.DELETE("/:id", patientHandler.Delete())
		patients.PUT("/:id", patientHandler.Put())
	}
	r.Run(":8080")
}