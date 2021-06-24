package http

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"test_task_advertising/internal/adsPost/mocks"
	"testing"
)

func TestCreateAdsPost(t *testing.T) {

	type TestCase struct {
		postJSON     string
		responseJSON string
	}

	t.Parallel()

	cases := []TestCase{
		{
			postJSON: "{\"title\":\"some car\",\"description\":\"it's cool car\",\"photos\":[\"link_1\"," +
				"\"link_2\"],\"price\":500000}",
			responseJSON: "{\"body\":{\"id\":1},\"code\":200,\"msg\":\"OK\"}\n",
		},
		{
			postJSON: "{\"title\":\"some\" car\",\"description\":\"it's cool car\",\"photos\":[\"link_1," +
				"\"link_2\"],\"price\":}",
			responseJSON: "{\"body\":null,\"code\":400,\"msg\":\"Bad JSON\"}\n",
		},
		{
			postJSON: "{\"title\":\"conflictTitle\",\"description\":\"it's cool car\",\"photos\":[\"link_1\"," +
				"\"link_2\"],\"price\":500000}",
			responseJSON: "{\"body\":null,\"code\":409,\"msg\":\"We already have this title\"}\n",
		},
		{
			postJSON: "{\"title\":\"" + strings.Repeat("tenSymbols", 101) + "\",\"description\":" +
				"\"it's cool car\",\"photos\":[\"link_1\",\"link_2\"],\"price\":500000}",
			responseJSON: "{\"body\":null,\"code\":400,\"msg\":\"You need send title with [3, 200] length\"}\n",
		},
	}

	for _, item := range cases {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/adsPost", strings.NewReader(item.postJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		adsPostHandler := &AdsPostHandler{AdsPostUseCase: mocks.AdsPostUseCaseMock{}}

		require.Equal(t, nil, adsPostHandler.CreateAdsPost(ctx))
		require.Equal(t, http.StatusOK, rec.Code)
		require.Equal(t, item.responseJSON, rec.Body.String())
	}
}

func TestGetAdsPost(t *testing.T) {

	type TestCase struct {
		postRequestJSON string
		responseJSON    string
	}

	t.Parallel()

	cases := []TestCase{
		{
			postRequestJSON: "{\"id\":1,\"fields\":[]}",
			responseJSON: "{\"body\":{\"id\":1,\"title\":\"\",\"description\":\"\",\"photos\":[\"url\"],\"price\"" +
				":0,\"date\":\"\"},\"code\":200,\"msg\":\"OK\"}\n",
		},
		{
			postRequestJSON: "{\"id\"1,\"fields\":]}",
			responseJSON:    "{\"body\":null,\"code\":400,\"msg\":\"Bad JSON\"}\n",
		},
		{
			postRequestJSON: "{\"id\":1,\"fields\":[\"bad_field\"]}",
			responseJSON:    "{\"body\":null,\"code\":400,\"msg\":\"Bad requested fields\"}\n",
		},
	}

	for _, item := range cases {

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/adsPost", strings.NewReader(item.postRequestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		adsPostHandler := &AdsPostHandler{AdsPostUseCase: mocks.AdsPostUseCaseMock{}}

		require.Equal(t, nil, adsPostHandler.GetAdsPost(ctx))
		require.Equal(t, http.StatusOK, rec.Code)
		require.Equal(t, item.responseJSON, rec.Body.String())
	}

}

func TestGetAdsPostArr(t *testing.T) {

	type TestCase struct {
		postRequestJSON string
		responseJSON    string
	}

	t.Parallel()

	cases := []TestCase{
		{
			postRequestJSON: "{\n  \"start\":0,\n  \"count\":10,\n  \"sort\":\"date\",\n  \"desc\":true\n}",
			responseJSON: "{\"body\":[{\"id\":1,\"title\":\"\",\"photo\":\"\",\"price\":0},{\"id\":2,\"title\":" +
				"\"\",\"photo\":\"\",\"price\":0},{\"id\":3,\"title\":\"\",\"photo\":\"\",\"price\":0}]," +
				"\"code\":200,\"msg\":\"OK\"}\n",
		},
		{
			postRequestJSON: "{\n  \"start\":0,\n  \"count\":0, sort\":\"date\",\n  \"desc\":true\n}",
			responseJSON:    "{\"body\":null,\"code\":400,\"msg\":\"Bad JSON\"}\n",
		},
		{
			postRequestJSON: "{\n  \"start\":0,\n  \"count\":11,\n  \"sort\":\"date\",\n  \"desc\":true\n}",
			responseJSON:    "{\"body\":null,\"code\":400,\"msg\":\"Too big count (max is 10)\"}\n",
		},
	}

	for _, item := range cases {

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/adsPosts", strings.NewReader(item.postRequestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		adsPostHandler := &AdsPostHandler{AdsPostUseCase: mocks.AdsPostUseCaseMock{}}

		require.Equal(t, nil, adsPostHandler.GetAdsPostArr(ctx))
		require.Equal(t, http.StatusOK, rec.Code)
		require.Equal(t, item.responseJSON, rec.Body.String())
	}

}
