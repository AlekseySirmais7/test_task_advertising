package http

import (
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"net/http"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/adsPost"
)

type AdsPostHandler struct {
	AdsPostUseCase adsPost.IUseCase
}

func (ah *AdsPostHandler) CreateAdsPost(ctx echo.Context) error {

	post := models.AdsPost{}

	err := easyjson.UnmarshalFromReader(ctx.Request().Body, &post)
	if err != nil {
		return ctx.JSON(http.StatusOK, &models.Answer{Code: http.StatusBadRequest, Msg: errorsConst.BAD_JSON})
	}

	postIdStruct, err := ah.AdsPostUseCase.CreateAdsPost(&post)
	if err != nil {

		if err.Error() == errorsConst.CONFLICT_UNIQUE_POST {
			// we can send recommendation in Msg
			return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusConflict, Msg: err.Error()})
		}

		if err.Error() == errorsConst.BAD_COUNT_OF_PHOTO_LINKS ||
			err.Error() == errorsConst.BAD_DESCRIPTION_LENGTH ||
			err.Error() == errorsConst.BAD_TITLE_LENGTH {
			return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusBadRequest, Msg: err.Error()})
		}

		return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusInternalServerError, Msg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &models.Answer{Body: postIdStruct, Code: http.StatusOK, Msg: "OK"})
}

func (ah *AdsPostHandler) GetAdsPost(ctx echo.Context) error {

	adsPostRequest := models.AdsPostRequest{}

	err := easyjson.UnmarshalFromReader(ctx.Request().Body, &adsPostRequest)
	if err != nil {
		return ctx.JSON(http.StatusOK, &models.Answer{Code: http.StatusBadRequest, Msg: errorsConst.BAD_JSON})
	}

	post, err := ah.AdsPostUseCase.GetAdsPost(adsPostRequest.Id, adsPostRequest.Fields)
	if err != nil {

		if err.Error() == errorsConst.BAD_REQUESTED_FIELDS ||
			err.Error() == errorsConst.BAD_REQUESTED_UNIQUE_FIELDS {
			return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusBadRequest, Msg: err.Error()})
		}
		if err.Error() == errorsConst.NOT_HAVE_POST_WITH_THIS_ID {
			return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusNotFound, Msg: err.Error()})
		}

		return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusInternalServerError, Msg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &models.Answer{Body: post, Code: http.StatusOK, Msg: "OK"})
}

func (ah *AdsPostHandler) GetAdsPostArr(ctx echo.Context) error {

	adsPostArrRequest := models.AdsPostArrRequest{}

	err := easyjson.UnmarshalFromReader(ctx.Request().Body, &adsPostArrRequest)
	if err != nil {
		return ctx.JSON(http.StatusOK, &models.Answer{Code: http.StatusBadRequest, Msg: errorsConst.BAD_JSON})
	}

	postArr, err := ah.AdsPostUseCase.GetAdsPostArr(adsPostArrRequest.Start, adsPostArrRequest.Count,
		adsPostArrRequest.Sort, adsPostArrRequest.Desc)
	if err != nil {

		if err.Error() == errorsConst.BAD_SORT_PARAM ||
			err.Error() == errorsConst.TOO_BIG_COUNT {
			return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusBadRequest, Msg: err.Error()})
		}

		return ctx.JSON(http.StatusOK, &models.Answer{Body: nil, Code: http.StatusInternalServerError, Msg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &models.Answer{Body: postArr, Code: http.StatusOK, Msg: "OK"})
}
