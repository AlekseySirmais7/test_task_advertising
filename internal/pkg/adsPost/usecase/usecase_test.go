package usecase

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/adsPost/mocks"
	"testing"
)

func TestCreateAdsPost(t *testing.T) {

	type TestCase struct {
		post   models.AdsPost
		postId models.AdsPostId
		err    error
	}

	t.Parallel()

	useCase := NewUseCase(mocks.AdsPostRepositoryMock{})

	cases := []TestCase{
		{
			post: models.AdsPost{Photos: []string{"1"}, Title: "cat",
				Description: "cat"},
			postId: models.AdsPostId{Id: 1},
			err:    nil,
		},
		{
			post:   models.AdsPost{Photos: []string{"1", "2", "3", "4"}, Title: "cat", Description: "cat"},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_COUNT_OF_PHOTO_LINKS),
		},
		{
			post:   models.AdsPost{Photos: []string{}, Title: "cat", Description: "cat"},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_COUNT_OF_PHOTO_LINKS),
		},
		{
			post:   models.AdsPost{Photos: []string{"1", "2"}, Title: "cat", Description: "=("},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_DESCRIPTION_LENGTH),
		},
		{
			post: models.AdsPost{Photos: []string{"1", "2"}, Title: "cat",
				Description: strings.Repeat("tenSymbols", 101)},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_DESCRIPTION_LENGTH),
		},
		{
			post:   models.AdsPost{Photos: []string{"1", "2"}, Title: "", Description: "cat"},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_TITLE_LENGTH),
		},
		{
			post: models.AdsPost{Photos: []string{"1", "2"}, Title: strings.Repeat("tenSymbols", 21),
				Description: "cat"},
			postId: models.AdsPostId{},
			err:    errors.New(errorsConst.BAD_TITLE_LENGTH),
		},
	}

	for caseNum, item := range cases {

		caseNumLabel := "Case №" + strconv.Itoa(caseNum+1)

		postId, err := useCase.CreateAdsPost(&item.post)

		require.Equal(t, item.postId, postId, caseNumLabel)

		require.Equal(t, item.err, err, caseNumLabel)

		//if !reflect.DeepEqual(postId, item.postId) {
		//	t.Errorf("[%d] wrong result: got %+v, expected %+v",
		//		caseNum, postId, item.postId)
		//}
		//
		//if !require.Equal(err, item.err) {
		//	t.Errorf("[%d] wrong err: got %+v, expected %+v",
		//		caseNum, err, item.err)
		//}
	}

}

func TestGetAdsPost(t *testing.T) {

	type TestCase struct {
		postReqParams models.AdsPostRequest
		post          models.AdsPost
		err           error
	}

	t.Parallel()

	useCase := NewUseCase(mocks.AdsPostRepositoryMock{})

	cases := []TestCase{
		{
			postReqParams: models.AdsPostRequest{Id: 1, Fields: []string{"photos", "description"}},
			post:          models.AdsPost{Id: 1},
			err:           nil,
		},
		{
			postReqParams: models.AdsPostRequest{Id: 1, Fields: []string{"BAD FIELD"}},
			post:          models.AdsPost{},
			err:           errors.New(errorsConst.BAD_REQUESTED_FIELDS),
		},
		{
			postReqParams: models.AdsPostRequest{Id: 1, Fields: []string{"photos", "photos"}},
			post:          models.AdsPost{},
			err:           errors.New(errorsConst.BAD_REQUESTED_UNIQUE_FIELDS),
		},
	}

	for caseNum, item := range cases {

		caseNumLabel := "Case №" + strconv.Itoa(caseNum+1)

		post, err := useCase.GetAdsPost(item.postReqParams.Id, item.postReqParams.Fields)

		require.Equal(t, item.post, post, caseNumLabel)

		require.Equal(t, item.err, err, caseNumLabel)

	}
}

func TestGetAdsPostArr(t *testing.T) {

	type TestCase struct {
		postReqArrParams models.AdsPostArrRequest
		posts            []models.AdsPostArrItem
		err              error
	}

	t.Parallel()

	useCase := NewUseCase(mocks.AdsPostRepositoryMock{})

	cases := []TestCase{
		{
			postReqArrParams: models.AdsPostArrRequest{Start: 0, Count: 10, Sort: "price", Desc: false},
			posts:            []models.AdsPostArrItem{{Id: 1}, {Id: 2}, {Id: 3}},
			err:              nil,
		},
		{
			postReqArrParams: models.AdsPostArrRequest{Start: 0, Count: 11, Sort: "price", Desc: false},
			posts:            nil,
			err:              errors.New(errorsConst.TOO_BIG_COUNT),
		},
		{
			postReqArrParams: models.AdsPostArrRequest{Start: 0, Count: 10, Sort: "BAD_SORT_FIELD", Desc: false},
			posts:            nil,
			err:              errors.New(errorsConst.BAD_SORT_PARAM),
		},
	}

	for caseNum, item := range cases {

		caseNumLabel := "Case №" + strconv.Itoa(caseNum+1)

		posts, err := useCase.GetAdsPostArr(item.postReqArrParams.Start, item.postReqArrParams.Count,
			item.postReqArrParams.Sort, item.postReqArrParams.Desc)

		require.Equal(t, item.posts, posts, caseNumLabel)

		require.Equal(t, item.err, err, caseNumLabel)

	}
}
