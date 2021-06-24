package mocks

import (
	"test_task_advertising/internal/models"
)

type AdsPostRepositoryMock struct {
}

func (a AdsPostRepositoryMock) CreateAdsPost(adsPost *models.AdsPost) (*models.AdsPostId, error) {
	return &models.AdsPostId{Id: 1}, nil
}

func (a AdsPostRepositoryMock) GetAdsPost(id uint64, fields []string) (*models.AdsPost, error) {
	return &models.AdsPost{Id: 1}, nil
}

func (a AdsPostRepositoryMock) GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem, error) {
	return []models.AdsPostArrItem{{Id: 1}, {Id: 2}, {Id: 3}}, nil
}

func (a AdsPostRepositoryMock) CloseAdsPost() error {
	return nil
}
