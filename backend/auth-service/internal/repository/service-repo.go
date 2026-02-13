package repository

import (
	"context"

	"github.com/Lemper29/Jarcy/auth-service/internal/database/queries"
	"github.com/Lemper29/Jarcy/auth-service/internal/models"
	"github.com/google/uuid"
)

type Repository struct {
	queries *queries.Queries
}

func NewRepo(queries *queries.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (q *Repository) CreateUser(ctx context.Context, req models.CreateUserParams) (queries.User, error) {
	dbParams := queries.CreateUserParams{
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
	}

	user, err := q.queries.CreateUser(ctx, dbParams)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (q *Repository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := q.queries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (queries.User, error) {
	user, err := q.queries.GetUserByID(ctx, id)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (q *Repository) GetUsers(ctx context.Context) ([]queries.User, error) {
	user, err := q.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (q *Repository) UpdateUser(ctx context.Context, req queries.UpdateUserParams) error {
	err := q.queries.UpdateUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
