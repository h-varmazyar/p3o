package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/entities"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

const tableName = "links"

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

	Model Model
}

func New(p Params) (Result, error) {
	if err := migration(p.Context, p.DB); err != nil {
		return Result{}, err
	}

	postgresModel := &postgresRepository{
		DB:  p.DB,
		log: p.Log,
	}
	return Result{
		Model: postgresModel,
	}, nil
}

func migration(_ context.Context, dbInstance *db.DB) error {
	var err error
	migrations := make(map[string]struct{})
	tags := make([]string, 0)
	err = dbInstance.PostgresDB.Table(db.MigrationTable).Where("table_name = ?", tableName).Select("tag").Find(&tags).Error
	if err != nil {
		return err
	}

	for _, tag := range tags {
		migrations[tag] = struct{}{}
	}

	newMigrations := make([]*db.Migration, 0)
	err = dbInstance.PostgresDB.Transaction(func(tx *gorm.DB) error {
		if _, ok := migrations["v1.0.0"]; !ok {
			err = tx.AutoMigrate(new(entities.Link))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				TableName:   tableName,
				Tag:         "v1.0.0",
				Description: fmt.Sprintf("create %s table", tableName),
			})
		}
		err = tx.Model(new(db.Migration)).CreateInBatches(&newMigrations, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) Create(ctx context.Context, link *entities.Link) error {
	err := r.PostgresDB.Save(link).Error
	if err != nil {
		return ErrFailedToCreateLink.AddOriginalError(err)
	}
	return nil
}

func (r *postgresRepository) TotalCounts(_ context.Context) (int64, error) {
	count := int64(0)
	if err := r.PostgresDB.Table(tableName).Count(&count).Error; err != nil {
		return 0, ErrLinkCountFetchFailed.AddOriginalError(err)
	}
	return count, nil
}

func (r *postgresRepository) TotalVisits(ctx context.Context) (int64, error) {
	sum := int64(0)
	if err := r.PostgresDB.Table(tableName).Select("sum(total_visit)").Row().Scan(&sum); err != nil {
		return 0, ErrVisitCountFetchFailed.AddOriginalError(err)
	}
	return sum, nil
}

func (r *postgresRepository) ReturnByKey(_ context.Context, key string) (*entities.Link, error) {
	link := new(entities.Link)
	err := r.PostgresDB.Model(new(entities.Link)).Where("key = ?", key).First(link).Error
	if err != nil {
		return nil, ErrLinkNotFound.AddOriginalError(err)
	}
	return link, nil
}

func (r *postgresRepository) Visit(_ context.Context, id uint) error {
	err := r.PostgresDB.
		Model(new(entities.Link)).
		Where("id = ?", id).
		Update("total_visit", gorm.Expr("total_visit + 1")).Error

	if err != nil {
		return ErrIncreaseVisitFailed.AddOriginalError(err)
	}
	return nil
}
