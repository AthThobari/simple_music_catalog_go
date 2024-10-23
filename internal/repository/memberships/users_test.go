package memberships

import (
	"reflect"
	"testing"
	"time"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		model memberships.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				model: memberships.User{
					Email:     "test@gmail.com",
					Username:  "test",
					Password:  "test",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: memberships.User{
					Email:     "test@gmail.com",
					Username:  "test",
					Password:  "test",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.CreatedBy,
						args.model.UpdatedBy,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.CreateUser(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	now := time.Now()

	type args struct {
		email    string
		username string
		id       uint
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mockFn func(args args)
	}{
		{
		name: "success",	
		args: args{
			email: "test@gmail.com",
			username: "test",
		},
		want: &memberships.User{
			Model: gorm.Model{
				ID: 1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			Email: "test@gmail.com",
			Username: "test",
			Password: "test",
			CreatedBy: "test@gmail.com",
			UpdatedBy: "test@gmail.com",
		},
		wantErr: false,
		mockFn: func(args args) {
			mock.ExpectBegin()

			mock.ExpectQuery(`SELECT FROM "users" .+`).WithArgs(args.email, args.username, args.id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "password", "created_by", "updated_by"}).AddRow(1, ))


			mock.ExpectCommit()
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUser() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
