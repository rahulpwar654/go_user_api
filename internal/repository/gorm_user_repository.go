package repository

import (
	"errors"

	"gin-user-api/internal/model"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetByID(id int64) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *GormUserRepository) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *GormUserRepository) ListPaged(limit, offset int) ([]model.User, int64, error) {
	var total int64
	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var users []model.User
	if err := r.db.Order("id ASC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *GormUserRepository) Update(u *model.User) error {
	res := r.db.Model(&model.User{}).Where("id = ?", u.ID).Updates(u)
	return res.Error
}

func (r *GormUserRepository) Delete(id int64) error {
	return r.db.Delete(&model.User{}, id).Error
}
