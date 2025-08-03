package api

import (
	"context"
	"golang-mongo/internal/application/core/domain"
	"golang-mongo/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(ctx, &payment)

	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}
