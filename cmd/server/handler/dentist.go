package handler

import (
	"net/http"
	"strconv"

	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/gin-gonic/gin"
)

type DentistsGetter interface {
	GetByMatriculation(matriculation int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
}

type DentistCreator interface {
	Create(dentist domain.Dentist) (domain.Dentist, error)
	ModifyByMatriculation(matriculation int, dentist domain.Dentist) (domain.Dentist, error)
	DeleteByMatriculation(matriculation int) error
}

type DentistsHandler struct {
	dentistsGetter  DentistsGetter
	dentistsCreator DentistCreator
}

func NewDentistsHandler(getter DentistsGetter, creator DentistCreator) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		dentistsCreator: creator,
	}
}

// GetDentistByMatriculation godoc
// @Summary      Gets a dentist by matriculation
// @Description  Gets a dentist by matriculation from the repository
// @Tags         dentists
// @Produce      json
// @Param        matriculation path int true "Matriculation"
// @Success      200 {object} domain.Dentist
// @Router       /dentists/{matriculation} [get]
func (ph *DentistsHandler) GetDentistByMatriculation(ctx *gin.Context) {
	matriculationParam := ctx.Param("matriculation")
	matriculation, err := strconv.Atoi(matriculationParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid matriculation"})
		return
	}
	dentist, err := ph.dentistsGetter.GetByMatriculation(matriculation)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// PutDentistByMatriculation godoc
// @Summary      Modifies a dentist by matriculation
// @Description  Modifies a dentist by matriculation in the repository
// @Tags         dentists
// @Produce      json
// @Param        matriculation path int true "Matriculation"
// @Param        dentist body DentistRequest true "Dentist data"
// @Success      200 {object} domain.Dentist
// @Router       /dentists/{matriculation} [put]
func (ph *DentistsHandler) PutDentistByMatriculation(ctx *gin.Context) {
	matriculationParam := ctx.Param("matriculation")
	matriculation, err := strconv.Atoi(matriculationParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid matriculation"})
		return
	}
	var dentistRequest domain.Dentist
	if err := ctx.ShouldBindJSON(&dentistRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dentist, err := ph.dentistsCreator.ModifyByMatriculation(matriculation, dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// DeleteDentistByMatriculation godoc
// @Summary      Deletes a dentist by matriculation
// @Description  Deletes a dentist by matriculation from the repository
// @Tags         dentists
// @Produce      json
// @Param        matriculation path int true "Matriculation"
// @Success      204 "No Content"
// @Router       /dentists/{matriculation} [delete]
func (ph *DentistsHandler) DeleteDentistByMatriculation(ctx *gin.Context) {
	matriculationParam := ctx.Param("matriculation")
	matriculation, err := strconv.Atoi(matriculationParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid matriculation"})
		return
	}
	err = ph.dentistsCreator.DeleteByMatriculation(matriculation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
