package adapters

import (
	"context"
	"errors"

	"surf-share/app/internal/modules/auth"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseAdapter struct {
	Db *pgxpool.Pool
}

// Connect establishes a connection to the PostgresSQL database using the provided connection string.
// It returns the connection pool or an error if the connection fails.
func (db *DatabaseAdapter) Connect(ctx context.Context, connStr string) error {
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return err
	}

	db.Db = pool
	return nil
}

// Close closes the database connection pool.
func (db *DatabaseAdapter) Close() {
	if db.Db != nil {
		db.Db.Close()
	}
}

// FindMany executes a query that returns multiple rows and scans the results into dest.
func (db *DatabaseAdapter) FindMany(ctx context.Context, dest any, query string, args ...any) error {
	if db.Db == nil {
		return errors.New("database is not initialized")
	}

	if err := pgxscan.Select(ctx, db.Db, dest, query, args...); err != nil {
		return err
	}

	return nil
}

// FindOne executes a query that returns a single row and scans the result into dest.
func (db *DatabaseAdapter) FindOne(ctx context.Context, dest any, query string, args ...any) error {
	if db.Db == nil {
		return errors.New("database is not initialized")
	}

	if err := pgxscan.Get(ctx, db.Db, dest, query, args...); err != nil {
		return err
	}

	return nil
}

// CreateOne executes an INSERT query with RETURNING clause and scans the created record into dest.
func (db *DatabaseAdapter) CreateOne(ctx context.Context, dest any, query string, args ...any) error {
	if db.Db == nil {
		return errors.New("database is not initialized")
	}

	if err := pgxscan.Get(ctx, db.Db, dest, query, args...); err != nil {
		return err
	}

	return nil
}

func (db *DatabaseAdapter) FindUserCredentialsByEmail(ctx context.Context, email string) (*auth.UserCredentials, error) {
	var userCredentials auth.UserCredentials
	if err := db.FindOne(ctx, &userCredentials,
		"SELECT id, username, email, password FROM app.users WHERE email = $1",
		email); err != nil {
		return nil, err
	}
	return &userCredentials, nil
}

func (db *DatabaseAdapter) FindUserByID(ctx context.Context, id string) (*auth.User, error) {
	var user auth.User
	if err := db.FindOne(ctx, &user,
		"SELECT id, username, email FROM app.users WHERE id = $1",
		id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DatabaseAdapter) CreateUser(ctx context.Context, username, email, hashedPassword string) (*auth.User, error) {
	var user auth.User
	if err := db.CreateOne(ctx, &user,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword); err != nil {
		return nil, err
	}
	return &user, nil
}
