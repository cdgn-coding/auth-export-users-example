package gateway

import "auth0-users-sync-job-poc/internal/business/domain"

type OperatorService interface {
	Save(operator *domain.Operator) error
	Get(ID string) (*domain.Operator, error)
}
