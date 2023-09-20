package handler

import (
	"net/http"
	"strconv"

	"github.com/andarroyave/reserva-turnos/internal/dentist"
	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	s dentist.Service
}

func NewDentistHandler(s dentist.Service) *DentistHandler {
	return &DentistHandler{
		s: s,
	}
}

func (dh *DentistHandler) GetDentistById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		dentist, err := dh.s.GetByIdDentist(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
			return
		}
		ctx.JSON(http.StatusOK, dentist)
	}
}

func (dh *DentistHandler) CreateDentist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentistRequest domain.Dentist
		if err := ctx.ShouldBindJSON(&dentistRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dentist, err := dh.s.CreateDentist(dentistRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		ctx.JSON(http.StatusCreated, dentist)
	}
}

func (dh *DentistHandler) PutDentistById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var dentistRequest domain.Dentist
		if err := ctx.ShouldBindJSON(&dentistRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dentist, err := dh.s.ModifyByIdDentist(id, dentistRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, dentist)
	}
}

func (dh *DentistHandler) DeleteDentistById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		err = dh.s.DeleteByIdDentist(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
