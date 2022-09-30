package services

import (
	"context"

	"github.com/ffelipelimao/delivery-service/internal/domain"
)

type Service interface {
	Create(ctx context.Context, register domain.Register) (domain.Register, error)
	Get(ctx context.Context, ID string) (domain.Register, error)
}
