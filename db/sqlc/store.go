package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mniudanri/go-auth-paseto/util"
	"github.com/rs/zerolog/log"
)

const UniqueViolation = "23505"

var (
	ErrRecordNotFound  = pgx.ErrNoRows
	ErrUniqueViolation = &pgconn.PgError{
		Code: UniqueViolation,
	}
)

// Store interface defines all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	ConnPool *pgxpool.Pool
	*Queries
}

// create SQL query to pg pool
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		ConnPool: connPool,
		Queries:  New(connPool),
	}
}

// create connection to pg pool from db source path
func CreateConnection(config util.Config) Store {
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	return NewStore(connPool)
}

// check error when connecting to postgres
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
