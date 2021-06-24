package postgres

import (
	"errors"
	"github.com/stretchr/testify/require"
	"log"
	"os/exec"
	"strconv"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/config"
	"testing"
	"time"
)

func clearDBAfterTests() {
	clearDBCommand := exec.Command("psql", "-h", "127.0.0.1", "-U", "docker", "-d", "myService", "-p",
		"5432", "-c", "DELETE  FROM  adsPosts;")
	clearDBCommand.Env = []string{"PGPASSWORD=docker"}
	err := clearDBCommand.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// if we have test posts in test database, we can separate get and create post tests
// in real project better to have second(test) database and script for filling
func TestCreateAdsPostAndGetAdsPost(t *testing.T) {

	type TestCase struct {
		post              *models.AdsPost
		err               error
		errSecondCreation error
	}

	confDB, err := config.GetTestConfig()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := NewAdsPostRepository(confDB)
	if err != nil {
		log.Println("TEST REPO, IS TEST DATABASE RUNNING?")
		log.Fatal(err)
	}
	defer func() {
		errClose := repo.CloseAdsPost()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}()

	testCase := TestCase{
		post: &models.AdsPost{Title: "testTitle" + time.Now().String(), Description: "some descp", Price: 123,
			Photos: []string{"link1", "link2"}},
		err:               nil,
		errSecondCreation: errors.New(errorsConst.CONFLICT_UNIQUE_POST),
	}

	postId, err := repo.CreateAdsPost(testCase.post)

	require.Equal(t, testCase.err, err)

	selectedPost, err := repo.GetAdsPost(postId.Id, []string{"photos", "description"})

	require.Equal(t, testCase.err, err)

	testCase.post.Id = postId.Id // determination in DB
	testCase.post.Date = selectedPost.Date

	require.Equal(t, testCase.post, selectedPost)

	_, err = repo.CreateAdsPost(testCase.post)
	require.Equal(t, testCase.errSecondCreation, err)

	// test GetAdsPost fields options

	// description
	postWithDescription := *selectedPost
	postWithDescription.Photos = selectedPost.Photos[:1]
	selectedWithDescription, err := repo.GetAdsPost(postId.Id, []string{"description"})
	require.Equal(t, nil, err)
	require.Equal(t, &postWithDescription, selectedWithDescription)

	// photos
	postWithPhotos := *selectedPost
	postWithPhotos.Description = ""
	selectedWithPhotos, err := repo.GetAdsPost(postId.Id, []string{"photos"})
	require.Equal(t, nil, err)
	require.Equal(t, &postWithPhotos, selectedWithPhotos)

	// not description and not photos
	postNoPhotosNoDescription := *selectedPost
	postNoPhotosNoDescription.Description = ""
	postNoPhotosNoDescription.Photos = selectedPost.Photos[:1]
	selectedWithNoPhotosNoDescription, err := repo.GetAdsPost(postId.Id, []string{})
	require.Equal(t, nil, err)
	require.Equal(t, &postNoPhotosNoDescription, selectedWithNoPhotosNoDescription)

	// bad field
	_, err = repo.GetAdsPost(postId.Id, []string{"BAD FIELD"})
	require.Equal(t, errors.New(errorsConst.BAD_REQUESTED_FIELDS), err)

	// not have posts
	_, err = repo.GetAdsPost(77777777, []string{})
	require.Equal(t, errors.New(errorsConst.NOT_HAVE_POST_WITH_THIS_ID), err)

	clearDBAfterTests()
}

func TestGetAdsPostArr(t *testing.T) {

	type TestCase struct {
		start         uint64
		count         uint64
		sort          string
		desc          bool
		err           error
		selectedPosts []models.AdsPostArrItem
	}

	confDB, err := config.GetTestConfig()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := NewAdsPostRepository(confDB)
	if err != nil {
		log.Println("TEST REPO, IS TEST DATABASE RUNNING?")
		log.Fatal(err)
	}
	defer func() {
		errClose := repo.CloseAdsPost()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}()

	createForSelectPosts := []models.AdsPost{
		{Title: "testTitle_1 " + time.Now().String(), Description: "some descp", Price: 22, Photos: []string{"link1"}},
		{Title: "testTitle_2 " + time.Now().String(), Description: "some descp", Price: 5, Photos: []string{"link1"}},
		{Title: "testTitle_3 " + time.Now().String(), Description: "some descp", Price: 11, Photos: []string{"link1"}}}

	// creating post for future selection
	for index, item := range createForSelectPosts {
		postId, err := repo.CreateAdsPost(&item)
		require.Equal(t, nil, err)
		createForSelectPosts[index].Id = postId.Id
	}

	cases := []TestCase{

		{start: 0, count: 3, sort: "date", desc: true, err: nil,
			selectedPosts: []models.AdsPostArrItem{
				{Id: createForSelectPosts[2].Id, Title: createForSelectPosts[2].Title,
					Photo: createForSelectPosts[2].Photos[0], Price: createForSelectPosts[2].Price},
				{Id: createForSelectPosts[1].Id, Title: createForSelectPosts[1].Title,
					Photo: createForSelectPosts[1].Photos[0], Price: createForSelectPosts[1].Price},
				{Id: createForSelectPosts[0].Id, Title: createForSelectPosts[0].Title,
					Photo: createForSelectPosts[0].Photos[0], Price: createForSelectPosts[0].Price},
			}},

		{start: 1, count: 2, sort: "date", desc: true, err: nil,
			selectedPosts: []models.AdsPostArrItem{
				{Id: createForSelectPosts[1].Id, Title: createForSelectPosts[1].Title,
					Photo: createForSelectPosts[1].Photos[0], Price: createForSelectPosts[1].Price},
				{Id: createForSelectPosts[0].Id, Title: createForSelectPosts[0].Title,
					Photo: createForSelectPosts[0].Photos[0], Price: createForSelectPosts[0].Price},
			}},

		{start: 0, count: 3, sort: "price", desc: false, err: nil,
			selectedPosts: []models.AdsPostArrItem{
				{Id: createForSelectPosts[1].Id, Title: createForSelectPosts[1].Title,
					Photo: createForSelectPosts[1].Photos[0], Price: createForSelectPosts[1].Price},
				{Id: createForSelectPosts[2].Id, Title: createForSelectPosts[2].Title,
					Photo: createForSelectPosts[2].Photos[0], Price: createForSelectPosts[2].Price},
				{Id: createForSelectPosts[0].Id, Title: createForSelectPosts[0].Title,
					Photo: createForSelectPosts[0].Photos[0], Price: createForSelectPosts[0].Price},
			}},
		{start: 0, count: 3, sort: "price", desc: true, err: nil,
			selectedPosts: []models.AdsPostArrItem{
				{Id: createForSelectPosts[0].Id, Title: createForSelectPosts[0].Title,
					Photo: createForSelectPosts[0].Photos[0], Price: createForSelectPosts[0].Price},
				{Id: createForSelectPosts[2].Id, Title: createForSelectPosts[2].Title,
					Photo: createForSelectPosts[2].Photos[0], Price: createForSelectPosts[2].Price},
				{Id: createForSelectPosts[1].Id, Title: createForSelectPosts[1].Title,
					Photo: createForSelectPosts[1].Photos[0], Price: createForSelectPosts[1].Price},
			}},
		{start: 1, count: 2, sort: "price", desc: true, err: nil,
			selectedPosts: []models.AdsPostArrItem{
				{Id: createForSelectPosts[2].Id, Title: createForSelectPosts[2].Title,
					Photo: createForSelectPosts[2].Photos[0], Price: createForSelectPosts[2].Price},
				{Id: createForSelectPosts[1].Id, Title: createForSelectPosts[1].Title,
					Photo: createForSelectPosts[1].Photos[0], Price: createForSelectPosts[1].Price},
			}},
	}

	for caseNum, item := range cases {

		caseNumLabel := "Case №" + strconv.Itoa(caseNum+1)

		posts, err := repo.GetAdsPostArr(item.start, item.count, item.sort, item.desc)

		require.Equal(t, item.err, err, caseNumLabel)

		require.Equal(t, item.selectedPosts, posts, caseNumLabel)

	}

	clearDBAfterTests()
}
