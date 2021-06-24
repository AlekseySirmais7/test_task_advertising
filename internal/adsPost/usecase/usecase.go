package usecase

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"test_task_advertising/internal/adsPost"
	"test_task_advertising/internal/adsPost/constants"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/uniqueSlice"
)

type AdsPostUseCase struct {
	adsPostRepo adsPost.IRepository
	logger      *zap.Logger
}

func NewUseCase(repo adsPost.IRepository, logger *zap.Logger) adsPost.IUseCase {
	return &AdsPostUseCase{
		adsPostRepo: repo,
		logger:      logger,
	}
}

func (au *AdsPostUseCase) CreateAdsPost(adsPost *models.AdsPost) (*models.AdsPostId, error) {

	if len(adsPost.Photos) < 1 || len(adsPost.Photos) > constants.MAX_PHOTOS_COUNT {
		err := errors.New(errorsConst.BAD_COUNT_OF_PHOTO_LINKS)
		au.logger.Error("AdsPostUseCase: CreateAdsPost : Check Photos length",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return &models.AdsPostId{}, err
	}

	if len(adsPost.Title) < constants.MIN_TITLE_LENGTH || len(adsPost.Title) > constants.MAX_TITLE_LENGTH {
		err := errors.New(errorsConst.BAD_TITLE_LENGTH)
		au.logger.Error("AdsPostUseCase: CreateAdsPost : Check title length",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return &models.AdsPostId{}, err
	}

	if len(adsPost.Description) < constants.MIN_DESCRIPTION_LENGTH ||
		len(adsPost.Description) > constants.MAX_DESCRIPTION_LENGTH {
		err := errors.New(errorsConst.BAD_DESCRIPTION_LENGTH)
		au.logger.Error("AdsPostUseCase: CreateAdsPost : Check description length",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return &models.AdsPostId{}, err
	}

	return au.adsPostRepo.CreateAdsPost(adsPost)
}

func (au *AdsPostUseCase) GetAdsPost(id uint64, fields []string) (*models.AdsPost, error) {

	// validate fields by white list
	for _, field := range fields {
		if field == constants.PHOTOS_FIELD || field == constants.DESCRIPTION_FIELD {
			continue
		}
		err := errors.New(errorsConst.BAD_REQUESTED_FIELDS)
		au.logger.Error("AdsPostUseCase: GetAdsPost : Check requested fields",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return &models.AdsPost{}, err
	}
	if !uniqueSlice.IsUniqueStrings(fields) {
		err := errors.New(errorsConst.BAD_REQUESTED_UNIQUE_FIELDS)
		au.logger.Error("AdsPostUseCase: GetAdsPost : Check unique fields",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return &models.AdsPost{}, err
	}

	return au.adsPostRepo.GetAdsPost(id, fields)
}

func (au *AdsPostUseCase) GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem,
	error) {

	if count > constants.MAX_ADS_POST_COUNT_PER_REQUEST {
		err := errors.New(errorsConst.TOO_BIG_COUNT)
		au.logger.Error("AdsPostUseCase: GetAdsPostArr : Check requested posts count",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return nil, err
	}

	if sort != constants.SORT_FIELD_DATE && sort != constants.SORT_FIELD_PRICE {
		err := errors.New(errorsConst.BAD_SORT_PARAM)
		au.logger.Error("AdsPostUseCase: GetAdsPostArr : Check sort parameter",
			zap.String("Error:", fmt.Sprintf("%v", err)),
		)
		return nil, err
	}

	return au.adsPostRepo.GetAdsPostArr(start, count, sort, desc)
}
