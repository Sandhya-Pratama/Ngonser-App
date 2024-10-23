package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	Userkey = "users:all"
)

type UserRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

// membuat constructor untuk dependency
func NewUserRepository(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		db:          db,
		redisClient: redisClient}
}

// menampilkan get all user
// menggunakan []*entity.User = karena akan membutuhkan data yg banyak dengan array slice of user.
func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	//melakukan returtn dari data user itu sendir, sehingga disimpan di variabel users
	users := make([]*entity.User, 0)

	val, err := r.redisClient.Get(context.Background(), Userkey).Result()
	if err != nil {
		err := r.db.WithContext(ctx).Find(&users).Error //menggunakan db untuk melakukan query ke database
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(users)
		if err != nil {
			return nil, err
		}

		// Set the data in Redis with an expiration time (e.g., 1 hour)
		err = r.redisClient.Set(ctx, Userkey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// membuat create user
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
	if err != nil {
		return err
	}
	return nil
}

// update data user byID
func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// get user by id
func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// detele user by id
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// get by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email not found")
	}
	return user, nil
}

// Update User Self
func (r *UserRepository) UpdateUserSelf(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// update user balance by id
func (r *UserRepository) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
