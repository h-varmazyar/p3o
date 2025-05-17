package visit

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/internal/repositories"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const tableName = "visits"

const (
	columnId        = "id"
	columnUserId    = "user_id"
	columnLinkId    = "link_id"
	columnCreatedAt = "created_at"
)

type Repository struct {
	*gorm.DB
	log *log.Logger
}

type DailyCount struct {
	VisitDate  time.Time
	VisitCount int
}

func New(log *log.Logger, db *gorm.DB) Repository {
	repo := Repository{
		DB:  db,
		log: log,
	}
	return repo
}

func (r Repository) Create(ctx context.Context, visit entities.Visit) (entities.Visit, error) {
	if err := r.DB.WithContext(ctx).Model(&entities.Visit{}).Create(&visit).Error; err != nil {
		return entities.Visit{}, err
	}
	return visit, nil
}

func (r Repository) ReturnByID(ctx context.Context, id string) (entities.Visit, error) {
	var visit entities.Visit
	if err := r.DB.WithContext(ctx).Model(&entities.Visit{}).Where("id = ?", id).First(&visit).Error; err != nil {
		return entities.Visit{}, err
	}
	return visit, nil
}

func (r Repository) Update(ctx context.Context, visit entities.Visit) error {
	if err := r.DB.WithContext(ctx).Model(&entities.Visit{}).Where("id = ?", visit.ID).Updates(&visit).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) DailyVisitCount(ctx context.Context, userId, linkId uint, days uint) ([]DailyCount, error) {
	var results []DailyCount

	tx := r.DB.WithContext(ctx)
	query := `
    WITH last_n_days AS (
        SELECT generate_series(
            CURRENT_DATE - CAST(? AS INTEGER) * INTERVAL '1 day', 
            CURRENT_DATE, 
            INTERVAL '1 day'
        )::DATE AS visit_date
    )
    SELECT 
        l.visit_date, 
        COALESCE(COUNT(v.id), 0) AS visit_count
    FROM last_n_days l
    LEFT JOIN visits v 
        ON DATE(v.created_at) = l.visit_date
        AND v.user_id = ?
`
	if linkId > 0 {
		query += `
		WHERE v.link_id = ?
`
		tx.Raw(query, days-1, userId, linkId)
	} else {
		tx.Raw(query, days-1, userId)
	}
	query += `
    GROUP BY l.visit_date
    ORDER BY l.visit_date;
	`

	err := tx.Scan(&results).Error
	return results, err
}

func (r Repository) VisitCount(ctx context.Context, userId uint, from, to time.Time) (int64, error) {
	count := int64(0)
	if err := r.DB.WithContext(ctx).Model(&entities.Visit{}).
		Where(repositories.Where(columnUserId), userId).
		Where(repositories.Between(columnCreatedAt), from, to).
		Count(&count).Error; err != nil {
		return 0, errors.ErrVisitCountFetchFailed.AddOriginalError(err)
	}
	return count, nil
}

func (r Repository) GetUnhandled(ctx context.Context) ([]entities.Visit, error) {
	visits := make([]entities.Visit, 0)
	err := r.DB.WithContext(ctx).
		Model(&entities.Visit{}).
		Where("handled_at is null").
		Where("created_at < ?", time.Now().Add(-1*time.Minute)).
		Limit(100).Find(&visits).Error
	if err != nil {
		return nil, err
	}

	return visits, nil
}
