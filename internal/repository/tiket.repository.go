package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewTicketRepository(db *gorm.DB, redisClient *redis.Client) *TicketRepository {
	return &TicketRepository{
		db:          db,
		redisClient: redisClient,
	}
}

// Memanggil data Ticket dari database
func (r *TicketRepository) GetAllTicket(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	val, err := r.redisClient.Get(context.Background(), "tickets_key").Result()

	if err != nil {
		// Jika tidak ada di Redis, ambil dari database
		err := r.db.WithContext(ctx).Find(&tickets).Error // SELECT * FROM tickets
		if err != nil {
			return nil, err
		}

		// Serialize data ke JSON untuk disimpan di Redis
		val, err := json.Marshal(tickets)
		if err != nil {
			return nil, err
		}

		// Simpan data di Redis dengan waktu kedaluwarsa (contoh: 1 jam)
		err = r.redisClient.Set(ctx, "tickets_key", val, time.Duration(1)*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		// Kembalikan data yang diambil dari database
		return tickets, nil
	}

	// Jika data ada di Redis, unmarshal dari JSON ke struct
	err = json.Unmarshal([]byte(val), &tickets)
	if err != nil {
		return nil, err
	}

	// Kembalikan data yang diambil dari Redis
	return tickets, nil
}

// Membuat tiker dalam database
func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	result := r.db.WithContext(ctx).Create(&ticket)
	if result.Error != nil {
		return result.Error
	}

	err := r.redisClient.Del(ctx, "tickets_key").Err()
	if err != nil {
		return err
	}

	return nil
}

// Update Ticket
func (r *TicketRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	result := r.db.WithContext(ctx).Model(&entity.Ticket{}).Where("id = ?", ticket.ID).Updates(&ticket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete Ticket
func (r *TicketRepository) DeleteTicket(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&entity.Ticket{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
