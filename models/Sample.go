package models

import (
	"github.com/google/uuid"
)

type Sample struct {
	// BaseModel include ID , CreatedAt , UpdatedAt , DeletedAt
	BaseModel

	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Check  string  `json:"check"`
}

func (Sample) Create(Sample *Sample) error {
	Sample.ID = uuid.New()
	return Lamp().Create(&Sample).Error
}

func (Sample Sample) GetByID(id uuid.UUID) (*Sample, error) {
	Sample.ID = id
	err := Lamp().First(&Sample).Error
	return &Sample, err
}

func (Sample) Update(Sample *Sample, id uuid.UUID) error {
	Sample.ID = id
	return Lamp().Model(&Sample).Updates(Sample).Error

}

func (Sample Sample) Delete(id uuid.UUID) error {
	Sample.ID = id
	_, err := Sample.GetByID(id)
	if err != nil {
		return err
	}

	return Lamp().Delete(&Sample).Error
}

func (Sample) GetList(page, limit int) ([]Sample, int64, error) {
	var list []Sample
	var totalCount int64

	// Create a query builder
	query := Lamp().Model(&Sample{})

	// Retrieve the total count
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and retrieve the list
	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, totalCount, nil
}
