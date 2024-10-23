package service

import (
	"context"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
)

// interface untuk service
// untuk memanngil repository
type TicketUseCase interface {
	GetAllTicket(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
}

type TicketRepository interface {
	GetAllTicket(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
}

type TicketService struct {
	repository TicketRepository
}

func NewTicketService(repository TicketRepository) *TicketService {
	return &TicketService{
		repository: repository,
	}
}

func (s *TicketService) GetAllTicket(ctx context.Context) ([]*entity.Ticket, error) {
	return s.repository.GetAllTicket(ctx)
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.repository.CreateTicket(ctx, ticket)
}

func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.repository.UpdateTicket(ctx, ticket)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int64) error {
	return s.repository.DeleteTicket(ctx, id)
}
