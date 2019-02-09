package graphql

import (
	"context"

	types "rocket-server/server/api/pkg/types"
	endpoint "rocket-server/server/api/pkg/endpoint"
)

type Resolver struct {
	e endpoint.Endpoints
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateEvent(ctx context.Context, input types.NewEvent) (types.Event, error) {
	// create types.NewEvent from input
	panic("not implemented")
	
	//return r.e.AddEvent(ctx, newEvent)
}
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (string, error) {
	// there is no service for deletion yet
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Events(ctx context.Context) ([]types.Event, error) {
	return r.e.GetEvents(ctx)
}
