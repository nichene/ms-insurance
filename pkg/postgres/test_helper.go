package postgres

import (
	"context"
	"database/sql"
	"ms-insurance/config"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pressly/goose"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TestDbConfig struct {
	testRef       *testing.T
	ctx           *context.Context
	cfg           *config.Config
	conn          *gorm.DB
	dbConn        *sql.DB
	migrationsDir string
}

var (
	_, b, _, _     = runtime.Caller(0)
	ProjectRootDir = filepath.Join(filepath.Dir(b), "../..")
)

func NewTestDb(testRef *testing.T, ctx context.Context, cfg *config.Config) *TestDbConfig {
	pgTestContainer, err := NewPostgresTestContainer(ctx, cfg)
	if err != nil {
		testRef.Fatal(err)
	}

	cfg.DBPort, err = pgTestContainer.Port(ctx)
	if err != nil {
		testRef.Fatal(err)
	}

	cfg.DBHost, err = pgTestContainer.Host(ctx)
	if err != nil {
		testRef.Fatal(err)
	}

	conn, err := InitDatabase(ctx, cfg)
	assert.Nil(testRef, err)

	dbConn, err := conn.DB()
	assert.Nil(testRef, err)

	migrationsDir := filepath.Join(ProjectRootDir, "/migrations")

	return &TestDbConfig{testRef, &ctx, cfg, conn, dbConn, migrationsDir}
}

func (i *TestDbConfig) InitDatabase() *gorm.DB {
	err := goose.Up(i.dbConn, i.migrationsDir)
	assert.Nil(i.testRef, err)

	return i.conn
}

func (i *TestDbConfig) ClearDatabase() {
	err := goose.Reset(i.dbConn, i.migrationsDir)
	assert.Nil(i.testRef, err)
}
