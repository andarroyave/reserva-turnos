package handler

import (
	"net/http"
	"strconv"

	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/andarroyave/reserva-turnos/internal/turn"
	"github.com/andarroyave/reserva-turnos/pkg/web"
	"github.com/gin-gonic/gin"
)

type TurnHandler struct {
	TurnService turn.IService
}

func (h *TurnHandler) GetTurnById(ctx *gin.Context) {
	turnId := ctx.Param("id")
	turnIdInt, err := strconv.Atoi(turnId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}
	turnFound, errFound := h.TurnService.GetTurnById(int64(turnIdInt))
	if errFound != nil {
		println(errFound.Error())
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
		return
	}

	ctx.JSON(http.StatusOK, &turnFound)

}

func (h *TurnHandler) CreateTurn(ctx *gin.Context) {
	var tur domain.Turn
	if err := ctx.ShouldBindJSON(&tur); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}
	turnCreated, errFound := h.TurnService.CreateTurn(tur)
	if errFound != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(errFound.Error()))
		return
	}

	ctx.JSON(http.StatusOK, &turnCreated)

}

func (h *TurnHandler) UpdateTurn(ctx *gin.Context) {
	var tur domain.Turn
	if err := ctx.ShouldBindJSON(&tur); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}
	turnId := ctx.Param("id")
	turnIdInt, err := strconv.Atoi(turnId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	tur.Id = int64(turnIdInt)
	turnUpdated, errFound := h.TurnService.UpdateTurn(tur)
	if errFound != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(errFound.Error()))
		return
	}

	ctx.JSON(http.StatusOK, &turnUpdated)

}

func (h *TurnHandler) UpdateTurnFields(ctx *gin.Context) {
	var tur domain.Turn
	if err := ctx.ShouldBindJSON(&tur); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}
	turnId := ctx.Param("id")
	turnIdInt, err := strconv.Atoi(turnId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	tur.Id = int64(turnIdInt)
	turnUpdated, errFound := h.TurnService.UpdateTurnFields(tur)
	if errFound != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(errFound.Error()))
		return
	}

	ctx.JSON(http.StatusOK, &turnUpdated)

}

func (h *TurnHandler) DeleteTurn(ctx *gin.Context) {
	turnId := ctx.Param("id")
	turnIdInt, err := strconv.Atoi(turnId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}
	turnFound, errFound := h.TurnService.DeleteTurn(int64(turnIdInt))
	if errFound != nil {
		println(errFound.Error())
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
		return
	}

	ctx.JSON(http.StatusOK, &turnFound)

}

func (h *TurnHandler) GetTurnByDNI(ctx *gin.Context) {
	dni := ctx.Query("dni")
	turnsFound, errFound := h.TurnService.GetTurnByDNI(dni)
	if errFound != nil {
		println(errFound.Error())
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
		return
	}

	ctx.JSON(http.StatusOK, &turnsFound)

}
