package services

import (
	"context"

	"github.com/ffelipelimao/delivery-service/internal/domain"
)

type Service interface {
	CreateRegister(ctx context.Context, register domain.Register) (domain.Register, error)
}
