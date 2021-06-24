package postgres

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"log"
	"sort"
	"strconv"
	"strings"
	"test_task_advertising/internal/adsPost"
	"test_task_advertising/internal/adsPost/constants"
	"test_task_advertising/internal/errorsConst"
	"test_task_advertising/internal/models"
	"test_task_advertising/internal/pkg/config"
)

type AdsPostRepository struct {
	adsPostCon *sql.DB
}

func NewAdsPostRepository(config config.DataBaseConfig) (adsPost.IRepository, error) {
	postgresConnStr := "user=" + config.User + " password=" + config.Password +
		" dbname=" + config.Database + " sslmode=disable port=" + strconv.Itoa(config.Port) + " host=" + config.Host
	postgresDB, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		return AdsPostRepository{}, err
	}
	postgresDB.SetMaxOpenConns(config.ConnectionCount)

	errPing := postgresDB.Ping()
	if errPing != nil {
		return AdsPostRepository{}, errPing
	}

	return &AdsPostRepository{adsPostCon: postgresDB}, nil
}

func (ar AdsPostRepository) CloseAdsPost() error {
	errClose := ar.adsPostCon.Close()
	return errClose
}

func (ar AdsPostRepository) CreateAdsPost(adsPost *models.AdsPost) (*models.AdsPostId, error) {

	queryRusTable := sq.Insert("AdsPosts").
		Columns("title", "description", "photos", "price", "date").
		Values(adsPost.Title, adsPost.Description, pq.Array(adsPost.Photos), adsPost.Price, sq.Expr("now()")).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(ar.adsPostCon)

	adsPostId := models.AdsPostId{}
	err := queryRusTable.QueryRow().Scan(&adsPostId.Id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint \"adsposts_title_key\"") {
			return &adsPostId, errors.New(errorsConst.CONFLICT_UNIQUE_POST)
		}
		return &models.AdsPostId{}, err
	}

	return &adsPostId, nil
}

func (ar AdsPostRepository) GetAdsPost(id uint64, fields []string) (*models.AdsPost, error) {

	selectedFields := []string{"id", "title", "price", "to_char(date, 'HH:mm:ss DD-MM-YYYY')"}
	isNeedSelectAllPhotos := false
	isNeedSelectDescription := false

	functionalFields := make([]string, 0, len(fields))
	for _, field := range fields {

		switch field {
		case constants.PHOTOS_FIELD: // it's ok to check twice (in useCase and in repo)
			functionalFields = append(functionalFields, field) // funcs have to check contracts
			isNeedSelectAllPhotos = true
		case constants.DESCRIPTION_FIELD:
			functionalFields = append(functionalFields, field)
			isNeedSelectDescription = true
		default:
			return &models.AdsPost{}, errors.New(errorsConst.BAD_REQUESTED_FIELDS)
		}

	}

	if !isNeedSelectAllPhotos { // always need first photo
		selectedFields = append(selectedFields, "photos[1]")
	}

	sort.Strings(functionalFields) // order: description, photos
	selectedFields = append(selectedFields, functionalFields...)

	querySelect := sq.Select(selectedFields...).
		From("AdsPosts").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(ar.adsPostCon)

	post := models.AdsPost{Photos: make([]string, 1)}
	err := errors.New("")
	scanner := querySelect.QueryRow()

	if isNeedSelectAllPhotos && isNeedSelectDescription {
		err = scanner.Scan(&post.Id, &post.Title, &post.Price, &post.Date, &post.Description, pq.Array(&post.Photos))
	} else if isNeedSelectAllPhotos && !isNeedSelectDescription {
		err = scanner.Scan(&post.Id, &post.Title, &post.Price, &post.Date, pq.Array(&post.Photos))
	} else if !isNeedSelectAllPhotos && isNeedSelectDescription {
		err = scanner.Scan(&post.Id, &post.Title, &post.Price, &post.Date, &post.Photos[0], &post.Description)
	} else {
		err = scanner.Scan(&post.Id, &post.Title, &post.Price, &post.Date, &post.Photos[0])
	}
	if err != nil {

		if strings.Contains(err.Error(), "no rows in result set") {
			return &post, errors.New(errorsConst.NOT_HAVE_POST_WITH_THIS_ID)
		}
		return &post, err
	}
	return &post, nil
}

func (ar AdsPostRepository) GetAdsPostArr(start uint64, count uint64, sort string, desc bool) ([]models.AdsPostArrItem, error) {

	if desc {
		sort += " DESC"
	}

	querySelect := sq.Select("id", "title", "photos[1]", "price").
		From("AdsPosts").
		Offset(start).
		Limit(count).
		OrderBy(sort).
		PlaceholderFormat(sq.Dollar).
		RunWith(ar.adsPostCon)

	// make capacity = count. if select post count less than count(not enough posts in DB), we spend extra memory,
	// but I suppose that in 90% cases we use full memory
	posts := make([]models.AdsPostArrItem, 0, count)
	rows, err := querySelect.Query()
	if err != nil {
		return posts, err
	}
	defer func(rows *sql.Rows) {
		errClose := rows.Close()
		if errClose != nil {
			log.Println("Err Repo GetAdsPostArr rows.Close():" + errClose.Error())
		}
	}(rows)

	for rows.Next() {
		post := models.AdsPostArrItem{}
		err = rows.Scan(&post.Id, &post.Title, &post.Photo, &post.Price)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
