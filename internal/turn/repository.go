package turn

import (
	"fmt"

	"github.com/andarroyave/reserva-turnos/internal/domain"
	"github.com/andarroyave/reserva-turnos/pkg/store"
	"github.com/andarroyave/reserva-turnos/pkg/web"
)

type IRepository interface {
	GetTurnById(id int64) (*domain.Turn, error)
	CreateTurn(t domain.Turn) (*domain.Turn, error)
	UpdateTurn(t domain.Turn) (*domain.Turn, error)
	UpdateTurnFields(t domain.Turn) (*domain.Turn, error)
	DeleteTurn(id int64) (string , error)
	GetTurnByDNI(dni string) ([]domain.Turn, error)
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) GetTurnById(id int64) (*domain.Turn, error) {
	turn, err := r.Store.GetTurnById(id)
	if err != nil {
		println(err.Error())
		return nil, web.NewNoFoundApiError(fmt.Sprintf("turn id: %d not found", id))
	}
	return turn, nil
}

func (r *Repository) CreateTurn(t domain.Turn) (*domain.Turn, error) {
	turn, err := r.Store.CreateTurn(t)
	if err != nil {
		println(err.Error())
		return nil, web.NewNoFoundApiError(fmt.Sprintf("error in turn cration"))
	}
	return turn, nil
}

func (r *Repository) UpdateTurn(t domain.Turn) (*domain.Turn, error) {
	turn, err := r.Store.UpdateTurn(t)
	if err != nil {
		println(err.Error())
		return nil, web.NewNoFoundApiError(fmt.Sprintf("error in turn update"))
	}
	return turn, nil
}

func (r *Repository) UpdateTurnFields(t domain.Turn) (*domain.Turn, error) {
	turn, err := r.Store.UpdateTurnFields(t)
	if err != nil {
		println(err.Error())
		return nil, web.NewNoFoundApiError(fmt.Sprintf("error in turn update"))
	}
	return turn, nil
}

func (r *Repository) DeleteTurn(id int64) (string , error) {
	res, err := r.Store.DeleteTurn(id)
	if err != nil {
		println(err.Error())
		return "", web.NewNoFoundApiError(fmt.Sprintf("turn id: %d not found for delete", id))
	}
	return res, nil
}

func (r *Repository) GetTurnByDNI(dni string) ([]domain.Turn, error) {
	turns, err := r.Store.GetTurnByDNI(dni)
	if err != nil {
		println(err.Error())
		return nil, web.NewNoFoundApiError(fmt.Sprintf("turn dni: %d not found", dni))
	}
	return turns, nil
}
