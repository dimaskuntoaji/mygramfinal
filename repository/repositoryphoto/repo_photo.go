package repositoryphoto

import (
	"finall/entity"

	"gorm.io/gorm"
)

type RepositoryPhoto interface {
	Create(data entity.Photo) (entity.Photo, error)
	GetPhotos() ([]entity.Photo, error)
	Update(data entity.Photo) (entity.Photo, error)
	Delete(id int) error
	GetPhotoByUserID(id uint) (entity.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryPhoto {
	return &repository{db: db}
}

func (r *repository) Update(data entity.Photo) (entity.Photo, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return data, nil
}

func (r *repository) Delete(id int) error {
	photo := entity.Photo{}
	photo.ID = uint(id)
	err := r.db.First(&photo).Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Create(data entity.Photo) (entity.Photo, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return data, nil
}

func (r *repository) GetPhotos() ([]entity.Photo, error) {
	var photo []entity.Photo
	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return []entity.Photo{}, err
	}
	return photo, nil
}

func (r *repository) GetPhotoByUserID(id uint) (entity.Photo, error) {
	var photo entity.Photo
	err := r.db.Preload("User").Where("user_id = ?", id).First(&photo).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return photo, nil
}
