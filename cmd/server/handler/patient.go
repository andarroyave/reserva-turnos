package handler

import (
	"errors"
	"strconv"
	"fmt"
	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/andarroyave/reserva-turnos/internal/patient"
	"github.com/andarroyave/reserva-turnos/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// GetPatientById godoc
// @Summary Get patient by Id
// @Tags Patient
// @Description Get patient
// @Produce json
// @Success 200
// @Router /patients/:id [get]
func (h *patientHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetById(id)
		if err != nil {
			fmt.Print(err.Error())
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// GetAllPatients godoc
// @Summary Get all patients
// @Tags Patient
// @Description Get All patients
// @Produce json
// @Success 200
// @Router /patients/GetAll [get]
func (h *patientHandler) GetAllPatients() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, err := h.s.GetAllPatients()
		if err != nil {
			web.Failure(c, 404, errors.New("patients not found"))
			return
		}
		web.Success(c, 200, patients)
	}
}

func validateEmptys(patient *domain.Patient) (bool, error) {
	if patient.DNI == "" || patient.Name == "" || patient.LastName == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// CreatePatient godoc
// @Summary Create patient
// @Tags Patient
// @Description Create patient
// @Accept json
// @Produce json
// @Success 200
// @Router /patients [post]
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 477, err)
			return
		}
		valid, err := validateEmptys(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.CreatePatient(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// DeletePatient godoc
// @Summary Delete patient by id
// @Tags Patient
// @Description Delete patient
// @Param id path int true "patient id"
// @Success 200
// @Router /patient [delete]
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.DeletePatient(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// UpdatePatient godoc
// @Summary Update patient by id
// @Tags Patient
// @Description Update patient
// @Accept json
// @Produce json
// @Param id path int true "patient id"
// @Success 200
// @Router /patients [put]
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetById(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdatePatient(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
