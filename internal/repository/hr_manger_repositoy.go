package repository

import (
	"github.com/mohit-051/hirego/internal/core/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

var HRRepo *HRRepository

type HRRepository struct {
	*repository[domain.HrManager]
}

func (r *HRRepository) Initialize(db *mongo.Client) *HRRepository {
	HRRepo = &HRRepository{
		repository: &repository[domain.HrManager]{db: db},
	}

	return HRRepo
}
