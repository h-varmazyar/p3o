package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

const tableName = "links"

const (
	ColumnUserId     = "user_id"
	ColumnTotalVisit = "total_visit"
	ColumnKey        = "key"
)

type postgresRepository struct {
	*db.DB
	log *log.Logger
}

type Params struct {
	fx.In

	Log     *log.Logger
	DB      *db.DB
	Context context.Context
}

type Result struct {
	fx.Out

	Repo postgresRepository
}

func New(p Params) (Result, error) {
	postgresModel := postgresRepository{
		DB:  p.DB,
		log: p.Log,
	}
	return Result{
		Repo: postgresModel,
	}, nil
}

func (r postgresRepository) Create(ctx context.Context, link entities.Link) (entities.Link, error) {
	err := r.PostgresDB.WithContext(ctx).Save(&link).Error
	if err != nil {
		return entities.Link{}, errors.ErrFailedToCreateLink.AddOriginalError(err)
	}
	return link, nil
}

func (r postgresRepository) TotalLinkCount(ctx context.Context, userId uint) (int64, error) {
	count := int64(0)
	where := fmt.Sprintf("%v = %v", ColumnUserId, userId)
	if err := r.PostgresDB.WithContext(ctx).Table(tableName).Where(where).Count(&count).Error; err != nil {
		return 0, errors.ErrLinkCountFetchFailed.AddOriginalError(err)
	}
	return count, nil
}

func (r postgresRepository) Visits(ctx context.Context, userId uint) (int64, error) {
	sum := int64(0)
	selectQuery := fmt.Sprintf("SUM(%v)", ColumnTotalVisit)
	where := fmt.Sprintf("%v = %v", ColumnUserId, userId)
	if err := r.PostgresDB.WithContext(ctx).Table(tableName).Select(selectQuery).Where(where).Row().Scan(&sum); err != nil {
		return 0, errors.ErrVisitCountFetchFailed.AddOriginalError(err)
	}
	return sum, nil
}

func (r postgresRepository) ReturnByKey(ctx context.Context, key string) (entities.Link, error) {
	link := entities.Link{}
	where := fmt.Sprintf("%v = %v", ColumnKey, key)
	err := r.PostgresDB.WithContext(ctx).Table(tableName).Where(where).First(&link).Error
	if err != nil {
		return entities.Link{}, errors.ErrLinkNotFound.AddOriginalError(err)
	}
	return link, nil
}

func (r postgresRepository) List(ctx context.Context, userId uint) ([]entities.Link, error) {
	links := make([]entities.Link, 0)

	where := fmt.Sprintf("%v = %v", ColumnUserId, userId)
	if err := r.PostgresDB.WithContext(ctx).Table(tableName).Where(where).Find(&links).Error; err != nil {
		return nil, errors.ErrUserHasNoLinks.AddOriginalError(err)
	}

	return links, nil
}

func (r postgresRepository) Update(ctx context.Context, link entities.Link) error {
	err := r.PostgresDB.WithContext(ctx).Updates(&link).Error

	if err != nil {
		return errors.ErrUnexpected.AddOriginalError(err)
	}

	return nil
}

func (r postgresRepository) Delete(ctx context.Context, key string) error {
	err := r.PostgresDB.WithContext(ctx).Delete(&entities.Link{Key: key}).Error
	if err != nil {
		return errors.ErrUnexpected.AddOriginalError(err)
	}

	return nil
}

func (r postgresRepository) Visit(_ context.Context, id uint) error {
	err := r.PostgresDB.
		Model(new(entities.Link)).
		Where("id = ?", id).
		Update("total_visit", gorm.Expr("total_visit + 1")).Error

	if err != nil {
		return errors.ErrIncreaseVisitFailed.AddOriginalError(err)
	}
	return nil
}
