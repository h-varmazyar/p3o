package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/internal/repositories"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "links"

const (
	ColumnId         = "id"
	ColumnOwnerId    = "owner_id"
	ColumnTotalVisit = "total_visit"
	ColumnKey        = "key"
)

type Repository struct {
	*gorm.DB
	log *log.Logger
}

func New(log *log.Logger, db *gorm.DB) Repository {
	return Repository{
		DB:  db,
		log: log,
	}
}

func (r Repository) Create(ctx context.Context, link entities.Link) (entities.Link, error) {
	err := r.DB.WithContext(ctx).Save(&link).Error
	if err != nil {
		return entities.Link{}, errors.ErrFailedToCreateLink.AddOriginalError(err)
	}
	return link, nil
}

func (r Repository) TotalLinkCount(ctx context.Context, userId uint) (int64, error) {
	count := int64(0)
	if err := r.DB.WithContext(ctx).
		Table(tableName).
		Where(repositories.Where(ColumnOwnerId), userId).
		Count(&count).Error; err != nil {
		return 0, errors.ErrLinkCountFetchFailed.AddOriginalError(err)
	}
	return count, nil
}

func (r Repository) TotalVisits(ctx context.Context, userId uint) (int64, error) {
	sum := int64(0)
	if err := r.DB.WithContext(ctx).Table(tableName).Select(repositories.Sum(ColumnTotalVisit)).Where(repositories.Where(ColumnOwnerId), userId).Row().Scan(&sum); err != nil {
		return 0, errors.ErrVisitCountFetchFailed.AddOriginalError(err)
	}
	return sum, nil
}

func (r Repository) ReturnById(ctx context.Context, id uint) (entities.Link, error) {
	link := entities.Link{}
	err := r.DB.WithContext(ctx).Table(tableName).Where(repositories.Where(ColumnId), id).First(&link).Error
	if err != nil {
		return entities.Link{}, errors.ErrLinkNotFound.AddOriginalError(err)
	}
	return link, nil
}

func (r Repository) ReturnByKey(ctx context.Context, key string) (entities.Link, error) {
	link := entities.Link{}
	err := r.DB.WithContext(ctx).Table(tableName).Where(repositories.Where(ColumnKey), key).First(&link).Error
	if err != nil {
		return entities.Link{}, errors.ErrLinkNotFound.AddOriginalError(err)
	}
	return link, nil
}

func (r Repository) List(ctx context.Context, userId uint) ([]entities.Link, error) {
	links := make([]entities.Link, 0)

	if err := r.DB.WithContext(ctx).Table(tableName).Where(repositories.Where(ColumnOwnerId), userId).Find(&links).Error; err != nil {
		return nil, errors.ErrUserHasNoLinks.AddOriginalError(err)
	}

	return links, nil
}

func (r Repository) Update(ctx context.Context, link entities.Link) error {
	err := r.DB.WithContext(ctx).Updates(&link).Error

	if err != nil {
		return errors.ErrUnexpected.AddOriginalError(err)
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, key string) error {
	err := r.DB.WithContext(ctx).Delete(&entities.Link{Key: key}).Error
	if err != nil {
		return errors.ErrUnexpected.AddOriginalError(err)
	}

	return nil
}

func (r Repository) Visit(_ context.Context, id uint) error {
	err := r.DB.
		Model(new(entities.Link)).
		Where(repositories.Where(ColumnId), id).
		Update("total_visit", gorm.Expr("total_visit + 1")).Error

	if err != nil {
		return errors.ErrIncreaseVisitFailed.AddOriginalError(err)
	}
	return nil
}
