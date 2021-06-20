package usecase

import (
	"errors"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/adsPost"
	"test_task_advertising/pkg/uniqueSlice"
)

type AdsPostUseCase struct {
	AdsPostRepo adsPost.IRepository
}

func NewUseCase(repo adsPost.IRepository) adsPost.IUseCase {
	return &AdsPostUseCase{
		AdsPostRepo: repo,
	}
}

func (au *AdsPostUseCase) CreateAdsPost(adsPost *models.AdsPost) (models.AdsPostId, error) {

	if len(adsPost.Photos) < 1 || len(adsPost.Photos) > 3 {
		return models.AdsPostId{}, errors.New(errorsConst.BAD_COUNT_OF_PHOTO_LINKS)
	}

	if len(adsPost.Title) < 3 || len(adsPost.Title) > 200 {
		return models.AdsPostId{}, errors.New(errorsConst.BAD_TITLE_LENGTH)
	}

	if len(adsPost.Description) < 3 || len(adsPost.Description) > 1000 {
		return models.AdsPostId{}, errors.New(errorsConst.BAD_DESCRIPTION_LENGTH)
	}

	post, err := au.AdsPostRepo.CreateAdsPost(adsPost)
	//if err != nil {
	//	// todo wrap
	//	return post, err
	//}
	return post, err
}

func (au *AdsPostUseCase) GetAdsPost(id uint64, fields []string) (models.AdsPost, error) {

	// validate fields by white list
	for _, field := range fields {
		if field == "photos" || field == "description" {
			continue
		}
		return models.AdsPost{}, errors.New(errorsConst.BAD_REQUESTED_FIELDS)
	}
	if !uniqueSlice.IsUniqueStrings(fields) {
		return models.AdsPost{}, errors.New(errorsConst.BAD_REQUESTED_UNIQUE_FIELDS)
	}

	post, err := au.AdsPostRepo.GetAdsPost(id, fields)
	//if err != nil {
	//	// todo Wrap
	//	return models.AdsPost{}, err
	//}
	return post, err
}

func (au *AdsPostUseCase) GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem, error) {

	if count > 10 {
		return nil, errors.New(errorsConst.TOO_BIG_COUNT)
	}

	if sort != "date" && sort != "price" {
		return nil, errors.New(errorsConst.BAD_SORT_PARAM)
	}

	postArr, err := au.AdsPostRepo.GetAdsPostArr(start, count, sort, desc)
	//if err != nil {
	//	// todo wrap
	//	return nil, err
	//}
	return postArr, err
}
