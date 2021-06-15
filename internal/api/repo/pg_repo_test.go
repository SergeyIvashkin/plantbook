// +build integration

package repo_test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/proplants/plantbook/internal/api/handlers"
	"github.com/proplants/plantbook/internal/api/handlers/users"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/repo"
)

const (
	dbURLTemplate      string        = "postgres://plantbook_admin:mypassword@%s/plantbook_admin?sslmode=disable"
	dbDefaultHostPort  string        = "localhost:54321"
	dbHOSTAndPortEnv   string        = "DB_HOST_PORT"
	defaultTestTimeout time.Duration = 30 * time.Second
)

// initRepo creates Repo instance by using environment variable DB_HOST_PORT
// for define host and port db server.
func initRepo(ctx context.Context) (*repo.PG, error) {
	dbHostPort := os.Getenv(dbHOSTAndPortEnv)
	if dbHostPort == "" {
		dbHostPort = dbDefaultHostPort
	}
	return repo.NewPG(ctx, fmt.Sprintf(dbURLTemplate, dbHostPort), false)
}

// TestPG_FindUserByLogin testify extract users from the just created db.
func TestPG_FindUserByLogin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	pg, err := initRepo(ctx)
	if err != nil {
		t.Fatalf("initRepo failed, %s", err)
	}
	tests := []struct {
		name    string
		login   string
		want    *models.User
		want1   string
		wantErr bool
	}{
		{
			"just_created_db_root_user",
			"root",
			&models.User{ID: 1, Username: "root", UserRole: handlers.UserRoleAdmin},
			"love",
			false,
		},
		{
			"just_created_db_notexists_user",
			"_gardener_",
			&models.User{ID: 1999, Username: "_gardener_", UserRole: handlers.UserRoleAdmin},
			"love-garden",
			true,
		}, // user not found
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := pg.FindUserByLogin(ctx, tt.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("PG.FindUserByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got != nil {
				// clear fields for correct comparison
				got.Email = ""
				got.FirstName = ""
				got.LastName = ""
				got.Password = ""
				got.Phone = ""
				got.UserStatus = 0

			}
			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("PG.FindUserByLogin() got = %v, want %v", *got, *tt.want)
			}
			if !users.CheckPass(got1, tt.want1) {
				t.Errorf("PG.FindUserByLogin() got1 = %s, want %s", got1, tt.want1)
			}
		})
	}
}
