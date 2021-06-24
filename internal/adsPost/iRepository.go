package adsPost

import "test_task_advertising/internal/models"

type IRepository interface {
	CreateAdsPost(adsPost *models.AdsPost) (*models.AdsPostId, error)
	GetAdsPost(id uint64, fields []string) (*models.AdsPost, error)
	GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem, error)
	CloseAdsPost() error
}
