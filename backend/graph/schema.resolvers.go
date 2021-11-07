package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/neutronas/warehouse/backend/database"
	"github.com/neutronas/warehouse/backend/graph/generated"
	"github.com/neutronas/warehouse/backend/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateWarehouse(ctx context.Context, input *model.NewWarehouse) (*model.Warehouse, error) {
	return db.CreateWarehouse(input), nil
}

func (r *queryResolver) Warehouse(ctx context.Context, id string) (*model.Warehouse, error) {
	return db.FindWarehouseByID(id), nil
}

func (r *queryResolver) Warehouses(ctx context.Context) ([]*model.Warehouse, error) {
	return db.FindAllWarehouses(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
