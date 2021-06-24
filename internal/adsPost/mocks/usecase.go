package mocks

import (
	"errors"
	"test_task_advertising/internal/adsPost/constants"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
)

type AdsPostUseCaseMock struct {
}

func (a AdsPostUseCaseMock) CreateAdsPost(adsPost *models.AdsPost) (*models.AdsPostId, error) {
	if adsPost.Title == "conflictTitle" {
		return &models.AdsPostId{}, errors.New(errorsConst.CONFLICT_UNIQUE_POST)
	}
	if len(adsPost.Title) > constants.MAX_TITLE_LENGTH || len(adsPost.Title) < constants.MIN_TITLE_LENGTH {
		return &models.AdsPostId{}, errors.New(errorsConst.BAD_TITLE_LENGTH)
	}
	return &models.AdsPostId{Id: 1}, nil
}

func (a AdsPostUseCaseMock) GetAdsPost(id uint64, fields []string) (*models.AdsPost, error) {
	for _, field := range fields {
		if field == constants.PHOTOS_FIELD || field == constants.DESCRIPTION_FIELD {
			continue
		}
		return &models.AdsPost{}, errors.New(errorsConst.BAD_REQUESTED_FIELDS)
	}
	return &models.AdsPost{Id: 1, Photos: []string{"url"}}, nil
}

func (a AdsPostUseCaseMock) GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem, error) {
	if count > constants.MAX_ADS_POST_COUNT_PER_REQUEST {
		return []models.AdsPostArrItem{}, errors.New(errorsConst.TOO_BIG_COUNT)

	}
	return []models.AdsPostArrItem{{Id: 1}, {Id: 2}, {Id: 3}}, nil
}
