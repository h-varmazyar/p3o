package auth

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/entities"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

const tableName = "users"

type Postgres struct {
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

	postgresModel := &Postgres{
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

func (r *Postgres) Create(ctx context.Context, link *entities.User) error {
	err := r.PostgresDB.Save(link).Error
	if err != nil {
		return ErrFailedToCreateUser.AddOriginalError(err)
	}
	return nil
}

func (r *Postgres) ReturnByMobile(ctx context.Context, mobile string) (*entities.User, error) {
	//err := r.PostgresDB.Save(link).Error
	//if err != nil {
	//	return ErrFailedToCreateUser.AddOriginalError(err)
	//}
	return nil, nil
}

func (r *Postgres) ReturnByEmail(ctx context.Context, email string) (*entities.User, error) {
	//err := r.PostgresDB.Save(link).Error
	//if err != nil {
	//	return ErrFailedToCreateUser.AddOriginalError(err)
	//}
	return nil, nil
}
