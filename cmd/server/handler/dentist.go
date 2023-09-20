package handler

import (
	"net/http"
	"strconv"

	"github.com/andarroyave/reserva-turnos/internal/dentist"
	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	DentistService dentist.Service
}

// GetDentistById godoc
// @Summary      Gets a dentist by ID
// @Description  Gets a dentist by ID from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Success      200 {object} domain.Dentist
// @Router       /dentists/{id} [get]
func (dh *DentistHandler) GetDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	dentist, err := dh.DentistService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// PutDentistById godoc
// @Summary      Modifies a dentist by ID
// @Description  Modifies a dentist by ID in the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Param        dentist body Dentist true "Dentist data"
// @Success      200 {object} domain.Dentist
// @Router       /dentists/{id} [put]
func (dh *DentistHandler) PutDentistById(ctx *gin.Context) {
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
	dentist, err := dh.DentistService.ModifyById(id, dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// DeleteDentistById godoc
// @Summary      Deletes a dentist by ID
// @Description  Deletes a dentist by ID from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Success      204 "No Content"
// @Router       /dentists/{id} [delete]
func (dh *DentistHandler) DeleteDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	err = dh.DentistService.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
