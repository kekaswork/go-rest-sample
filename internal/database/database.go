package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	ctx        context.Context
	connection *pgx.Conn
}

var (
	instance *Service
	once     sync.Once
)

func NewService() *Service {
	once.Do(func() {
		address := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, address)
		if err != nil {
			log.Fatal(err)
		}

		instance = &Service{
			ctx:        ctx,
			connection: conn,
		}

		if err := instance.applySchema(); err != nil {
			log.Fatalf("failed to apply schema: %v", err)
		}
	})

	return instance
}

func (s *Service) applySchema() error {
	if s.isCreated() {
		return nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}
	schemaPath := filepath.Join(cwd, "internal", "database", "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	_, err = s.connection.Exec(s.ctx, string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil
}

func (s *Service) isCreated() bool {
	var count int
	err := s.connection.QueryRow(context.Background(), "SELECT COUNT(table_name) FROM INFORMATION_SCHEMA.TABLES WHERE table_name = 'students' OR table_name = 'marks' OR table_name = 'subjects'").Scan(&count)
	if err != nil {
		return false
	}

	return count == 3
}

func (s *Service) Close() {
	s.connection.Close(s.ctx)
}

func (s *Service) GetConn() *pgx.Conn {
	return s.connection
}

func (s *Service) GetCtx() context.Context {
	return s.ctx
}
